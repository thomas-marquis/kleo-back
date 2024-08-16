package user

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserRepository      = errors.New("user repository error")
	ErrHouseholdRepository = errors.New("household repository error")
)
