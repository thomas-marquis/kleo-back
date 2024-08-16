package user

import "github.com/google/uuid"

type UserRepository interface {
	GetById(id uuid.UUID) (User, error)
	ListByHouseholdId(id uuid.UUID) ([]User, error)
	Create(name string) (User, error)
	Update(u User) (User, error)
}

type HouseholdRepository interface {
	GetById(id uuid.UUID) (Household, error)
}
