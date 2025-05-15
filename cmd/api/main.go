package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/repo/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/service/primary"
	"git.fruzit.pp.ua/weather/api/internal/transport/http"
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

		subscriptionRepo := sqlite.NewSubscriptionRepo(db)
		weatherRepo := sqlite.NewWeatherRepo(db)

		subscriptionService := primary.NewSubscriptionService(subscriptionRepo)
		weatherService := primary.NewWeatherService(weatherRepo)

		http.ServeHTTP(addr, subscriptionService, weatherService)

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

		if err := http.IsHealthy(addr); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("invalid subcommand %s\n", args[0])
	}
}
