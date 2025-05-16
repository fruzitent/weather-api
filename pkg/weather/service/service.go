package service

import "git.fruzit.pp.ua/weather/api/pkg/weather/command"

type IWeather interface {
	GetWeather(c *command.GetWeather) (*command.GetWeatherRes, error)
}
