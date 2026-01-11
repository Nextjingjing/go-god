package user

type User struct {
	id    int
	Name  string
	Email string
}

func (u *User) GetID() int {
	return u.id
}

func (u *User) SetID(id int) {
	u.id = id
}

func NewUser(id int, name string, email string) User {
	return User{
		id:    id,
		Name:  name,
		Email: email,
	}
}
