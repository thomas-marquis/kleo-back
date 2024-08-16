package transactions

import (
	"time"

	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/user"
)

type Transaction struct {
	ID     uuid.UUID
	Amount float64
	Label  string
	Date   time.Time

	reporitory Transactionrepository
}

func InitTransaction(t *Transaction, r Transactionrepository) (*Transaction, error) {
	t.reporitory = r
	return t, nil
}

func (t *Transaction) Categorize(c Category) error {
	return nil
}

func (t *Transaction) AddTag(tag Tag) error {
	return nil
}

func (t *Transaction) RemoveTag(tag Tag) error {
	return nil
}

func (t *Transaction) Allocate(user user.User, allocation float32) error {
	if allocation < 0.0 || allocation > 1.0 {
		return ErrInvalidAccolcationRate
	}

	t.reporitory.SaveAllocation(*t, user, float64(allocation))

	return nil
}

func (t *Transaction) GetUserAllocation(user user.User) (float32, error) {
	return 0.0, nil
}
