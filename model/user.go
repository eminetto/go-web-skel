package model

//User dados dos usuários que podem logar no sistema
type User struct {
	ID       int64  `valid:"-"`
	Name     string `valid:"required"`
	Picture  string `valid:"url,optional"`
	Email    string `valid:"email,required"`
	Password string `valid: required`
}
