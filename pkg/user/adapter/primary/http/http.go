package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	valueShared "git.fruzit.pp.ua/weather/api/internal/shared/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/user/adapter/primary/http/oapi_gen"
	"git.fruzit.pp.ua/weather/api/pkg/user/core"
	"git.fruzit.pp.ua/weather/api/pkg/user/core/command"
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/value"
	valueWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
	"github.com/oklog/ulid/v2"
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

func (t *Http) ConfirmSubscription(ctx context.Context, request oapi_gen.ConfirmSubscriptionRequestObject) (oapi_gen.ConfirmSubscriptionResponseObject, error) {
	return &oapi_gen.ConfirmSubscription200Response{}, nil
}

// TODO: DOES NOT HANDLE 406 https://github.com/oapi-codegen/oapi-codegen/issues/736
func (t *Http) Subscribe(ctx context.Context, request oapi_gen.SubscribeRequestObject) (oapi_gen.SubscribeResponseObject, error) {
	cmd := &command.Subscribe{}

	// Generate ID
	id, err := valueShared.NewId(ulid.Make().String())
	if err != nil {
		return nil, err
	}

	var reqCity, reqEmail, reqFrequency string
	if request.FormdataBody != nil {
		reqCity = request.FormdataBody.City
		reqEmail = request.FormdataBody.Email
		reqFrequency = string(request.FormdataBody.Frequency)
	}
	if request.JSONBody != nil {
		reqCity = request.JSONBody.City
		reqEmail = request.JSONBody.Email
		reqFrequency = string(request.JSONBody.Frequency)
	}

	frequency, err := value.NewFrequency(reqFrequency)
	if err != nil {
		return nil, err
	}
	cmd.Frequency = *frequency

	location, err := valueWeather.NewLocation(reqCity)
	if err != nil {
		return nil, err
	}
	cmd.Location = *location

	mail, err := valueShared.NewMail(reqEmail)
	if err != nil {
		return nil, err
	}
	cmd.User = entity.NewUser(*id, *mail)

	if err := t.app.Command.Subscribe.Handle(ctx, *cmd); err != nil {
		return nil, err
	}

	return &oapi_gen.Subscribe200Response{}, nil
}

func (t *Http) Unsubscribe(ctx context.Context, request oapi_gen.UnsubscribeRequestObject) (oapi_gen.UnsubscribeResponseObject, error) {
	return &oapi_gen.Unsubscribe200Response{}, nil
}
