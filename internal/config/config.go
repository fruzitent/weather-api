package config

import (
	"flag"
)

type Config struct {
	Http   Http
	Sqlite Sqlite
}

type Http struct {
	Host string
	Port int
}

type Sqlite struct {
	DataSourceName string
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")
	flag.StringVar(&config.Sqlite.DataSourceName, "sqlite.dataSourceName", "db.sqlite", "")
	flag.Parse()

	return config
}
