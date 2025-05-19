package command

import (
	"context"

	"git.fruzit.pp.ua/weather/api/internal/decorator"
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/user/port"
	valueWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

type Subscribe struct {
	Frequency value.Frequency
	Location  valueWeather.Location
	User      entity.User
}

type SubscribeHandler decorator.CommandHandler[Subscribe]

type subscribeHandler struct {
	Notification port.Notification
	Storage      port.Storage
}

var _ SubscribeHandler = (*subscribeHandler)(nil)

func NewSubscribeHandler(notification port.Notification, storage port.Storage) *subscribeHandler {
	return &subscribeHandler{notification, storage}
}

func (h *subscribeHandler) Handle(ctx context.Context, cmd Subscribe) error {
	if err := h.Notification.SendConfirmation(cmd.User); err != nil {
		return err
	}
	if err := h.Storage.SaveUser(cmd.User); err != nil {
		return err
	}
	return nil
}
