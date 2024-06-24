package entities

type User struct {
	id int
}

func CreateUser(id int) User {
	return User{id: id}
}

func (u User) GetID() int {
	return u.id
}
