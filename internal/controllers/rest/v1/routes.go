package v1

import (
	"encoding/json"
	"net/http"
)

func (r *Router) AttachRoutes() {
	r.router.Get("/ping", r.Ping)
	r.router.Get("/login", r.Login)
	r.router.Get("/signup", r.SignUp)
}

func (r *Router) Ping(rw http.ResponseWriter, rq *http.Request) {
	_ = json.NewEncoder(rw).Encode(RespMessage{Message: "pong", Status: http.StatusOK})
}
