package service

import "git.fruzit.pp.ua/weather/api/internal/command"

type IWeather interface {
	GetWeather(c *command.GetWeather) (*command.GetWeatherRes, error)
}
