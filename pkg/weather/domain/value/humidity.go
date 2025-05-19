package value

import "fmt"

type Humidity struct {
	Percentage float64
}

func NewHumidity(percentage float64) (*Humidity, error) {
	if percentage < 0 || percentage > 1 {
		return nil, fmt.Errorf("humidity must be in range [0; 1] %f", percentage)
	}
	return &Humidity{percentage}, nil
}
