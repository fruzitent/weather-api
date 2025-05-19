package entity

import "git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"

type Forecast struct {
	Description string
	Humidity    value.Humidity
	Temperature value.Temperature
}

func NewForecast(
	description string,
	humidity value.Humidity,
	temperature value.Temperature,
) Forecast {
	return Forecast{description, humidity, temperature}
}
