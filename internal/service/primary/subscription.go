package primary

import (
	"git.fruzit.pp.ua/weather/api/internal/command"
	"git.fruzit.pp.ua/weather/api/internal/service"
)

type subscription struct{}

func NewSubscriptionService() service.ISubscription {
	return &subscription{}
}

func (s *subscription) ConfirmEmail(c *command.ConfirmEmail) (*command.ConfirmEmailRes, error) {
	return nil, nil
}

func (s *subscription) Subscribe(c *command.Subscribe) (*command.SubscribeRes, error) {
	return nil, nil
}

func (s *subscription) Unsubscribe(c *command.Unsubscribe) (*command.UnsubscribeRes, error) {
	return nil, nil
}
