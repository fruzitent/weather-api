package primary

import (
	"git.fruzit.pp.ua/weather/api/internal/repo/weatherapi"
	"git.fruzit.pp.ua/weather/api/pkg/weather/command"
	"git.fruzit.pp.ua/weather/api/pkg/weather/repo"
	"git.fruzit.pp.ua/weather/api/pkg/weather/service"
)

type Service struct {
	repo     repo.IRepo
	provider *weatherapi.WeatherApi
}

func New(repo repo.IRepo, provider *weatherapi.WeatherApi) service.IService {
	return &Service{repo, provider}
}

func (s *Service) GetWeather(c *command.GetWeather) (*command.GetWeatherRes, error) {
	res, err := s.provider.Current(&weatherapi.CurrentReq{
		Q: c.City,
	})
	if err != nil {
		return nil, err
	}
	return &command.GetWeatherRes{
		Temperature: res.Current.TempC,
		Humidity:    res.Current.Humidity,
		Description: res.Current.Condition.Text,
	}, nil
}
