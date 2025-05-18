package http

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/lib/http"
	"git.fruzit.pp.ua/weather/api/pkg/user/adapter/driving/http/oapi_gen"
)

type Http struct{}

var _ oapi_gen.StrictServerInterface = (*Http)(nil)

func New(mux *http.ServeMux) *Http {
	a := &Http{}
	_ = oapi_gen.HandlerFromMux(oapi_gen.NewStrictHandler(a, []oapi_gen.StrictMiddlewareFunc{}), mux)
	return a
}

func (t *Http) ConfirmSubscription(ctx context.Context, request oapi_gen.ConfirmSubscriptionRequestObject) (oapi_gen.ConfirmSubscriptionResponseObject, error) {
	return &oapi_gen.ConfirmSubscription200Response{}, nil
}

// TODO: DOES NOT HANDLE 406 https://github.com/oapi-codegen/oapi-codegen/issues/736
func (t *Http) Subscribe(ctx context.Context, request oapi_gen.SubscribeRequestObject) (oapi_gen.SubscribeResponseObject, error) {
	if request.FormdataBody != nil {
	}
	if request.JSONBody != nil {
	}
	return &oapi_gen.Subscribe200Response{}, nil
}

func (t *Http) Unsubscribe(ctx context.Context, request oapi_gen.UnsubscribeRequestObject) (oapi_gen.UnsubscribeResponseObject, error) {
	return &oapi_gen.Unsubscribe200Response{}, nil
}
