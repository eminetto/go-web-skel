package mysql

import (
	"database/sql"

	"fmt"
	"os"
	//blank import required by sql
	_ "github.com/go-sql-driver/mysql"
)

type mDB struct {
	db *sql.DB
}

//DBConfig database config
type DBConfig struct {
	User   string
	Passwd string
	DBName string
	DBHost string
	DBPort string
}

//InitDb faz a conexão com o banco
func InitDb(dbConfig DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Passwd, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), dbConfig.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
