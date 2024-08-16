package repository

import (
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/user"
)

type UserMySQLRepository struct{}

func (r *UserMySQLRepository) GetById(id uuid.UUID) (user.User, error) {
	// Get user logic
	return user.User{}, nil
}

func (r *UserMySQLRepository) StoreOrUpdate(u user.User) error {
	// Store or update user logic
	return nil
}

var _ user.UserRepository = &UserMySQLRepository{}
