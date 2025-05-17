package config

import (
	"flag"

	"git.fruzit.pp.ua/weather/api/internal/repo/smtp"
	"git.fruzit.pp.ua/weather/api/internal/repo/weatherapi"
)

type Config struct {
	Http       Http
	Smtp       smtp.Config
	Sqlite     Sqlite
	WeatherApi weatherapi.Config
}

type Http struct {
	Host string
	Port int
}

type Sqlite struct {
	DataSourceName string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	var err error

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")

	flag.StringVar(&config.Smtp.From, "smtp.from", "", "")
	flag.StringVar(&config.Smtp.Host, "smtp.host", "", "")
	flag.StringVar(&config.Smtp.Password, "smtp.password", "", "")
	flag.IntVar(&config.Smtp.Port, "smtp.port", 0, "")
	flag.StringVar(&config.Smtp.Username, "smtp.username", "", "")

	flag.StringVar(&config.Sqlite.DataSourceName, "sqlite.dataSourceName", "db.sqlite3", "")

	flag.StringVar(&config.WeatherApi.Secret, "weatherApi.secret", "", "")

	flag.Parse()

	if config.Smtp.Password, err = loadSecret(config.Smtp.Password, "smtp.password"); err != nil {
		return nil, err
	}

	if config.WeatherApi.Secret, err = loadSecret(config.WeatherApi.Secret, "weatherApi.secret"); err != nil {
		return nil, err
	}

	return config, nil
}
