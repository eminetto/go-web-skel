package model

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/eminetto/go-web-skel/model"
	"testing"
)

func TestUserValidation(t *testing.T) {
	u := model.User{
		Picture:  "https://avatars0.githubusercontent.com/u/19712121939?:3",
		Email:    "eminetto@email.com",
		Name:     "Elton Minetto",
		Password: "sfsdfdsdsf",
	}
	_, err := valid.ValidateStruct(u)
	if err != nil {
		t.Errorf("expected %s result %s", nil, err)
	}
}

func TestCompanyValidation(t *testing.T) {
	c := model.Company{
		Email: "eminetto@email.com",
		Name:  "Big Co",
		URL:   "http://bigco.com",
	}
	_, err := valid.ValidateStruct(c)
	if err != nil {
		t.Errorf("expected %s result %s", nil, err)
	}
}
