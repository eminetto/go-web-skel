package db

import (
	"github.com/eminetto/go-web-skel/model"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUsers() ([]*model.User, error) {
	var people []*model.User
	rows, err := r.db.Queryx("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u model.User
		err = rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
		people = append(people, &u)
	}
	return people, nil
}

func (r *UserRepo) GetUser(id int64) (*model.User, error) {
	var u model.User
	err := r.db.QueryRowx("SELECT * FROM user where id = ?", id).StructScan(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) CreateUser(u model.User) (int64, error) {
	sql := "insert into user values (null, ?,?,?,?)"
	result, err := r.db.Exec(sql, u.Name, u.Picture, u.Email, u.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}

func (r *UserRepo) UpdateUser(u model.User) error {
	return nil
}

func (r *UserRepo) DeleteUser(id int64) error {
	sql := "delete from user where id = ?"
	_, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
