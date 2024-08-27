package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

func newTransactionId() domain.TransactionId {
	return domain.TransactionId(uuid.New())
}

func Test_Allocate_ShouldReturnErrorWhenInvalidNewRatio(t *testing.T) {
	t.Run("Ratio > 1", func(t *testing.T) {
		// Given
		tr := domain.NewTransaction(100.0, "some expense", time.Now())

		// When
		err := tr.Allocate(newUserId(), 1.5)

		// Then
		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidAccolcationRate, err)
	})

	t.Run("Ratio < 0", func(t *testing.T) {
		// Given
		tr := domain.NewTransaction(100.0, "some expense", time.Now())

		// When
		err := tr.Allocate(newUserId(), -0.5)

		// Then
		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidAccolcationRate, err)
	})
}

func Test_Allocate_ShouldUpdateAllocations(t *testing.T) {
	// Given
	userId := newUserId()
	tr := domain.NewTransaction(100.0, "some expense", time.Now())

	// When
	err := tr.Allocate(userId, 0.5)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0.5, tr.GetUserAllocation(userId))
}

func Test_Allocate_ShouldReturnErrorWhenSumOfRatiosIsGreaterThanOne(t *testing.T) {
	// Given
	tr := domain.NewTransaction(100.0, "some expense", time.Now())
	tr.Allocate(newUserId(), 0.5)

	// When
	err := tr.Allocate(newUserId(), 0.6)

	// Then
	assert.Error(t, err)
	assert.Equal(t, domain.ErrOverAllocation, err)
}

func Test_Allocate_ShouldAllocateWhenAllocationAlreadyExistsForUser(t *testing.T) {
	// Given
	userId := newUserId()
	tr := domain.NewTransaction(100.0, "some expense", time.Now())
	tr.Allocate(newUserId(), 0.5)
	tr.Allocate(userId, 0.3)

	// When
	err := tr.Allocate(userId, 0.5)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0.5, tr.GetUserAllocation(userId))
}

func Test_GetUserAllocation_ShouldReturnZeroWhenUserNotAllocated(t *testing.T) {
	// Given
	tr := domain.NewTransaction(100.0, "some expense", time.Now())

	// When
	ratio := tr.GetUserAllocation(newUserId())

	// Then
	assert.Equal(t, 0.0, ratio)
}

func Test_GetUserAllocation_ShouldReturnUserAllocation(t *testing.T) {
	// Given
	userId1 := newUserId()
	userId2 := newUserId()
	tr := domain.NewTransaction(100.0, "some expense", time.Now())
	tr.Allocate(userId1, 0.5)
	tr.Allocate(userId2, 0.4)

	// When
	ratio1 := tr.GetUserAllocation(userId1)
	ratio2 := tr.GetUserAllocation(userId2)

	// Then
	assert.Equal(t, 0.5, ratio1)
	assert.Equal(t, 0.4, ratio2)
}
