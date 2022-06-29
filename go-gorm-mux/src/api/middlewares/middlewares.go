package middlewares

import (
	"errors"
	"go-gorm-mux/src/api/auth"
	"go-gorm-mux/src/api/responses"
	"net/http"
)

// SetMiddlewareJSON sets the Content-Type to application/json in the response header.
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication sets the authentication token in the request header.
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
