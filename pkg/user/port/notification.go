package port

import (
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	entityWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
)

type Notification interface {
	SendWeatherReport(user entity.User, report entityWeather.Report) error
}
