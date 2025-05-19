package config

import (
	"flag"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/internal/lib/smtp"
	"git.fruzit.pp.ua/weather/api/internal/lib/sqlite"
	"git.fruzit.pp.ua/weather/api/internal/lib/weatherapi"
)

type RootConfig struct {
	Http *http.Config
}

func NewRootConfig(fs *flag.FlagSet, args []string) (*RootConfig, error) {
	config := &RootConfig{}
	config.Http = http.NewConfig(fs)
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	return config, nil
}

type DaemonConfig struct {
	Smtp       *smtp.Config
	Sqlite     *sqlite.Config
	Weatherapi *weatherapi.Config
}

func NewDaemonConfig(fs *flag.FlagSet, args []string) (*DaemonConfig, error) {
	config := &DaemonConfig{}
	config.Smtp = smtp.NewConfig(fs)
	config.Sqlite = sqlite.NewConfig(fs)
	config.Weatherapi = weatherapi.NewConfig(fs)
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	return config, nil
}

func (config *DaemonConfig) Load() (err error) {
	if config.Smtp.Password, err = loadSecret(config.Smtp.Password, "smtp.password"); err != nil {
		return err
	}
	if config.Weatherapi.Secret, err = loadSecret(config.Weatherapi.Secret, "weatherapi.secret"); err != nil {
		return err
	}
	return nil
}
