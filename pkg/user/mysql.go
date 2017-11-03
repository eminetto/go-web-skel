package user

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
	dbmap.AddTableWithName(User{}, "user").SetKeys(true, "id")
	return &service{
		dbmap: dbmap,
	}
}

func (s *service) Find(id int64) (*User, error) {
	d, err := s.dbmap.Get(User{}, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, nil
	}
	return d.(*User), nil
}

func (s *service) FindAll() ([]*User, error) {
	var d []*User
	_, err := s.dbmap.Select(&d, "select * from user order by id")
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Search(query string) ([]*User, error) {
	var d []*User
	_, err := s.dbmap.Select(&d, "select * from user where name like ? order by name", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Store(p *User) (int64, error) {
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
	user, err := s.Find(ID)
	if err != nil {
		return errors.New("Error reading user")
	}
	_, err = s.dbmap.Delete(user)
	return err
}

func (s *service) ToJSON(u *User) (ToJSON, error) {
	var d ToJSON
	d.User = u
	d.Type = "Regular User"
	return d, nil
}
