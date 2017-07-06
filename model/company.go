package model

type companydb interface {
    GetCompanies() ([]*Company, error)
    GetCompnay(id int64) (*Company, error)
    CreateCompany(u Company) (int64, error)
    UpdateCompany(u Company) error
    DeleteCompany(id int64) error
}

type CompanyModel struct {
    companydb
}

func NewCompanyModel(db companydb) *CompanyModel {
    return &CompanyModel{
        companydb: db,
    }
}

// func (m *CompanyModel) GetAllCompanies() ([]*Company, error) {
//     return m.GetAll()
// }

//Company dados da empresa
type Company struct {
    ID    int64  `valid:"-"`
    Name  string `valid:"required"`
    Email string `valid:"email,required"`
    URL   string `valid:"url,required"`
}
