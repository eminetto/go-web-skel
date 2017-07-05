package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	mysql "github.com/go-sql-driver/mysql"
	handle_error "gitlab.com/thecodenation/thecodenation/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type parameters struct {
	Email    string `valid:"email,required"`
	Password string `valid:"required"`
}

//AuthUser dados do usuário logado
type AuthUser struct {
	ID           int64  `json:"user_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	IsAdmin      int    `json:"is_admin"`
	CompanyID    int64  `json:"company_id"`
}

//AuthStorage onde está armazenado o login
var AuthStorage *sql.DB

//AuthHandler trata a autenticação
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("Invalid request method")
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var p parameters

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	_, err = valid.ValidateStruct(p)
	if err != nil {
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	u, err := Auth(p.Email, p.Password)
	if err != nil {
		handle_error.HandleError(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"nbf":     time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	u.Token, err = token.SignedString([]byte("slkm211221oesdnadsajdsa"))
	if err != nil {
		handle_error.HandleError(err)
		http.Error(w, err.Error(), 500)
		return
	}
	h := sha256.New()
	h.Write([]byte(u.Token))
	u.RefreshToken = hex.EncodeToString(h.Sum(nil))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

//Auth do the authentication
func Auth(email, password string) (AuthUser, error) {
	var u AuthUser
	var dbPassword string
	AuthStorage, err := getAuthStorage()
	if err != nil {
		return u, err
	}
	defer AuthStorage.Close()
	// Prepare statement for reading data
	stmtOut, err := AuthStorage.Prepare("SELECT user.id, user.email, user.password, user.is_admin, COALESCE(user_company.company_id, 0) company_id FROM user LEFT JOIN user_company ON user.id = user_company.user_id where email=?")
	if err != nil {
		return u, err
	}
	defer stmtOut.Close()
	// Query
	err = stmtOut.QueryRow(email).Scan(&u.ID, &u.Email, &dbPassword, &u.IsAdmin, &u.CompanyID)
	if err != nil {
		return u, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		err := errors.New("Invalid password")
		return u, err
	}
	return u, nil
}

//SetAuthStorage configura uma storage diferente do padrão
func SetAuthStorage(st *sql.DB) {
	AuthStorage = st
}

func getAuthStorage() (*sql.DB, error) {
	if AuthStorage != nil {
		return AuthStorage, nil
	}
	var dbConfig mysql.Config
	dbConfig.User = os.Getenv("DATABASE_USER")
	dbConfig.Passwd = os.Getenv("DATABASE_PASSWORD")
	dbConfig.DBName = os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Passwd, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), dbConfig.DBName)
	AuthStorage, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return AuthStorage, nil
}
