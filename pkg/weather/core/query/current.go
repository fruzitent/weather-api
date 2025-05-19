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
	Meteo port.Meteo
}

var _ CurrentHandler = (*currentHandler)(nil)

func NewCurrentHandler(meteo port.Meteo) *currentHandler {
	return &currentHandler{meteo}
}

func (h *currentHandler) Handle(ctx context.Context, query Current) (entity.Forecast, error) {
	res, err := h.Meteo.GetCurrentWeather(query.Location)
	return *res, err
}
