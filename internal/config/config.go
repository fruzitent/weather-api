package config

import "flag"

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
	Secret string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	var err error

	flag.StringVar(&config.Http.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Http.Port, "http.port", 8000, "")
	flag.StringVar(&config.Sqlite.DataSourceName, "sqlite.dataSourceName", "db.sqlite3", "")

	flag.StringVar(&config.WeatherApi.Secret, "weatherApi.secret", "", "")

	flag.Parse()

	if config.WeatherApi.Secret, err = loadSecret(config.WeatherApi.Secret, "weatherApi.secret"); err != nil {
		return nil, err
	}

	return config, nil
}
