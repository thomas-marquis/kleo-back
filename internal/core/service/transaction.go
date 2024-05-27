package service

import (
	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/port/service"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
)

type TransactionServiceImpl struct {
}

func NewTransactionServiceImpl() *TransactionServiceImpl {
	return &TransactionServiceImpl{}
}

func (s *TransactionServiceImpl) Find(filter value.Filter, page, page_size int32) ([]entity.Transaction, bool, error) {
	return nil, false, nil
}

func (s *TransactionServiceImpl) FindById(id string) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

var _ service.TransactionService = &TransactionServiceImpl{}
