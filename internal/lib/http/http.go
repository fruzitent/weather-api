package http

import (
	"log"
	"net/http"
)

type Config struct {
	Host string
	Port int
}

func NewServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func ListenAndServe(addr string, handler http.Handler) error {
	log.Printf("[http] Listening on %s\n", addr)
	return http.ListenAndServe(addr, handler)
}
