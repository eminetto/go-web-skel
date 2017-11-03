package user

import (
	valid "github.com/asaskevich/govalidator"
	"testing"
)

func TestUserValidation(t *testing.T) {
	u := User{
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

