package http

import (
	"log"
	"net/http"
)

func ServeHTTP(addr string) {
	mux := http.NewServeMux()

	_ = NewProbeController(mux)
	_ = NewSubscriptionController(mux)
	_ = NewWeatherController(mux)

	log.Fatal(http.ListenAndServe(addr, mux))
}
