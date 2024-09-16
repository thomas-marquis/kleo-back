package domain

import "github.com/google/uuid"

type BankAccountId uuid.UUID

type BankAccount struct {
	ID       BankAccountId
	Label    string
	Users    map[UserId]*User
	IsActive bool
}

func NewBankAccount(label string) *BankAccount {
	return &BankAccount{
		ID:       BankAccountId(uuid.New()),
		Label:    label,
		Users:    make(map[UserId]*User),
		IsActive: true,
	}
}

func (b *BankAccount) AssociateUser(u *User) {
	b.Users[u.ID] = u
}
