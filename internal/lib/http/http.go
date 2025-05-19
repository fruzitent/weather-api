package http

import (
	"flag"
	"log"
	"net/http"
)

type Config struct {
	Host string
	Port int
}

func NewConfig(fs *flag.FlagSet) *Config {
	config := &Config{}
	flag.StringVar(&config.Host, "http.host", "[::]", "")
	flag.IntVar(&config.Port, "http.port", 8000, "")
	return config
}

type Handler = http.Handler
type ServeMux = http.ServeMux

func NewServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func ListenAndServe(addr string, handler http.Handler) error {
	log.Printf("[http] Listening on %s\n", addr)
	return http.ListenAndServe(addr, handler)
}
