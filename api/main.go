package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/eminetto/go-web-skel/api/handler"
	"github.com/eminetto/go-web-skel/pkg/company"
	"github.com/eminetto/go-web-skel/pkg/mysql"
	"github.com/eminetto/go-web-skel/pkg/user"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {

	env := os.Getenv("SKEL_ENV")
	err := godotenv.Load("config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConf := mysql.DBConfig{
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
	}
	dbConn, err := mysql.InitDb(dbConf)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
	}
	migrateDatabase(dbConn)
	StartServer(dbConn)
}

func migrateDatabase(dbConn *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	n, err := migrate.Exec(dbConn, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Printf("Error applying migrations: %v\n", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

//StartServer rotas e handlers
func StartServer(dbConn *sql.DB) {

	r := mux.NewRouter()
	//company
	cService := company.NewService(dbConn)
	handler.MakeCompanyHandlers(r, cService)

	//user
	uService := user.NewService(dbConn)
	handler.MakeUserHandlers(r, uService)

	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(http.DefaultServeMux))

}
