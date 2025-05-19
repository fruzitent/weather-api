package port

import (
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

type Meteo interface {
	GetCurrentWeather(location value.Location) (*entity.Forecast, error)
}
