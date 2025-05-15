package entity

type Weather struct {
	Temperature float32 `json:"temperature"` // Current temperature
	Humidity    float32 `json:"humidity"`    // Current humidity percentage
	Description string  `json:"description"` // Weather description
}
