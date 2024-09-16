package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/thomas-marquis/kleo-back/internal/domain"
	"github.com/thomas-marquis/kleo-back/internal/infrastructure"
)

type TransactionService struct {
	repo domain.TransactionRepository
}

func NewTransactionService(repo domain.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) AllocateToUser(ctx context.Context, transaction *domain.Transaction, userID domain.UserId, ratio float64) error {
	if err := transaction.Allocate(userID, ratio); err != nil {
		return err
	}

	if err := s.repo.Save(ctx, transaction); err != nil {
		return fmt.Errorf("an error occurred when saving transaction: %w", err)
	}

	return nil
}

func (s *TransactionService) ListUserTransactions(ctx context.Context, userId domain.UserId, limit int, offset int) ([]*domain.Transaction, error) {
	transactions, err := s.repo.FindByUserId(ctx, userId, limit, offset)
	if err != nil {
		if errors.Is(err, infrastructure.ErrDataNotFound) {
			return nil, domain.ErrTransactionNotFound
		}
		return nil, fmt.Errorf("an error occurred when fetching transactions: %w", err)
	}
	return transactions, nil
}
