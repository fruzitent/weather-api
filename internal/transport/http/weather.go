package http

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"git.fruzit.pp.ua/weather/api/internal/command"
)

type weatherController struct{}

func NewWeatherController(mux *http.ServeMux) *weatherController {
	c := &weatherController{}
	mux.HandleFunc("GET /weather", c.getWeather)
	return c
}

func (c *weatherController) getWeather(w http.ResponseWriter, r *http.Request) {
	v, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	city := v.Get("city")
	if city == "" {
		http.Error(w, "missing parameter (city)", http.StatusBadRequest)
		return
	}
	log.Printf("GetWeather: city=%s\n", city)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(command.GetWeatherRes{
		Temperature: 0,
		Humidity:    0,
		Description: "string",
	})
}
