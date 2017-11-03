package company

import (
	"database/sql"
	"errors"
	"time"

	gorp "gopkg.in/gorp.v1"
)

type service struct {
	dbmap *gorp.DbMap
}

//NewService create new service
func NewService(db *sql.DB) Service {
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Company{}, "company").SetKeys(true, "id")
	return &service{
		dbmap: dbmap,
	}
}

func (s *service) Find(id int64) (*Company, error) {
	d, err := s.dbmap.Get(Company{}, id)
	if err != nil {
		return nil, err
	}
	return d.(*Company), nil
}

func (s *service) FindAll() ([]*Company, error) {
	var d []*Company
	_, err := s.dbmap.Select(&d, "select * from company order by id")
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Search(query string) ([]*Company, error) {
	var d []*Company
	_, err := s.dbmap.Select(&d, "select * from company where name like ? order by name", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Store(p *Company) (int64, error) {
	if p.ID == 0 {
		p.CreatedAt = time.Now()
		err := s.dbmap.Insert(p)
		if err != nil {
			return 0, err
		}
		return p.ID, nil
	}
	_, err := s.dbmap.Update(p)
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (s *service) Remove(ID int64) error {
	var err error
	company, err := s.Find(ID)
	if err != nil {
		return errors.New("Error reading company")
	}
	_, err = s.dbmap.Delete(company)
	return err
}
