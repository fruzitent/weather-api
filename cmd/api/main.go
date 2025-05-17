package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/repo/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/repo/weatherapi"
	serviceProbe "git.fruzit.pp.ua/weather/api/pkg/probe/service/primary"
	transportProbe "git.fruzit.pp.ua/weather/api/pkg/probe/transport/http"
	repoSubscription "git.fruzit.pp.ua/weather/api/pkg/subscription/repo/sqlite"
	serviceSubscription "git.fruzit.pp.ua/weather/api/pkg/subscription/service/primary"
	trasportSubscription "git.fruzit.pp.ua/weather/api/pkg/subscription/transport/http"
	repoWeather "git.fruzit.pp.ua/weather/api/pkg/weather/repo/sqlite"
	serviceWeather "git.fruzit.pp.ua/weather/api/pkg/weather/service/primary"
	transportWeather "git.fruzit.pp.ua/weather/api/pkg/weather/transport/http"
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

		db, err := sqlite.Open(ctx, config.Sqlite.DataSourceName)
		if err != nil {
			log.Fatal(err)
		}
		repoSubscription := repoSubscription.New(db)
		repoWeather := repoWeather.New(db)

		providerWeather, err := weatherapi.NewWeatherAPI(config.WeatherApi.Secret)
		if err != nil {
			log.Fatal(err)
		}

		serviceProbe := serviceProbe.New()
		serviceSubscription := serviceSubscription.New(repoSubscription)
		serviceWeather := serviceWeather.New(repoWeather, providerWeather)

		mux := http.NewServeMux()
		_ = transportProbe.New(mux, serviceProbe)
		_ = transportWeather.New(mux, serviceWeather)
		_ = trasportSubscription.New(mux, serviceSubscription)
		log.Printf("Listening on http://%s", addr)
		log.Fatal(http.ListenAndServe(addr, mux))

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

		if err := transportProbe.IsHealthy(addr); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("invalid subcommand %s\n", args[0])
	}
}
