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
	smtpUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/smtp"
	sqliteUser "git.fruzit.pp.ua/weather/api/pkg/user/adapter/secondary/sqlite"
	coreUser "git.fruzit.pp.ua/weather/api/pkg/user/core"
	commandUser "git.fruzit.pp.ua/weather/api/pkg/user/core/command"

	httpWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/http"
	weatherapiWeather "git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/weatherapi"
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
		log.Fatalf("[flag]: %s", err.Error())
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("[flag] not enough arguments")
	}

	if err := func() error {
		switch args[0] {
		case CMD_DAEMON:
			return runDaemon(ctx, cfgRoot, args)
		case CMD_HEALTH:
			return runHealth(ctx, cfgRoot, args)
		default:
			return fmt.Errorf("[flag] invalid subcommand %s\n", args[0])
		}
	}(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func runDaemon(ctx context.Context, cfgRoot *config.RootConfig, args []string) error {
	fs := flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)

	cfg, err := config.NewDaemonConfig(fs, args[1:])
	if err != nil {
		return err
	}
	if err := cfg.Load(); err != nil {
		return err
	}

	schema := slices.Concat(
		sqliteUser.Schema,
		sqliteWeather.Schema,
	)
	db, err := sqlite.Open(ctx, cfg.Sqlite.Source, schema)
	if err != nil {
		return err
	}

	notification := smtpUser.NewSmtp(cfg.Smtp)
	storage := sqliteUser.NewSqlite(db)

	appUser := coreUser.App{
		Command: coreUser.Command{
			Subscribe: commandUser.NewSubscribeHandler(notification, storage),
		},
		Query: coreUser.Query{},
	}

	meteo, err := weatherapiWeather.NewWeatherapi(cfg.Weatherapi)
	if err != nil {
		return err
	}

	appWeather := coreWeather.App{
		Command: coreWeather.Command{},
		Query: coreWeather.Query{
			Current: queryWeather.NewCurrentHandler(meteo),
		},
	}

	mux := http.NewServeMux()
	_ = httpProbe.New(mux)
	_ = httpUser.New(mux, &appUser)
	_ = httpWeather.New(mux, &appWeather)
	addr := fmt.Sprintf("%s:%d", cfgRoot.Http.Host, cfgRoot.Http.Port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("[http] %s", err.Error())
	}

	return nil
}

func runHealth(ctx context.Context, cfgRoot *config.RootConfig, args []string) error {
	_ = flag.NewFlagSet(CMD_HEALTH, flag.ExitOnError)

	addr := fmt.Sprintf("%s:%d", cfgRoot.Http.Host, cfgRoot.Http.Port)
	if err := httpProbe.IsHealthy(addr); err != nil {
		return err
	}

	return nil
}
