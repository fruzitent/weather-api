package http

import (
	"context"
	"net/http"

	"git.fruzit.pp.ua/weather/api/pkg/weather/command"
	"git.fruzit.pp.ua/weather/api/pkg/weather/service"
	"git.fruzit.pp.ua/weather/api/pkg/weather/transport/http/openapi"
)

type Transport struct {
	service service.IService
}

func New(mux *http.ServeMux, service service.IService) *Transport {
	c := &Transport{service}
	_ = openapi.HandlerFromMux(openapi.NewStrictHandler(c, []openapi.StrictMiddlewareFunc{}), mux)
	return c
}

var _ openapi.StrictServerInterface = (*Transport)(nil)

func (t *Transport) GetWeather(ctx context.Context, request openapi.GetWeatherRequestObject) (openapi.GetWeatherResponseObject, error) {
	res, err := t.service.GetWeather(&command.GetWeather{
		City: request.Params.City,
	})
	if err != nil {
		return nil, err
	}
	humidity := float32(res.Humidity)
	temperature := float32(res.Temperature)
	return openapi.GetWeather200JSONResponse{
		Description: &res.Description,
		Humidity:    &humidity,
		Temperature: &temperature,
	}, nil
}
