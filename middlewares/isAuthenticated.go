package middlewares

import (
	"gitlab.com/thecodenation/thecodenation/errors"
	"gitlab.com/thecodenation/thecodenation/session"
	"net/http"
)

//IsAuthenticated verifica se o usuário tem uma sessão
func IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, err := session.Store.Get(r, "auth-session")
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}
