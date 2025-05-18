package decorator

import "context"

// Example:
//
//	package query
//
//	import (
//		"context"
//
//		"git.fruzit.pp.ua/weather/api/internal/decorator"
//	)
//
//	type MyQuery struct{}
//
//	type MyQueryHandler decorator.QueryHandler[MyQuery, string]
//
//	type myQueryHandler struct{}
//
//	var _ MyQueryHandler = (*myQueryHandler)(nil)
//
//	func NewMyQueryHandler() *myQueryHandler {
//		return &myQueryHandler{}
//	}
//
//	func (h *myQueryHandler) Handle(ctx context.Context, query MyQuery) (string, error) {
//		return "", nil
//	}
type QueryHandler[Query any, Return any] interface {
	Handle(ctx context.Context, query Query) (Return, error)
}
