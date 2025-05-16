package primary

import (
	"git.fruzit.pp.ua/weather/api/pkg/subscription/command"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/repo"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/service"
)

type subscription struct {
	repo repo.ISubscription
}

func New(repo repo.ISubscription) service.ISubscription {
	return &subscription{repo}
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
