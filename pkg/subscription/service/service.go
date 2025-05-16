package service

import "git.fruzit.pp.ua/weather/api/pkg/subscription/command"

type IService interface {
	ConfirmEmail(c *command.ConfirmEmail) (*command.ConfirmEmailRes, error)
	Subscribe(c *command.Subscribe) (*command.SubscribeRes, error)
	Unsubscribe(c *command.Unsubscribe) (*command.UnsubscribeRes, error)
}
