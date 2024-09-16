package viewmodel

import "github.com/thomas-marquis/kleo-back/internal/domain"

type AppState struct {
	Transactions []domain.Transaction
	IsLoading    bool
	CurrentView  string
}

func NewState(initView string) *AppState {
	return &AppState{
		Transactions: make([]domain.Transaction, 0),
		CurrentView:  initView,
		IsLoading:    false,
	}
}
