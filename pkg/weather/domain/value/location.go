package value

import "fmt"

type Location struct {
	City string
}

func NewLocation(city string) (*Location, error) {
	if city == "" {
		return nil, fmt.Errorf("missing parameter city")
	}
	return &Location{city}, nil
}
