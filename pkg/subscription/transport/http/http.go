package http

import (
	"encoding/json"
	"net/http"

	"git.fruzit.pp.ua/weather/api/pkg/subscription/command"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/service"
)

type subscription struct {
	service service.ISubscription
}

func New(mux *http.ServeMux, service service.ISubscription) *subscription {
	c := &subscription{service}
	mux.HandleFunc("GET /confirm/{token}", c.getConfirm)
	mux.HandleFunc("GET /unsubscribe/{token}", c.getUnsubscribe)
	mux.HandleFunc("POST /subscribe", c.postSubscribe)
	return c
}

func (c *subscription) getConfirm(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")

	res, err := c.service.ConfirmEmail(&command.ConfirmEmail{
		Token: token,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *subscription) getUnsubscribe(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")

	res, err := c.service.Unsubscribe(&command.Unsubscribe{
		Token: token,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *subscription) postSubscribe(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	s := &command.Subscribe{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.service.Subscribe(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
