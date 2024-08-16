package transactions_test

import "testing"

func Test_AllocationRatio_ShouldReturnErrorWhenInvalidNewRatio(t *testing.T) {
	t.Run("Ratio > 1", func(t *testing.T) {
		// Given
		tr := transactions.Transaction{}
		u := user.User{Name: "Toto"}

		// When
		res := tr.Allocate(u, 1.5)

		// Then
		assert.Error(t, res)
		assert.Equal(t, transactions.ErrInvalidAccolcationRate, res)
	})

	t.Run("Ratio < 0", func(t *testing.T) {
		// Given
		tr := transactions.Transaction{}
		u := user.User{Name: "Toto"}

		// When
		res := tr.Allocate(u, -0.5)

		// Then
		assert.Error(t, res)
		assert.Equal(t, transactions.ErrInvalidAccolcationRate, res)
	})
}
