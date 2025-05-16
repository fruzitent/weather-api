package http

import (
	"fmt"
	"net/http"

	"git.fruzit.pp.ua/weather/api/pkg/probe/service"
)

type Transport struct {
	service service.IService
}

func New(mux *http.ServeMux, service service.IService) *Transport {
	c := &Transport{service}
	mux.HandleFunc("GET /-/healthy", c.getHealthy)
	mux.HandleFunc("GET /-/ready", c.getReady)
	return c
}

func (t *Transport) getHealthy(w http.ResponseWriter, r *http.Request) {
	if err := t.service.IsHealthy(); err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.WriteHeader(http.StatusOK)
}

func (t *Transport) getReady(w http.ResponseWriter, r *http.Request) {
	if err := t.service.IsReady(); err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.WriteHeader(http.StatusOK)
}

func IsHealthy(addr string) error {
	_, err := http.Get(fmt.Sprintf("http://%s/-/healthy", addr))
	if err != nil {
		return err
	}
	return nil
}
