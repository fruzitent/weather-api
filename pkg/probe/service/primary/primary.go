package primary

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) IsHealthy() error {
	return nil
}

func (s *Service) IsReady() error {
	return nil
}
