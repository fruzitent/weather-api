package primary

import (
	"git.fruzit.pp.ua/weather/api/internal/command"
	"git.fruzit.pp.ua/weather/api/internal/repo"
	"git.fruzit.pp.ua/weather/api/internal/service"
)

type weather struct {
	repo repo.IWeather
}

func NewWeatherService(repo repo.IWeather) service.IWeather {
	return &weather{repo}
}

func (s *weather) GetWeather(c *command.GetWeather) (*command.GetWeatherRes, error) {
	return &command.GetWeatherRes{
		Temperature: 0,
		Humidity:    0,
		Description: "string",
	}, nil
}
