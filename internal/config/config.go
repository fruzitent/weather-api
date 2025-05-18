package config

import (
	"flag"
	"net/mail"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/internal/lib/smtp"
	"git.fruzit.pp.ua/weather/api/internal/lib/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/lib/weatherapi"
)

type Config struct {
	Http       http.Config
	Smtp       smtp.Config
	Sqlite     sqlite.Config
	Weatherapi weatherapi.Config
}

func NewConfig() (*Config, error) {
	config := &Config{}
	var err error

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")

	flag.Func("smtp.from", "", func(s string) error {
		addr, err := mail.ParseAddress(s)
		if err != nil {
			return err
		}
		config.Smtp.From = *addr
		return nil
	})
	flag.StringVar(&config.Smtp.Host, "smtp.host", "", "")
	flag.StringVar(&config.Smtp.Password, "smtp.password", "", "")
	flag.IntVar(&config.Smtp.Port, "smtp.port", 0, "")
	flag.StringVar(&config.Smtp.Username, "smtp.username", "", "")

	flag.StringVar(&config.Sqlite.DataSourceName, "sqlite.dataSourceName", "db.sqlite3", "")

	flag.StringVar(&config.Weatherapi.Secret, "weatherapi.secret", "", "")

	flag.Parse()

	if config.Smtp.Password, err = loadSecret(config.Smtp.Password, "smtp.password"); err != nil {
		return nil, err
	}

	if config.Weatherapi.Secret, err = loadSecret(config.Weatherapi.Secret, "weatherapi.secret"); err != nil {
		return nil, err
	}

	return config, nil
}
