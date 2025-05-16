package command

type GetWeather struct {
	City string // City name for weather forecast
}

type GetWeatherRes struct {
	Temperature float32 // Current temperature
	Humidity    float32 // Current humidity percentage
	Description string  // Weather description
}
