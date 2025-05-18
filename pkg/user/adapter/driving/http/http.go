package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/pkg/user/adapter/driving/http/openapi"
)

type Http struct{}

var _ openapi.StrictServerInterface = (*Http)(nil)

func New(mux *http.ServeMux) *Http {
	a := &Http{}
	_ = openapi.HandlerFromMux(openapi.NewStrictHandler(a, []openapi.StrictMiddlewareFunc{}), mux)
	return a
}

func (t *Http) ConfirmSubscription(ctx context.Context, request openapi.ConfirmSubscriptionRequestObject) (openapi.ConfirmSubscriptionResponseObject, error) {
	return &openapi.ConfirmSubscription200Response{}, nil
}

// TODO: DOES NOT HANDLE 406 https://github.com/oapi-codegen/oapi-codegen/issues/736
func (t *Http) Subscribe(ctx context.Context, request openapi.SubscribeRequestObject) (openapi.SubscribeResponseObject, error) {
	if request.FormdataBody != nil {
	}
	if request.JSONBody != nil {
	}
	return &openapi.Subscribe200Response{}, nil
}

func (t *Http) Unsubscribe(ctx context.Context, request openapi.UnsubscribeRequestObject) (openapi.UnsubscribeResponseObject, error) {
	return &openapi.Unsubscribe200Response{}, nil
}
