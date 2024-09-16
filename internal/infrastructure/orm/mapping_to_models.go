package orm

import (
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

func ToAccountModel(a *domain.BankAccount) *Account {
	users := make([]*User, len(a.Users))
	for _, u := range a.Users {
		users = append(users, ToUserModel(u))
	}

	return &Account{
		ID:          uuid.UUID(a.ID),
		DisplayName: a.Label,
		IsActive:    a.IsActive,
		Users:       users,
	}
}

func ToRawTransactionModel(rt *domain.RawTransaction) *RawTransaction {
	return &RawTransaction{
		ID:     uuid.UUID(rt.ID),
		Label:  rt.Label,
		Date:   rt.Date,
		Amount: rt.Amount,
	}
}

func ToTransactionModel(t *domain.Transaction) *Transaction {
	allocations := make([]Allocation, 0)

	return &Transaction{
		ID:          uuid.UUID(t.ID),
		Label:       t.Label,
		Date:        t.Date,
		Amount:      t.Amount,
		Allocations: allocations,
	}
}

func ToUserModel(u *domain.User) *User {
	accounts := make([]*Account, 0)

	return &User{
		ID:       uuid.UUID(u.ID),
		UserName: u.Name,
		Email:    u.Email,
		Accounts: accounts,
	}
}
