package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"

	"git.fruzit.pp.ua/weather/api/internal/config"
	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/internal/lib/sqlite"
	httpProbe "git.fruzit.pp.ua/weather/api/pkg/probe/adapter/primary/http"
	httpUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/primary/http"
	"git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/smtp"
	sqliteUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/sqlite"
	coreUser "git.fruzit.pp.ua/weather/api/pkg/user/core"
	"git.fruzit.pp.ua/weather/api/pkg/user/core/command"
	httpWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/http"
	"git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/weatherapi"
	sqliteWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/secondary/sqlite"
	coreWeather "git.fruzit.pp.ua/weather/api/pkg/weather/core"
	queryWeather "git.fruzit.pp.ua/weather/api/pkg/weather/core/query"
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

		notification := &smtp.Smtp{
			Config: cfgDaemon.Smtp,
		}

		storageSqlite := sqliteUser.NewSqlite(db)

		providerWeather, err := weatherapi.NewWeatherapi(cfgDaemon.Weatherapi)
		if err != nil {
			log.Fatalf("[weatherapi] %s", err.Error())
		}

		appUser := coreUser.App{
			Command: coreUser.Command{
				Subscribe: command.NewSubscribeHandler(notification, storageSqlite),
			},
			Query: coreUser.Query{},
		}

		appWeather := coreWeather.App{
			Command: coreWeather.Command{},
			Query: coreWeather.Query{
				Current: queryWeather.NewCurrentHandler(providerWeather),
			},
		}

		mux := http.NewServeMux()
		_ = httpProbe.New(mux)
		_ = httpUser.New(mux, &appUser)
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
