package domain

import (
	"context"
	"regexp"
)

type TransactionRepository interface {
	// Save saves a transaction in the repository or updates it if it already exists.
	Save(ctx context.Context, transaction *Transaction) error
	SaveRaw(ctx context.Context, rawTransaction *RawTransaction) error

	// FindByUserId returns the paginated list of transactions that belong to given user.
	FindByUserId(ctx context.Context, userId UserId, limit, offset int) ([]*Transaction, error)
}

type BankAccountRepository interface {
	// Save saves a bank account in the repository or updates it if it already exists.
	Save(ctx context.Context, bankAccount *BankAccount) error

	// FindById returns the bank account with the given ID.
	FindById(ctx context.Context, id BankAccountId) (*BankAccount, error)
}

type LegacyRepository interface {
	GetBankAccounts(ctx context.Context) ([]*BankAccount, error)
	GetRawTransactionsByAccountId(ctx context.Context, accountID BankAccountId) ([]*RawTransaction, error)
	GetCategoryByOldLabel(ctx context.Context, oldLabel string) (*Category, error)
	GetDateParseRegexpByAccountId(ctx context.Context, accountID BankAccountId) (regexp.Regexp, error)
	GetCategoryFromMetadata(ctx context.Context, metadata map[string]interface{}) (*Category, error)
}
