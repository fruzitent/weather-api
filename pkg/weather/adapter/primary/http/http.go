package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/http/oapi_gen"
)

type Http struct{}

var _ oapi_gen.StrictServerInterface = (*Http)(nil)

func New(mux *http.ServeMux) *Http {
	a := &Http{}
	_ = oapi_gen.HandlerFromMux(oapi_gen.NewStrictHandler(a, []oapi_gen.StrictMiddlewareFunc{}), mux)
	return a
}

func (a *Http) GetWeather(ctx context.Context, request oapi_gen.GetWeatherRequestObject) (oapi_gen.GetWeatherResponseObject, error) {
	return &oapi_gen.GetWeather200JSONResponse{}, nil
}
