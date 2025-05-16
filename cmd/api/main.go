package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/repo/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/transport/http"
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

	config := config.NewConfig()
	addr := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("not enough arguments")
	}

	switch args[0] {
	case CMD_DAEMON:
		daemonCmd := flag.NewFlagSet(CMD_DAEMON, flag.ExitOnError)
		daemonCmd.Parse(args[1:])

		db, err := sqlite.Open(ctx)
		if err != nil {
			log.Fatal(err)
		}
		repoSubscription := repoSubscription.New(db)
		repoWeather := repoWeather.New(db)

		serviceProbe := serviceProbe.New()
		serviceSubscription := serviceSubscription.New(repoSubscription)
		serviceWeather := serviceWeather.New(repoWeather)

		mux := http.NewServeMux()
		_ = transportProbe.New(mux, serviceProbe)
		_ = trasportSubscription.New(mux, serviceSubscription)
		_ = transportWeather.New(mux, serviceWeather)
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
