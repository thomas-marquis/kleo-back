package transactions

import "github.com/thomas-marquis/kleo-back/internal/user"

type Transactionrepository interface {
	GetAllocationByUser(tr Transaction, u user.User) (float32, error)
	SaveAllocation(tr Transaction, u user.User, alloc float32) error
}
