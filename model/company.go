package model

//Company dados da empresa
type Company struct {
	ID    int64  `valid:"-"`
	Name  string `valid:"required"`
	Email string `valid:"email,required"`
	URL   string `valid:"url,required"`
}
