package decorator

import "context"

// Example:
//
//	package command
//
//	import (
//		"context"
//
//		"git.fruzit.pp.ua/weather/api/internal/decorator"
//	)
//
//	type MyCommand struct{}
//
//	type MyCommandHandler decorator.CommandHandler[MyCommand]
//
//	type myCommandHandler struct{}
//
//	var _ MyCommandHandler = (*myCommandHandler)(nil)
//
//	func NewMyCommandHandler() *myCommandHandler {
//		return &myCommandHandler{}
//	}
//
//	func (h *myCommandHandler) Handle(ctx context.Context, cmd MyCommand) error {
//		return nil
//	}
type CommandHandler[Command any] interface {
	Handle(ctx context.Context, cmd Command) error
}
