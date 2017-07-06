package db

import (
    "github.com/eminetto/go-web-skel/model"
)

func (p *mDB) GetCompanies() ([]*model.Company, error) {
    var people []*model.Company

    for i := 0; i < 5; i++ {
        people = append(people, &model.Company{1, "Name", "email", "url"})
    }
    return people, nil
}

func (p *mDB) GetCompany(id int64) (*model.Company, error) {
    var u *model.Company
    return u, nil
}

func (p *mDB) CreateCompany(u model.Company) (int64, error) {
    return 1, nil
}

func (p *mDB) UpdateCompany(u model.Company) error {
    return nil
}

func (p *mDB) DeleteCompany(id int64) error {
    return nil
}
