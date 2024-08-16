package user

import "github.com/google/uuid"

type Household struct {
	ID   uuid.UUID
	Name string
}

func NewHousehold(name string) Household {
	return Household{
		ID:   uuid.New(),
		Name: name,
	}
}

func (h *Household) AddMember(u User) {
	// Add member logic
}

func (h *Household) ChangeRole(u User, newRole Role) {
	// Change role logic
}

func (h *Household) ListUsers() []User {
	// List users logic
	return nil
}
