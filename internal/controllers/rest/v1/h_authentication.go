package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (r *Router) Login(rw http.ResponseWriter, rq *http.Request) {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := JsonDecode(rq.Body, &body); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(rq.Context(), time.Minute*1)
	defer cancel()
	_, err := r.userService.Login(ctx, body.Username, body.Password)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusUnauthorized, Error: err.Error()})
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusOK, Message: "authenticated"})
}

func (r *Router) SignUp(rw http.ResponseWriter, rq *http.Request) {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := JsonDecode(rq.Body, &body); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()
	_, err := r.userService.SignUp(ctx, body.Username, body.Password)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusUnauthorized, Error: err.Error()})
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(RespMessage{Status: http.StatusOK, Message: "user created"})
}
