package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	handle_error "gitlab.com/thecodenation/thecodenation/errors"
	"net/http"
	"strconv"
)

//IsJWTAuthenticated verifica se o usuário tem uma sessão
func IsJWTAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := errors.New("Unauthorized")
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("slkm211221oesdnadsajdsa"), nil
	})
	if err != nil {
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		r.Header.Add("user_id", strconv.Itoa(int(claims["user_id"].(float64))))
	} else {
		err := errors.New("Unauthorized")
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	next(w, r)
}
