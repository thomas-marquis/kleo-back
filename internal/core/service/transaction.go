package service

import (
	"fmt"

	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/port/repository"
	"github.com/thomas-marquis/kleo-back/internal/core/port/service"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
)

const (
	defaultPageSize = 50
)

type TransactionServiceImpl struct {
	repository repository.TransactionRepository
}

func NewTransactionServiceImpl(repository repository.TransactionRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		repository: repository,
	}
}

func (s *TransactionServiceImpl) Find(filter value.Filter, page, page_size int32) ([]entity.Transaction, bool, error) {
	if page_size == 0 {
		page_size = defaultPageSize
	}
	var offset int32
	if page != 0 {
		offset = (page - 1) * page_size
	} else {
		offset = 0
	}

	tr, err := s.repository.FindTransactionsByFilter(filter, page_size+1, offset)
	if err != nil {
		return nil, false, fmt.Errorf("failed to find asked transactions: %s", err.Error())
	}

	hasNext := len(tr) > int(page_size)

	return tr, hasNext, nil
}

func (s *TransactionServiceImpl) FindById(id string) (entity.Transaction, error) {
	tr, err := s.repository.FindTransactionById(id)
	if err != nil {
		if err == value.ErrTransactionNotFound {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, fmt.Errorf("failed to find transaction with id %s: %s", id, err.Error())
	}

	return tr, nil
}

var _ service.TransactionService = &TransactionServiceImpl{}
