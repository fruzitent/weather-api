package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"slices"
	"time"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/internal/lib/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/shared/domain/value"
	httpProbe "git.fruzit.pp.ua/weather/api/pkg/probe/adapter/primary/http"
	httpUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/primary/http"
	"git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/smtp"
	sqliteUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/sqlite"
	entityUser "git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	httpWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/http"
	"git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/weatherapi"
	sqliteWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/secondary/sqlite"
	coreWeather "git.fruzit.pp.ua/weather/api/pkg/weather/core"
	queryWeather "git.fruzit.pp.ua/weather/api/pkg/weather/core/query"
	entityWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
	valueWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

const (
	CMD_DAEMON = "daemon"
	CMD_HEALTH = "health"
)

func main() {
	ctx := context.Background()

	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("not enough arguments")
	}

	switch args[0] {
	case CMD_DAEMON:
		daemonCmd := flag.NewFlagSet(CMD_DAEMON, flag.ExitOnError)
		daemonCmd.Parse(args[1:])

		schema := slices.Concat(
			sqliteUser.Schema,
			sqliteWeather.Schema,
		)
		db, err := sqlite.Open(ctx, config.Sqlite.DataSourceName, schema)
		if err != nil {
			log.Fatal(err)
		}
		_ = db

		notifications(config)

		providerWeather, err := weatherapi.NewWeatherapi(&config.Weatherapi)
		if err != nil {
			log.Fatal(err)
		}

		appWeather := coreWeather.App{
			Command: coreWeather.Command{},
			Query: coreWeather.Query{
				Current: queryWeather.NewCurrentHandler(providerWeather),
			},
		}

		mux := http.NewServeMux()
		_ = httpProbe.New(mux)
		_ = httpUser.New(mux)
		_ = httpWeather.New(mux, &appWeather)
		log.Fatal(http.ListenAndServe(addr, mux))

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

		if err := httpProbe.IsHealthy(addr); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("invalid subcommand %s\n", args[0])
	}
}

func notifications(config *config.Config) {
	ntfyProvider := smtp.Smtp{
		Config: &config.Smtp,
	}

	user := (func() entityUser.User {
		id, err := value.NewId("hi")
		if err != nil {
			log.Fatal(err)
		}
		mail, err := value.NewMail("fruzit@fruzit.pp.ua")
		if err != nil {
			log.Fatal(err)
		}
		return entityUser.NewUser(*id, *mail)
	})()

	report := (func() entityWeather.Report {
		createdAt := time.Now().Unix()
		id, err := value.NewId("test-id")
		if err != nil {
			log.Fatal(err)
		}
		location, err := valueWeather.NewLocation("Kyiv")
		if err != nil {
			log.Fatal(err)
		}
		description := "Cloudy"
		humidity, err := valueWeather.NewHumidity(0.5)
		if err != nil {
			log.Fatal(err)
		}
		temperature, err := valueWeather.NewTemperature(23)
		if err != nil {
			log.Fatal(err)
		}
		forecast := entityWeather.NewForecast(description, *humidity, *temperature)
		return entityWeather.NewReport(createdAt, *id, *location, forecast)
	})()

	if err := ntfyProvider.Notify(user, report); err != nil {
		log.Fatal(err)
	}
}
