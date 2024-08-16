package user

import (
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/bank"
)

type User struct {
	ID   uuid.UUID
	Name string
}

func NewUser(name string) User {
	return User{
		ID:   uuid.New(),
		Name: name,
	}
}

func (u *User) Login() {
	// Login logic
}

func (u *User) Logout() {
	// Logout logic
}

func (u *User) CreateHousehold(name string) Household {
	hh := NewHousehold(name)

	return hh
}

func (u *User) AssociateToBankAccount(ba bank.BankAccount) error {
	return nil
}
