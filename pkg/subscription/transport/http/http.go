package http

import (
	"encoding/json"
	"net/http"

	"git.fruzit.pp.ua/weather/api/pkg/subscription/command"
	"git.fruzit.pp.ua/weather/api/pkg/subscription/service"
)

type Transport struct {
	service service.IService
}

func New(mux *http.ServeMux, service service.IService) *Transport {
	c := &Transport{service}
	mux.HandleFunc("GET /confirm/{token}", c.getConfirm)
	mux.HandleFunc("GET /unsubscribe/{token}", c.getUnsubscribe)
	mux.HandleFunc("POST /subscribe", c.postSubscribe)
	return c
}

func (t *Transport) getConfirm(w http.ResponseWriter, r *http.Request) {
	// TODO: Code 404 Invalid token
	res, err := t.service.ConfirmEmail(&command.ConfirmEmail{
		Token: r.PathValue("token"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (t *Transport) getUnsubscribe(w http.ResponseWriter, r *http.Request) {
	// TODO: Code 404 Invalid token
	res, err := t.service.Unsubscribe(&command.Unsubscribe{
		Token: r.PathValue("token"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (t *Transport) postSubscribe(w http.ResponseWriter, r *http.Request) {
	s := &command.Subscribe{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := t.service.Subscribe(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
