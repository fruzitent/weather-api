package http

import (
	"log"
	"net/http"

	"git.fruzit.pp.ua/weather/api/internal/service"
)

func ServeHTTP(addr string, subscription service.ISubscription, weather service.IWeather) {
	mux := http.NewServeMux()

	_ = NewProbeController(mux)
	_ = NewSubscriptionController(mux, subscription)
	_ = NewWeatherController(mux, weather)

	log.Fatal(http.ListenAndServe(addr, mux))
}
