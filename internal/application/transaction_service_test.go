package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thomas-marquis/kleo-back/internal/application"
	"github.com/thomas-marquis/kleo-back/internal/domain"
	mocks_domain "github.com/thomas-marquis/kleo-back/mocks"
	"go.uber.org/mock/gomock"
)

func getTransactionServiceMocks(t *testing.T) (*gomock.Controller, *mocks_domain.MockTransactionRepository) {
	ctrl := gomock.NewController(t)
	mRepo := mocks_domain.NewMockTransactionRepository(ctrl)
	return ctrl, mRepo
}

func Test_AllocateToUser_ShouldSaveTransactionAfterAllocation(t *testing.T) {
	// Given
	_, mRepo := getTransactionServiceMocks(t)
	svc := application.NewTransactionService(mRepo)
	tr := domain.Transaction{Allocations: make(map[domain.UserId]float64)}
	userID := domain.UserId{}
	ctx := context.TODO()

	mRepo.EXPECT().Save(ctx, &tr).Return(nil)

	// When
	err := svc.AllocateToUser(ctx, &tr, userID, 0.5)

	// Then
	assert.NoError(t, err)
}

func Test_AllocateToUser_ShouldReturnErrorWhenSaveFails(t *testing.T) {
	// Given
	_, mRepo := getTransactionServiceMocks(t)
	svc := application.NewTransactionService(mRepo)
	tr := domain.Transaction{Allocations: make(map[domain.UserId]float64)}
	userID := domain.UserId{}
	ctx := context.TODO()

	mRepo.EXPECT().Save(ctx, &tr).Return(errors.New("ckc"))

	// When
	err := svc.AllocateToUser(ctx, &tr, userID, 0.5)

	// Then
	assert.Error(t, err)
	assert.Equal(t, "an error occurred when saving transaction: ckc", err.Error())
}

func Test_AllocateToUser_ShouldAllocateTransactionToUser(t *testing.T) {
	// Given
	_, mRepo := getTransactionServiceMocks(t)
	svc := application.NewTransactionService(mRepo)
	tr := domain.Transaction{Allocations: make(map[domain.UserId]float64)}
	userID := domain.UserId(uuid.New())
	ctx := context.TODO()

	mRepo.EXPECT().Save(ctx, &tr).Return(nil).Do(func(ctx context.Context, tr *domain.Transaction) {
		assert.Equal(t, 0.5, tr.GetUserAllocation(userID))
	})

	// When
	err := svc.AllocateToUser(ctx, &tr, userID, 0.5)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0.5, tr.GetUserAllocation(userID))
}

func Test_AllocateToUser_ShouldReturnErrorWhenInvalidNewRatio(t *testing.T) {
	// Given
	_, mRepo := getTransactionServiceMocks(t)
	svc := application.NewTransactionService(mRepo)
	tr := domain.Transaction{Allocations: make(map[domain.UserId]float64)}
	userID := domain.UserId(uuid.New())
	ctx := context.TODO()

	mRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// When
	err := svc.AllocateToUser(ctx, &tr, userID, 1.5)

	// Then
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidAccolcationRate, err)
}

func Test_AllocateToUser_ShouldReturnErrorWhenSumOfRatiosIsGreaterThanOne(t *testing.T) {
	// Given
	_, mRepo := getTransactionServiceMocks(t)
	svc := application.NewTransactionService(mRepo)
	tr := domain.Transaction{Allocations: make(map[domain.UserId]float64)}
	tr.Allocate(domain.UserId(uuid.New()), 0.5)
	userID := domain.UserId(uuid.New())
	ctx := context.TODO()

	mRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// When
	err := svc.AllocateToUser(ctx, &tr, userID, 0.6)

	// Then
	assert.Error(t, err)
	assert.Equal(t, domain.ErrOverAllocation, err)
}
