package config

import (
	"flag"
)

type Config struct {
	Http       Http
	Sqlite     Sqlite
	WeatherApi WeatherApi
}

type Http struct {
	Host string
	Port int
}

type Sqlite struct {
	DataSourceName string
}

type WeatherApi struct {
	SecretFile string
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")
	flag.StringVar(&config.Sqlite.DataSourceName, "sqlite.dataSourceName", "db.sqlite3", "")
	flag.StringVar(&config.WeatherApi.SecretFile, "weatherApi.secretFile", "weatherapi-token", "")
	flag.Parse()

	return config
}
