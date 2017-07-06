package main

import (
	"log"
	"net/http"

	"github.com/eminetto/go-web-skel/api/handlers"
	"github.com/eminetto/go-web-skel/db"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	env := os.Getenv("SKEL_ENV")
	err := godotenv.Load("config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	StartServer()
}

//StartServer rotas e handlers
func StartServer() {
	dbConf := db.DBConfig{
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
	}
	dbConn, err := db.InitDb(dbConf)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
	}
	u := db.NewUserRepo(dbConn)
	c := db.NewCompanyRepo(dbConn)
	r := mux.NewRouter()

	r.Handle("/user", handlers.UserIndex(u)).Methods("GET")
	r.Handle("/user", handlers.UserCreate(u)).Methods("POST")
	r.Handle("/user/{id}", handlers.UserDelete(u)).Methods("DELETE")
	r.Handle("/user/{id}", handlers.UserGet(u)).Methods("GET")
	r.Handle("/company", handlers.CompanyIndex(c))

	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(http.DefaultServeMux))
}
