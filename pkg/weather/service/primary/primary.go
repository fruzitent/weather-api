package primary

import (
	"git.fruzit.pp.ua/weather/api/pkg/weather/command"
	"git.fruzit.pp.ua/weather/api/pkg/weather/repo"
	"git.fruzit.pp.ua/weather/api/pkg/weather/service"
)

type Service struct {
	repo repo.IRepo
}

func New(repo repo.IRepo) service.IService {
	return &Service{repo}
}

func (s *Service) GetWeather(c *command.GetWeather) (*command.GetWeatherRes, error) {
	return &command.GetWeatherRes{
		Temperature: 0,
		Humidity:    0,
		Description: "string",
	}, nil
}
