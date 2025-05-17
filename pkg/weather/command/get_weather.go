package command

type GetWeather struct {
	City string // City name for weather forecast
}

type GetWeatherRes struct {
	Temperature float64 // Current temperature
	Humidity    int     // Current humidity percentage
	Description string  // Weather description
}
