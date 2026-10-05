package user

import "github.com/google/uuid"

type User struct {
	ID   string
	Name string
}

func NewUser(name string) User {
	id := uuid.NewString()
	return User{
		ID:   id,
		Name: name,
	}
}
