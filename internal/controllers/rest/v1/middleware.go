package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
)

func checkAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		tokenString := rq.Header.Get("Authorization")
		if tokenString == "" {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(RespMessage{
				Status: http.StatusBadRequest,
				Error:  apperror.NewAuthenticationError(errors.New("bearer token missing")).Error(),
			})
			return
		}
		tokenString = tokenString[len("Bearer "):]

		id, err := verifyToken(tokenString)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(RespMessage{
				Status: http.StatusBadRequest,
				Error:  apperror.NewAuthenticationError(err).Error(),
			})
			return
		}
		ctx := context.WithValue(rq.Context(), UserIDKey, id)
		next.ServeHTTP(rw, rq.WithContext(ctx))
	})
}

type ctxKey string

const UserIDKey ctxKey = "user-id"
