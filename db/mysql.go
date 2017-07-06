package db

import (
    "database/sql"

    "fmt"
    mysql "github.com/go-sql-driver/mysql"
    "os"
)

type mDB struct {
    db *sql.DB
}

// type Config struct {
//     ConnectString string
// }

//InitDb faz a conexão com o banco
func InitDb() (*mDB, error) {
    // func InitDb(cfg Config) (*sql.DB, error) {
    //@TODO receber as configurações via Config
    //@TODO usar sqlx
    var dbConfig mysql.Config
    dbConfig.User = os.Getenv("DATABASE_USER")
    dbConfig.Passwd = os.Getenv("DATABASE_PASSWORD")
    dbConfig.DBName = os.Getenv("DATABASE_NAME")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Passwd, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), dbConfig.DBName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    p := &mDB{db: db}
    return p, nil
}
