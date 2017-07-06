package db

import (
    "github.com/eminetto/go-web-skel/model"
)

func (p *mDB) GetUsers() ([]*model.User, error) {
    var people []*model.User

    for i := 0; i < 5; i++ {
        people = append(people, &model.User{1, "Name", "picture", "email", "password"})
    }
    return people, nil
}

func (p *mDB) GetUser(id int64) (*model.User, error) {
    var u *model.User
    return u, nil
}

func (p *mDB) CreateUser(u model.User) (int64, error) {
    return 1, nil
}

func (p *mDB) UpdateUser(u model.User) error {
    return nil
}

func (p *mDB) DeleteUser(id int64) error {
    return nil
}
