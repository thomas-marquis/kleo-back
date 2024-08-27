package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionId uuid.UUID

func (i TransactionId) String() string {
	return uuid.UUID(i).String()
}

type Transaction struct {
	ID          TransactionId
	Amount      float64
	Label       string
	Date        time.Time
	Allocations map[UserId]float64
	Category    *Category
}

// Create a new transaction with the specified amount, label and date
func NewTransaction(amount float64, label string, date time.Time) *Transaction {
	return &Transaction{
		ID:          TransactionId(uuid.New()),
		Amount:      amount,
		Label:       label,
		Date:        date,
		Allocations: make(map[UserId]float64),
		Category:    nil,
	}
}

// Perform allocation for the specified user.
// Return ErrOverAllocation if the specified allocation sumed to existing ones are greater than 1
// Return ErrInvalidAccolcationRate if the specified allocation rate is not between 0 and 1
func (t *Transaction) Allocate(userID UserId, ratio float64) error {
	if ratio > 1 || ratio < 0 {
		return ErrInvalidAccolcationRate
	}

	totAllocRate := ratio
	for uId, alloc := range t.Allocations {
		if uId == userID {
			continue
		}
		totAllocRate += alloc
	}
	if totAllocRate > 1.0 {
		return ErrOverAllocation
	}

	t.Allocations[userID] = ratio
	return nil
}

// Get the allocation ratio for the specified user or 0 if not found
func (t *Transaction) GetUserAllocation(userID UserId) float64 {
	if ratio, ok := t.Allocations[userID]; ok {
		return ratio
	}
	return 0
}

func (t *Transaction) UpdateCategory(category *Category) {
	t.Category = category
}
