package primary

import (
	"git.fruzit.pp.ua/weather/api/pkg/subscription/command"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/repo"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/service"
)

type Service struct {
	repo repo.IRepo
}

func New(repo repo.IRepo) service.IService {
	return &Service{repo}
}

func (s *Service) ConfirmEmail(c *command.ConfirmEmail) (*command.ConfirmEmailRes, error) {
	return nil, nil
}

func (s *Service) Subscribe(c *command.Subscribe) (*command.SubscribeRes, error) {
	return nil, nil
}

func (s *Service) Unsubscribe(c *command.Unsubscribe) (*command.UnsubscribeRes, error) {
	return nil, nil
}
