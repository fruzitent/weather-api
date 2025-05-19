package value

import "fmt"

type Id struct{ string }

func NewId(id string) (*Id, error) {
	if id == "" {
		return nil, fmt.Errorf("missing parameter id")
	}
	return &Id{id}, nil
}
