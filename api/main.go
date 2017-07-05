package main

import (
	"log"
	"net/http"

	// "github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	api "gitlab.com/thecodenation/thecodenation/api/handlers"
	"gitlab.com/thecodenation/thecodenation/datastore/database"
	"gitlab.com/thecodenation/thecodenation/errors"
	// "gitlab.com/thecodenation/thecodenation/middlewares"
	"fmt"
	"os"
)

func main() {

	env := os.Getenv("CODENATION_ENV")
	err := godotenv.Load("config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	errors.Init()
	StartServer()
}

//DBStorage banco de dados
// var DBStorage *sql.DB

//StartServer rotas e handlers
func StartServer() {
	db, err := database.Connect()
	if err != nil {
		errors.HandleError(err)
	}
	defer db.Close()
	var userDS = database.New(db)
	var applicantDS = database.New(db)
	r := mux.NewRouter()

	r.HandleFunc("/v1/auth", api.AuthHandler)
	r.HandleFunc("/v1/applicant", api.ApplicantIndex(applicantDS))
	r.HandleFunc("/v1/user", api.ApplicantIndex(userDS))
	// r.Handle("/v1/applicant", negroni.New(
	// 	negroni.HandlerFunc(middlewares.IsJWTAuthenticated),
	// 	negroni.Wrap(http.HandlerFunc(api.ApplicantHandler)),
	// ))

	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(http.DefaultServeMux))
}
