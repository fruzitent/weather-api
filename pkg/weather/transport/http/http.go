package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"git.fruzit.pp.ua/weather/api/pkg/weather/command"
	"git.fruzit.pp.ua/weather/api/pkg/weather/service"
)

type Transport struct {
	service service.IService
}

func New(mux *http.ServeMux, service service.IService) *Transport {
	c := &Transport{service}
	mux.HandleFunc("GET /weather", c.getWeather)
	return c
}

func (t *Transport) getWeather(w http.ResponseWriter, r *http.Request) {
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

	res, err := t.service.GetWeather(&command.GetWeather{
		City: city,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
