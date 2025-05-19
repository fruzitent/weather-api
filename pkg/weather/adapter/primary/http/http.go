package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/pkg/weather/adapter/primary/http/oapi_gen"
	"git.fruzit.pp.ua/weather/api/pkg/weather/core"
	"git.fruzit.pp.ua/weather/api/pkg/weather/core/query"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

type Http struct {
	app *core.App
}

var _ oapi_gen.StrictServerInterface = (*Http)(nil)

func New(mux *http.ServeMux, app *core.App) *Http {
	a := &Http{app}
	_ = oapi_gen.HandlerFromMux(oapi_gen.NewStrictHandler(a, []oapi_gen.StrictMiddlewareFunc{}), mux)
	return a
}

func (a *Http) GetWeather(ctx context.Context, request oapi_gen.GetWeatherRequestObject) (oapi_gen.GetWeatherResponseObject, error) {
	location, err := value.NewLocation(request.Params.City)
	if err != nil {
		return nil, err
	}

	res, err := a.app.Query.Current.Handle(ctx, query.Current{
		Location: *location,
	})
	if err != nil {
		return nil, err
	}
	humidity := float32(res.Humidity.Percentage)
	temperature := float32(res.Temperature.Celcius)
	return &oapi_gen.GetWeather200JSONResponse{
		Description: &res.Description,
		Humidity:    &humidity,
		Temperature: &temperature,
	}, nil
}
