package application

import (
	"context"
	"fmt"

	"github.com/thomas-marquis/kleo-back/internal/domain"
)

type InitLegacyService struct {
	repo            domain.LegacyRepository
	transactionRepo domain.TransactionRepository
}

func NewInitLegacyService(repo domain.LegacyRepository, transactionRepo domain.TransactionRepository) *InitLegacyService {
	return &InitLegacyService{repo: repo, transactionRepo: transactionRepo}
}

func (s *InitLegacyService) Sync(ctx context.Context) error {
	accounts, err := s.repo.GetBankAccounts(ctx)
	if err != nil {
		return fmt.Errorf("an error occurred when fetching bank accounts: %w", err)
	}

	for _, acc := range accounts {
		dateParseRegexp, err := s.repo.GetDateParseRegexpByAccountId(ctx, acc.ID)
		if err != nil {
			return fmt.Errorf("an error occurred when fetching date parse regexp: %w", err)
		}

		rawTranasactions, err := s.repo.GetRawTransactionsByAccountId(ctx, acc.ID)
		if err != nil {
			return fmt.Errorf("an error occurred when fetching raw transactions: %w", err)
		}

		for _, rt := range rawTranasactions {
			err := s.transactionRepo.SaveRaw(ctx, rt)
			if err != nil {
				return fmt.Errorf("an error occurred when saving raw transaction: %w", err)
			}

			cleanDate, err := rt.ExtractDateFromLabel(&dateParseRegexp)
			if err != nil {
				return fmt.Errorf("an error occurred when extracting date from label: %w", err)
			}
			cleanLabel, err := rt.CleanupLabel(&dateParseRegexp)
			if err != nil {
				return fmt.Errorf("an error occurred when cleaning label: %w", err)
			}

			tr := domain.NewTransaction(rt.Amount, cleanLabel, cleanDate)

			categ, err := s.repo.GetCategoryFromMetadata(ctx, rt.Metadata)
			if err != nil {
				return fmt.Errorf("an error occurred when fetching category from metadata: %w", err)
			}
			if categ != nil {
				tr.UpdateCategory(categ)
			}

			err = s.transactionRepo.Save(ctx, tr)
			if err != nil {
				return fmt.Errorf("an error occurred when saving transaction: %w", err)
			}
		}
	}

	return nil
}
