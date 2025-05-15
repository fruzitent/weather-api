package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"git.fruzit.pp.ua/weather/api/internal/command"
)

type subscriptionController struct{}

func NewSubscriptionController(mux *http.ServeMux) *subscriptionController {
	c := &subscriptionController{}
	mux.HandleFunc("GET /confirm/{token}", c.getConfirm)
	mux.HandleFunc("GET /unsubscribe/{token}", c.getUnsubscribe)
	mux.HandleFunc("POST /subscribe", c.postSubscribe)
	return c
}

func (c *subscriptionController) getConfirm(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	log.Printf("GetConfirm: token=%s", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *subscriptionController) getUnsubscribe(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	log.Printf("GetUnsubscribe: token=%s\n", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *subscriptionController) postSubscribe(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	log.Printf("PostSubscribe: body=%s\n", body)

	s := &command.Subscribe{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
