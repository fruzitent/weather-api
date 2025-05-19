package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
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

	cfgRoot, err := config.NewRootConfig(flag.CommandLine, os.Args[1:])
	if err != nil {
		log.Fatalf("[flag/root] %s", err.Error())
	}
	addr := fmt.Sprintf("%s:%d", cfgRoot.Http.Host, cfgRoot.Http.Port)

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("[flag] not enough arguments")
	}

	switch args[0] {
	case CMD_DAEMON:
		daemonCmd := flag.NewFlagSet(CMD_DAEMON, flag.ExitOnError)
		cfgDaemon, err := config.NewDaemonConfig(daemonCmd, args[1:])
		if err != nil {
			log.Fatalf("[flag] %s", err.Error())
		}
		if err := cfgDaemon.Load(); err != nil {
			log.Fatalf("[flag] %s", err.Error())
		}

		schema := slices.Concat(
			sqliteUser.Schema,
			sqliteWeather.Schema,
		)
		db, err := sqlite.Open(ctx, cfgDaemon.Sqlite.Source, schema)
		if err != nil {
			log.Fatalf("[sqlite] %s", err.Error())
		}
		_ = db

		if err := notifications(cfgDaemon); err != nil {
			log.Fatalf("[notification] %s", err.Error())
		}

		providerWeather, err := weatherapi.NewWeatherapi(cfgDaemon.Weatherapi)
		if err != nil {
			log.Fatalf("[weatherapi] %s", err.Error())
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
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatalf("[http] %s", err.Error())
		}

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

		if err := httpProbe.IsHealthy(addr); err != nil {
			log.Fatalf("[health] %s", err.Error())
		}

	default:
		log.Fatalf("[flag] invalid subcommand %s\n", args[0])
	}
}

func notifications(config *config.DaemonConfig) error {
	notification := smtp.Smtp{
		Config: config.Smtp,
	}

	user, err := (func() (entityUser.User, error) {
		id, err := value.NewId("hi")
		if err != nil {
			return entityUser.User{}, err
		}
		mail, err := value.NewMail("fruzit@fruzit.pp.ua")
		if err != nil {
			return entityUser.User{}, err
		}
		return entityUser.NewUser(*id, *mail), nil
	})()
	if err != nil {
		return err
	}

	report, err := (func() (entityWeather.Report, error) {
		createdAt := time.Now().Unix()
		id, err := value.NewId("test-id")
		if err != nil {
			return entityWeather.Report{}, err
		}
		location, err := valueWeather.NewLocation("Kyiv")
		if err != nil {
			return entityWeather.Report{}, err
		}
		description := "Cloudy"
		humidity, err := valueWeather.NewHumidity(0.5)
		if err != nil {
			return entityWeather.Report{}, err
		}
		temperature, err := valueWeather.NewTemperature(23)
		if err != nil {
			return entityWeather.Report{}, err
		}
		forecast := entityWeather.NewForecast(description, *humidity, *temperature)
		return entityWeather.NewReport(createdAt, *id, *location, forecast), nil
	})()
	if err != nil {
		return err
	}

	if err := notification.SendWeatherReport(user, report); err != nil {
		return err
	}

	return nil
}
