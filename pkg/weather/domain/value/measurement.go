package value

import "fmt"

type Measurement struct {
	Description string
	Humidity    float64
	Temperature float64
}

func NewMeasurement(description string, humidity float64, temperature float64) (*Measurement, error) {
	if humidity < 0 || humidity > 1 {
		return nil, fmt.Errorf("humidity must be in range [0; 1] %f", humidity)
	}
	if temperature < -273.15 {
		return nil, fmt.Errorf("temperature below absolute zero %f", temperature)
	}
	return &Measurement{description, humidity, temperature}, nil
}
