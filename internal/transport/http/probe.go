package http

import (
	"fmt"
	"log"
	"net/http"
)

type probe struct{}

func NewProbeController(mux *http.ServeMux) *probe {
	c := &probe{}
	mux.HandleFunc("GET /-/healthy", c.healthy)
	mux.HandleFunc("GET /-/ready", c.ready)
	return c
}

func (c *probe) healthy(w http.ResponseWriter, r *http.Request) {
	log.Printf("Healthy: %s\n", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}

func (c *probe) ready(w http.ResponseWriter, r *http.Request) {
	log.Printf("Ready: %s\n", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}

func IsHealthy(addr string) error {
	_, err := http.Get(fmt.Sprintf("http://%s/-/healthy", addr))
	if err != nil {
		return err
	}
	return nil
}
