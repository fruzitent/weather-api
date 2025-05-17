package http

import (
	"context"
	"net/http"

	"git.fruzit.pp.ua/weather/api/pkg/subscription/command"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/service"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/transport/http/openapi"
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

func (t *Transport) ConfirmSubscription(ctx context.Context, request openapi.ConfirmSubscriptionRequestObject) (openapi.ConfirmSubscriptionResponseObject, error) {
	_, err := t.service.ConfirmEmail(&command.ConfirmEmail{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}
	return &openapi.ConfirmSubscription200Response{}, nil
}

// TODO: DOES NOT HANDLE 406 https://github.com/oapi-codegen/oapi-codegen/issues/736
func (t *Transport) Subscribe(ctx context.Context, request openapi.SubscribeRequestObject) (openapi.SubscribeResponseObject, error) {
	cmd := &command.Subscribe{}

	if request.FormdataBody != nil {
		cmd.City = request.FormdataBody.City
		cmd.Email = request.FormdataBody.Email
		cmd.Frequency = string(request.FormdataBody.Frequency)
	}

	if request.JSONBody != nil {
		cmd.City = request.JSONBody.City
		cmd.Email = request.JSONBody.Email
		cmd.Frequency = string(request.JSONBody.Frequency)
	}

	_, err := t.service.Subscribe(cmd)
	if err != nil {
		return nil, err
	}

	return &openapi.Subscribe200Response{}, nil
}

func (t *Transport) Unsubscribe(ctx context.Context, request openapi.UnsubscribeRequestObject) (openapi.UnsubscribeResponseObject, error) {
	_, err := t.service.Unsubscribe(&command.Unsubscribe{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}
	return &openapi.Unsubscribe200Response{}, nil
}
