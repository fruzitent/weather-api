package value

import "fmt"

type Frequency struct {
	slug string
}

var (
	FrequencyUnknown = Frequency{""}
	FrequencyDaily   = Frequency{"daily"}
	FrequencyHourly  = Frequency{"hourly"}
)

func NewFrequency(slug string) (*Frequency, error) {
	switch slug {
	case FrequencyDaily.slug:
		return &FrequencyDaily, nil
	case FrequencyHourly.slug:
		return &FrequencyHourly, nil
	default:
		return &FrequencyUnknown, fmt.Errorf("unknown frequency %s", slug)
	}
}
