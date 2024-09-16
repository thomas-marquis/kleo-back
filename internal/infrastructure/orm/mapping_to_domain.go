package orm

import "github.com/thomas-marquis/kleo-back/internal/domain"

func ToBankAccount(a *Account) *domain.BankAccount {
	users := make(map[domain.UserId]*domain.User, len(a.Users))
	for _, u := range a.Users {
		users[domain.UserId(u.ID)] = ToUser(u)
	}

	return &domain.BankAccount{
		ID:    domain.BankAccountId(a.ID),
		Label: a.DisplayName,
		Users: users,
	}
}

func ToRawTransaction(t *RawTransaction) *domain.RawTransaction {
	return &domain.RawTransaction{
		ID:     domain.RawTransactionId(t.ID),
		Label:  t.Label,
		Date:   t.Date,
		Amount: t.Amount,
	}
}

func ToTransaction(t *Transaction) *domain.Transaction {
	allocations := make(map[domain.UserId]float64, len(t.Allocations))
	for _, a := range t.Allocations {
		allocations[domain.UserId(a.UserID)] = a.Rate
	}

	return &domain.Transaction{
		ID:          domain.TransactionId(t.ID),
		Label:       t.Label,
		Date:        t.Date,
		Amount:      t.Amount,
		Allocations: allocations,
		// TODO: category
	}
}

func ToUser(u *User) *domain.User {
	accounts := make([]domain.BankAccount, len(u.Accounts))
	for i, a := range u.Accounts {
		accounts[i] = *ToBankAccount(a)
	}
	return &domain.User{
		ID:    domain.UserId(u.ID),
		Name:  u.UserName,
		Email: u.Email,
	}
}
