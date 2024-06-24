package v1

import (
	"encoding/json"
	"net/http"
)

func (r *Router) AttachRoutes() {
	r.router.Get("/ping", r.Ping)
	r.router.Get("/login", r.Login)
	r.router.Get("/signup", r.SignUp)

	protected := r.router.With(checkAuthentication)
	protected.Get("/proping", r.PingProtected)
}

func (r *Router) Ping(rw http.ResponseWriter, rq *http.Request) {
	_ = json.NewEncoder(rw).Encode(RespMessage{Message: "pong", Status: http.StatusOK})
}

func (r *Router) PingProtected(rw http.ResponseWriter, rq *http.Request) {
	_ = json.NewEncoder(rw).Encode(RespMessage{Message: "protected pong", Status: http.StatusOK})
}
