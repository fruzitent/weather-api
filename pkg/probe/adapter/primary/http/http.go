package http

import (
	"fmt"
	"net/http"
)

type Http struct{}

func New(mux *http.ServeMux) *Http {
	c := &Http{}
	mux.HandleFunc("GET /-/healthy", c.getHealthy)
	mux.HandleFunc("GET /-/ready", c.getReady)
	return c
}

func (t *Http) getHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (t *Http) getReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func IsHealthy(addr string) error {
	_, err := http.Get(fmt.Sprintf("http://%s/-/healthy", addr))
	if err != nil {
		return err
	}
	return nil
}
