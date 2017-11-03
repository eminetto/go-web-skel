package company

import "time"

//Company dados da empresa
type Company struct {
	ID        int64     `valid:"-" json:"id" db:"id"`
	Name      string    `valid:"required" json:"name" db:"name"`
	Email     string    `valid:"email,required" json:"email" db:"email"`
	URL       string    `valid:"url,required" json:"url" db:"url"`
	CreatedAt time.Time `valid:"-" db:"created_at" json:"created_at"`
}

//Service service interface
type Service interface {
	Find(id int64) (*Company, error)
	Search(query string) ([]*Company, error)
	FindAll() ([]*Company, error)
	Remove(ID int64) error
	Store(company *Company) (int64, error)
}
