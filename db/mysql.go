package db

import (
	"database/sql"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type mDB struct {
	db *sql.DB
}

type DBConfig struct {
	User   string
	Passwd string
	DBName string
	DBHost string
	DBPort string
}

//InitDb faz a conexão com o banco
func InitDb(dbConfig DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Passwd, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), dbConfig.DBName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
