package primary

type probe struct{}

func New() *probe {
	return &probe{}
}

func (s *probe) IsHealthy() error {
	return nil
}

func (s *probe) IsReady() error {
	return nil
}
