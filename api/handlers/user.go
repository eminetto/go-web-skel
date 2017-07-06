package handlers

import (
	"encoding/json"
	valid "github.com/asaskevich/govalidator"
	"github.com/eminetto/go-web-skel/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserActions interface {
	GetUsers() ([]*model.User, error)
	GetUser(id int64) (*model.User, error)
	CreateUser(u model.User) (int64, error)
	UpdateUser(u model.User) error
	DeleteUser(id int64) error
}

func UserIndex(repo UserActions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := repo.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func UserGet(repo UserActions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data, err := repo.GetUser(id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func UserDelete(repo UserActions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = repo.DeleteUser(id)
		if err := json.NewEncoder(w).Encode(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func UserCreate(repo UserActions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		_, err = valid.ValidateStruct(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		id, err := repo.CreateUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
