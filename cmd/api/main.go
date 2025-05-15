package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"git.fruzit.pp.ua/weather/api/internal/config"
	repo "git.fruzit.pp.ua/weather/api/internal/repo/sqlite"
	service "git.fruzit.pp.ua/weather/api/internal/service/primary"
	transport "git.fruzit.pp.ua/weather/api/internal/transport/http"
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

		db, err := repo.Open(ctx)
		if err != nil {
			log.Fatal(err)
		}
		subscriptionRepo := repo.NewSubscriptionRepo(db)
		weatherRepo := repo.NewWeatherRepo(db)

		probeService := service.NewProbeService()
		subscriptionService := service.NewSubscriptionService(subscriptionRepo)
		weatherService := service.NewWeatherService(weatherRepo)

		mux := transport.NewServeMux()
		_ = transport.NewProbeController(mux, probeService)
		_ = transport.NewSubscriptionController(mux, subscriptionService)
		_ = transport.NewWeatherController(mux, weatherService)
		log.Fatal(transport.ListenAndServe(addr, mux))

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

		if err := transport.IsHealthy(addr); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("invalid subcommand %s\n", args[0])
	}
}
