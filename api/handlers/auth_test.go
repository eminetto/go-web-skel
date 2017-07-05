package handlers

import (
	"bytes"
	"encoding/json"
	api "gitlab.com/thecodenation/thecodenation/api/handlers"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

var query = "SELECT (.+), (.+), (.+), (.+) FROM"

type parameters struct {
	Email    string `valid:"email,required"`
	Password string `valid:"required"`
}

func TestValidLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	api.SetAuthStorage(db)
	defer db.Close()

	hash, _ := bcrypt.GenerateFromPassword([]byte("valid_password"), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "is_admin", "company_id"}).
		AddRow(1, "eminetto@gmail.com", string(hash), 1, 0)

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs("eminetto@gmail.com").
		WillReturnRows(rows)

	u, err := api.Auth("eminetto@gmail.com", "valid_password")

	if err != nil {
		t.Errorf("Should not be nil: %s", err)
	}
	if u.ID != 1 {
		t.Errorf("Should be 1: %s", u.ID)
	}
}

func TestInvalidLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	api.SetAuthStorage(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "is_admin"})

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs("eminetto@gmail.com").
		WillReturnRows(rows)

	u, err := api.Auth("eminetto@gmail.com", "valid_password")

	if err == nil {
		t.Errorf("Should be nil: %s", err)
	}
	if u.ID == 1 {
		t.Errorf("Should be 0: %s", u.ID)
	}
}

func TestInvalidPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	api.SetAuthStorage(db)
	defer db.Close()

	hash, _ := bcrypt.GenerateFromPassword([]byte("valid_password"), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "is_admin"}).
		AddRow(1, "eminetto@gmail.com", string(hash), 1)

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs("eminetto@gmail.com").
		WillReturnRows(rows)

	_, err = api.Auth("eminetto@gmail.com", "invalid_password")

	if err == nil {
		t.Errorf("Should be nil: %s", err)
	}
}

func TestInvalidParameters(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	api.SetAuthStorage(db)
	defer db.Close()
	u, err := api.Auth("eminettogmail.com", "1")

	if err == nil {
		t.Errorf("Should be not nil: %s", err)
	}
	if u.ID != 0 {
		t.Errorf("Should be 0: %s", u.ID)
	}
}

func TestIntegrationTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("valid_password"), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "is_admin", "company_id"}).
		AddRow(1, "eminetto@gmail.com", string(hash), 1, 0)

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs("eminetto@gmail.com").
		WillReturnRows(rows)

	api.SetAuthStorage(db)
	defer db.Close()

	p := parameters{Email: "eminetto@gmail.com", Password: "valid_password"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	req, _ := http.NewRequest("POST", "/v1/auth", b)
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	api.AuthHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var u api.AuthUser
	if err := json.NewDecoder(rr.Body).Decode(&u); err != nil {
		panic(err)
	}
	if u.Email != "eminetto@gmail.com" {
		t.Errorf("Invalid email: %s", u.Email)
	}
	if u.Token == "" {
		t.Errorf("Invalid token: %s", u.Token)
	}
	if u.RefreshToken == "" {
		t.Errorf("Invalid refresh token: %s", u.RefreshToken)
	}
}
