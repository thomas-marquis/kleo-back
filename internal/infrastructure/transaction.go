package infrastructure

import (
	"context"
	"errors"

	"github.com/thomas-marquis/kleo-back/internal/domain"
	"github.com/thomas-marquis/kleo-back/internal/infrastructure/orm"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

var _ domain.TransactionRepository = &TransactionRepositoryImpl{}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (r *TransactionRepositoryImpl) Save(ctx context.Context, transaction *domain.Transaction) error {
	tr := orm.ToTransactionModel(transaction)
	if err := r.db.WithContext(ctx).Create(tr).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) SaveRaw(ctx context.Context, rawTransaction *domain.RawTransaction) error {
	rt := orm.ToRawTransactionModel(rawTransaction)
	if err := r.db.WithContext(ctx).Create(rt).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) FindByUserId(ctx context.Context, userId domain.UserId, limit, offset int) ([]*domain.Transaction, error) {
	var transactions []*orm.Transaction
	if err := r.db.WithContext(ctx).Model(&orm.Transaction{}).
		Where("user_id = ?", userId.String()).
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, ErrDataNotFound
	}

	domainTransactions := make([]*domain.Transaction, len(transactions))
	for i, t := range transactions {
		domainTransactions[i] = orm.ToTransaction(t)
	}

	return domainTransactions, nil
}
