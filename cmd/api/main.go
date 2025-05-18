package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	httpUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/driving/http"
	httpWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/driving/http"
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

		mux := http.NewServeMux()
		_ = httpUser.New(mux)
		_ = httpWeather.New(mux)
		log.Fatal(http.ListenAndServe(addr, mux))

	case CMD_HEALTH:
		healthCmd := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)
		healthCmd.Parse(args[1:])

	default:
		log.Fatalf("invalid subcommand %s\n", args[0])
	}
}
