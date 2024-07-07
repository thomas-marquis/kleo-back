package repository

import (
	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
)

type TransactionRepository interface {
	FindTransactionsByFilter(filter value.Filter, size, offset int32) ([]entity.Transaction, error)
	FindTransactionById(id string) (entity.Transaction, error)
}
