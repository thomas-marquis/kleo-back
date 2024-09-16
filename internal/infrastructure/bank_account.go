package infrastructure

import (
	"context"

	"github.com/thomas-marquis/kleo-back/internal/domain"
	"github.com/thomas-marquis/kleo-back/internal/infrastructure/orm"
	"gorm.io/gorm"
)

type BankAccountRepositoryImpl struct {
	db *gorm.DB
}

var _ domain.BankAccountRepository = &BankAccountRepositoryImpl{}

func NewBankAccountRepositoryImpl(db *gorm.DB) *BankAccountRepositoryImpl {
	return &BankAccountRepositoryImpl{
		db: db,
	}
}

func (r *BankAccountRepositoryImpl) Save(ctx context.Context, bankAccount *domain.BankAccount) error {
	accModel := orm.ToAccountModel(bankAccount)
	if err := r.db.Create(accModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *BankAccountRepositoryImpl) FindById(ctx context.Context, id domain.BankAccountId) (*domain.BankAccount, error) {
	return nil, nil
}
