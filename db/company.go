package db

import (
	"github.com/eminetto/go-web-skel/model"
	"github.com/jmoiron/sqlx"
)

type CompanyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

func (r *CompanyRepo) GetCompanies() ([]*model.Company, error) {
	var people []*model.Company

	for i := 0; i < 5; i++ {
		people = append(people, &model.Company{1, "Name", "email", "url"})
	}
	return people, nil
}

func (r *CompanyRepo) GetCompany(id int64) (*model.Company, error) {
	var u *model.Company
	return u, nil
}

func (r *CompanyRepo) CreateCompany(u model.Company) (int64, error) {
	return 1, nil
}

func (r *CompanyRepo) UpdateCompany(u model.Company) error {
	return nil
}

func (r *CompanyRepo) DeleteCompany(id int64) error {
	return nil
}
