package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/thecodenation/thecodenation/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWithoutToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/auth", nil)
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	middlewares.IsJWTAuthenticated(rr, req, nil)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

func TestWithInvalidToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/auth", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "invalid_token")
	rr := httptest.NewRecorder()

	middlewares.IsJWTAuthenticated(rr, req, nil)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

func TestWithValidToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/auth", nil)
	req.Header.Add("Content-Type", "application/json")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
		"nbf":     time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("slkm211221oesdnadsajdsa"))
	req.Header.Add("Authorization", tokenString)
	rr := httptest.NewRecorder()

	middlewares.IsJWTAuthenticated(rr, req, http.HandlerFunc(handler))

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//@TODO ver como incluir um teste aqui, para verificar se o r.Header.Get("user_id") == 1
}
