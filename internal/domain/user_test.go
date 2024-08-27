package domain_test

import (
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

func newUserId() domain.UserId {
	return domain.UserId(uuid.New())
}
