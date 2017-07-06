package model

type userdb interface {
    GetUsers() ([]*User, error)
    GetUser(id int64) (*User, error)
    CreateUser(u User) (int64, error)
    UpdateUser(u User) error
    DeleteUser(id int64) error
}

type UserModel struct {
    userdb
}

func NewUserModel(db userdb) *UserModel {
    return &UserModel{
        userdb: db,
    }
}

// func (m *UserModel) GetAllUsers() ([]*User, error) {
//     return m.GetAll()
// }

//User dados dos usuários que podem logar no sistema
type User struct {
    ID       int64  `valid:"-"`
    Name     string `valid:"required"`
    Picture  string `valid:"url,optional"`
    Email    string `valid:"email,required"`
    Password string `valid: required`
}
