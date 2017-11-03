package user

import "time"

//User dados dos usuários que podem logar no sistema
type User struct {
	ID        int64     `valid:"-" json:"id" db:"id"`
	Name      string    `valid:"required" json:"name" db:"name"`
	Picture   string    `valid:"url,optional" json:"picture" db:"picture"`
	Email     string    `valid:"email,required" json:"email" db:"email"`
	Password  string    `valid:"required" json:"password" db:"password"`
	CreatedAt time.Time `valid:"-" db:"created_at" json:"created_at"`
}

//ToJSON extra information used when translating to json
type ToJSON struct {
	*User
	Type string `json:"type"`
}

//Service service interface
type Service interface {
	Find(id int64) (*User, error)
	Search(query string) ([]*User, error)
	FindAll() ([]*User, error)
	Remove(ID int64) error
	Store(user *User) (int64, error)
	ToJSON(u *User) (ToJSON, error)
}
