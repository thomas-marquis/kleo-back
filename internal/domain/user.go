package domain

import "github.com/google/uuid"

type UserId uuid.UUID

func (i UserId) String() string {
	return uuid.UUID(i).String()
}

type User struct {
	ID    UserId
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{
		ID:    UserId(uuid.New()),
		Name:  name,
		Email: email,
	}
}
