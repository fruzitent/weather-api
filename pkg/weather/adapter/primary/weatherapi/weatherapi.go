package weatherapi

import (
	"git.fruzit.pp.ua/weather/api/internal/lib/weatherapi"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/weather/port"
)

type Weatherapi struct {
	api    *weatherapi.Weatherapi
	config *weatherapi.Config
}

var _ port.Provider = (*Weatherapi)(nil)

func NewWeatherapi(config *weatherapi.Config) (*Weatherapi, error) {
	api, err := weatherapi.NewWeatherapi(config.Secret)
	if err != nil {
		return nil, err
	}
	return &Weatherapi{api, config}, nil
}

func (a *Weatherapi) GetCurrentWeather(location value.Location) (*entity.Forecast, error) {
	res, err := a.api.Realtime(&weatherapi.RealtimeReq{
		Q: location.City,
	})
	if err != nil {
		return nil, err
	}

	humidity, err := value.NewHumidity(float64(res.Current.Humidity) / 100)
	if err != nil {
		return nil, err
	}

	temperature, err := value.NewTemperature(float64(res.Current.TempC))
	if err != nil {
		return nil, err
	}

	forecast := entity.NewForecast(res.Current.Condition.Text, *humidity, *temperature)
	return &forecast, nil
}
