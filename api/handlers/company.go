package handlers

import (
	"encoding/json"
	"github.com/eminetto/go-web-skel/model"
	"net/http"
)

type CompanyActions interface {
	GetCompanies() ([]*model.Company, error)
	GetCompany(id int64) (*model.Company, error)
	CreateCompany(u model.Company) (int64, error)
	UpdateCompany(u model.Company) error
	DeleteCompany(id int64) error
}

func CompanyIndex(repo CompanyActions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := repo.GetCompanies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
