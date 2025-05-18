package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/pkg/weather/adapter/driving/http/openapi"
)

type Http struct{}

var _ openapi.StrictServerInterface = (*Http)(nil)

func New(mux *http.ServeMux) *Http {
	a := &Http{}
	_ = openapi.HandlerFromMux(openapi.NewStrictHandler(a, []openapi.StrictMiddlewareFunc{}), mux)
	return a
}

func (a *Http) GetWeather(ctx context.Context, request openapi.GetWeatherRequestObject) (openapi.GetWeatherResponseObject, error) {
	return &openapi.GetWeather200JSONResponse{}, nil
}
