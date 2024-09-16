package viewmodel

import (
	"context"

	"fyne.io/fyne/v2/data/binding"
	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/desktop/data"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

type ViewModel struct {
	state *AppState
	repo  *data.KlepRepository

	transactionsbinding binding.UntypedList
	isLoadingBinding    binding.Bool
}

func New() *ViewModel {
	repo, _ := data.NewKleoRepository("localhost:50051")

	return &ViewModel{
		state:               NewState("transactions"),
		transactionsbinding: binding.NewUntypedList(),
		isLoadingBinding:    binding.NewBool(),
		repo:                repo,
	}
}

func (vm *ViewModel) Transactions() binding.UntypedList {
	return vm.transactionsbinding
}

func (vm *ViewModel) SearchTransactions() {
	vm.transactionsbinding.Set(make([]any, 0))

	transactions, _ := vm.repo.ListTransactions(context.TODO(), domain.UserId(uuid.New()), 10, 0)

	for _, t := range transactions {
		vm.transactionsbinding.Append(t)
	}
}

func (vm *ViewModel) IsLoading() binding.Bool {
	return vm.isLoadingBinding
}

func (vm *ViewModel) LoadingStart() {
	vm.isLoadingBinding.Set(true)
}

func (vm *ViewModel) LoadingEnd() {
	vm.isLoadingBinding.Set(false)
}
