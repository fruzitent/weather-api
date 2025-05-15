package config

import (
	"flag"
)

type Config struct {
	Http Http
}

type Http struct {
	Host string
	Port int
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")
	flag.Parse()

	return config
}
