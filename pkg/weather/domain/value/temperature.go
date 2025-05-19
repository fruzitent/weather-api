package value

import "fmt"

type Temperature struct {
	Celcius float64
}

func NewTemperature(celcius float64) (*Temperature, error) {
	if celcius < -273.15 {
		return nil, fmt.Errorf("celcius below absolute zero %f", celcius)
	}
	return &Temperature{celcius}, nil
}
