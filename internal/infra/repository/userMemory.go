package repository

import (
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/user"
)

type UserMemoryRepository struct {
	users        map[uuid.UUID]user.User
	byHousholdId map[uuid.UUID][]uuid.UUID
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{}
}

func (r *UserMemoryRepository) GetById(id uuid.UUID) (user.User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	} else {
		return user.User{}, user.ErrUserNotFound
	}
}

func (r *UserMemoryRepository) ListByHouseholdId(id uuid.UUID) ([]user.User, error) {
	if ids, ok := r.byHousholdId[id]; ok {
		users := make([]user.User, len(ids))
		for i, id := range ids {
			if u, ok := r.users[id]; !ok {
				return nil, user.ErrUserNotFound
			} else {
				users[i] = u
			}
		}
		return users, nil
	} else {
		return nil, user.ErrUserNotFound
	}
}

func (r *UserMemoryRepository) Create(name string) (user.User, error) {
	u := user.NewUser(name)
	r.users[u.ID] = u

	return u, nil
}

func (r *UserMemoryRepository) Update(u user.User) (user.User, error) {
	if _, ok := r.users[u.ID]; !ok {
		return user.User{}, user.ErrUserNotFound
	}

	r.users[u.ID] = u

	return u, nil
}

var _ user.UserRepository = &UserMemoryRepository{}
