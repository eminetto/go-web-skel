package middleware

import (
	"errors"
	"net/http"
)

//IsAuthenticated verifica se o usuário tem uma sessão
func IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := errors.New("Unauthorized")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	next(w, r)
}
