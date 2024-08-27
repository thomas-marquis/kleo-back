package domain

import (
	"context"
	"regexp"
)

type TransactionRepository interface {
	Save(ctx context.Context, transaction *Transaction) error
	SaveRaw(ctx context.Context, rawTransaction *RawTransaction) error
	FindByUserId(ctx context.Context, userId UserId) ([]*Transaction, error)
}

type BankAccountRepository interface {
	Save(ctx context.Context, bankAccount *BankAccount) error
	FindById(ctx context.Context, id BankAccountId) (*BankAccount, error)
}

type LegacyRepository interface {
	GetBankAccounts(ctx context.Context) ([]*BankAccount, error)
	GetRawTransactionsByAccountId(ctx context.Context, accountID BankAccountId) ([]*RawTransaction, error)
	GetCategoryByOldLabel(ctx context.Context, oldLabel string) (*Category, error)
	GetDateParseRegexpByAccountId(ctx context.Context, accountID BankAccountId) (regexp.Regexp, error)
	GetCategoryFromMetadata(ctx context.Context, metadata map[string]interface{}) (*Category, error)
}
