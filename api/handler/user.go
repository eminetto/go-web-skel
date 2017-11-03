package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/go-web-skel/pkg/middleware"
	"github.com/eminetto/go-web-skel/pkg/user"
	"github.com/gorilla/mux"
)

func userFindAll(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*user.User
		var err error
		var dataToJSON []user.ToJSON
		data, err = service.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(data) == 0 {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		for _, j := range data {
			d, err := service.ToJSON(j)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dataToJSON = append(dataToJSON, d)
		}
		if err := json.NewEncoder(w).Encode(dataToJSON); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func userFind(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u, err := service.Find(id)
		if err != nil || u == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		d, err := service.ToJSON(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(d); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func userRemove(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = service.Remove(id)
		if err != nil {
			http.Error(w, "Error removing user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func userAdd(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			user.User
		}
		var param parameters
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userID, err := service.Store(&param.User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type data struct {
			ID int64 `json:"id"`
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(data{ID: userID}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, service user.Service) {
	r.Handle("/v1/user", negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.IsAuthenticated),
		negroni.Wrap(userFindAll(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/user", negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.Wrap(userAdd(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/user/{id}", negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.IsAuthenticated),
		negroni.Wrap(userFind(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/user/{id}", negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.IsAuthenticated),
		negroni.Wrap(userRemove(service)),
	)).Methods("DELETE", "OPTIONS")
}
