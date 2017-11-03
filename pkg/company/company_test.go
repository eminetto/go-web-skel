package company

import (
	valid "github.com/asaskevich/govalidator"
	"testing"
)
func TestCompanyValidation(t *testing.T) {
	c := Company{
		Email: "eminetto@email.com",
		Name:  "Big Co",
		URL:   "http://bigco.com",
	}
	_, err := valid.ValidateStruct(c)
	if err != nil {
		t.Errorf("expected %s result %s", nil, err)
	}
}
