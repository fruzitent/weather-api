package query

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/decorator"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/weather/port"
)

type Current struct {
	Location value.Location
}

type CurrentHandler decorator.QueryHandler[Current, entity.Forecast]

type currentHandler struct {
	WeatherProvider port.Provider
}

var _ CurrentHandler = (*currentHandler)(nil)

func NewCurrentHandler(weatherProvider port.Provider) *currentHandler {
	return &currentHandler{weatherProvider}
}

func (h *currentHandler) Handle(ctx context.Context, query Current) (entity.Forecast, error) {
	res, err := h.WeatherProvider.GetCurrentWeather(query.Location)
	return *res, err
}
