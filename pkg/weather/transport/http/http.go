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

	res, err := t.service.GetWeather(&command.GetWeather{
		City: v.Get("city"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
