package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Weather struct {
	Temperature float32 `json:"temperature"` // Current temperature
	Humidity    float32 `json:"humidity"`    // Current humidity percentage
	Description string  `json:"description"` // Weather description
}

type Subscription struct {
	Email     string `json:"email"`     // Email address
	City      string `json:"city"`      // City for weather updates
	Frequency string `json:"frequency"` // Frequency of updates
	Confirmed bool   `json:"confirmed"` // Whether the subscription is confirmed
}

func GetConfirm(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	log.Printf("GetConfirm: token=%s", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetUnsubscribe(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	log.Printf("GetUnsubscribe: token=%s\n", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(Weather{
		Temperature: 0,
		Humidity:    0,
		Description: "string",
	})
}

func PostSubscribe(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	log.Printf("PostSubscribe: body=%s\n", body)

	var s Subscription
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /confirm/{token}", GetConfirm)
	mux.HandleFunc("GET /unsubscribe/{token}", GetUnsubscribe)
	mux.HandleFunc("GET /weather", GetWeather)
	mux.HandleFunc("POST /subscribe", PostSubscribe)

	mux.HandleFunc("GET /-/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("GET /-/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
