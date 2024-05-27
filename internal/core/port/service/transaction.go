package service

import (
	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
)

type TransactionService interface {
	Find(filter value.Filter, page, page_size int32) ([]entity.Transaction, bool, error)
	FindById(id string) (entity.Transaction, error)
}
