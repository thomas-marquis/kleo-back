package transactions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomas-marquis/kleo-back/internal/transactions"
	"github.com/thomas-marquis/kleo-back/internal/user"
	"github.com/thomas-marquis/kleo-back/mocks"
	"go.uber.org/mock/gomock"
)

func Test_Allocate_ShouldReturnErrorWhenInvalidNewRatio(t *testing.T) {
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

func Test_Allocate_ShouldSaveAllocationWhenValidAndNoOneAlreadyAffected(t *testing.T) {
	// Given
	u := user.User{Name: "Toto"}
	ctrl := gomock.NewController(t)
	tRepo := mocks.NewMockTransactionrepository(ctrl)
	tr, _ := transactions.InitTransaction(&transactions.Transaction{}, tRepo)

	tRepo.EXPECT().SaveAllocation(*tr, u, 0.5).Return(nil)

	// When
	res := tr.Allocate(u, 0.5)

	// Then
	assert.NoError(t, res)
}

// func Test_Allocate_ShouldReturnErrorWhenFailToGetAllocations(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	tRepo := mocks.NewMockTransactionrepository(ctrl)
// 	suite.allocRepository.EXPECT().GetByTransactionID("111").Return(nil, errors.New("ckc"))
//
// 	account := entities.BankAccount{Id: "666"}
// 	tr := entities.Transaction{
// 		Id:      "111",
// 		Date:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
// 		Amount:  100,
// 		Account: account,
// 	}
// 	user := entities.User{Id: "222", UserName: "thomas", Household: &entities.Household{Id: "333", Name: "marquis"}}
//
// 	// When
// 	res := suite.service.Allocate(tr, user, 0.5)
//
// 	// Then
// 	assert.Error(suite.T(), res)
// 	assert.Equal(suite.T(), "an error occurred while getting allocations for transaction 111: ckc", res.Error())
// }
