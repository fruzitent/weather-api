package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

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

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments")
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	switch os.Args[1] {
	case CMD_DAEMON:
		daemonCmd := flag.NewFlagSet(CMD_DAEMON, flag.ExitOnError)
		daemonCmd.Parse(os.Args[2:])

		_, err := sqlite.Open(ctx)
		if err != nil {
			log.Fatal(err)
		}

		subscriptionService := primary.NewSubscriptionService()
		weatherService := primary.NewWeatherService()

		http.ServeHTTP(addr, subscriptionService, weatherService)

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(os.Args[2:])

		if err := http.IsHealthy(addr); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("invalid subcommand %s\n", os.Args[1])
	}
}
