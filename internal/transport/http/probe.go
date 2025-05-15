package http

import (
	"fmt"
	"log"
	"net/http"
)

type probeController struct{}

func NewProbeController(mux *http.ServeMux) *probeController {
	c := &probeController{}
	mux.HandleFunc("GET /-/healthy", c.healthy)
	mux.HandleFunc("GET /-/ready", c.ready)
	return c
}

func (c *probeController) healthy(w http.ResponseWriter, r *http.Request) {
	log.Printf("Healthy: %s\n", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}

func (c *probeController) ready(w http.ResponseWriter, r *http.Request) {
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
