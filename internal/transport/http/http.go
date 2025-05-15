package http

import "net/http"

func NewServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
