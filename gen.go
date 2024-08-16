package main

import _ "go.uber.org/mock/gomock"

//go:generate mockgen -package mocks -destination mocks/transactions.go github.com/thomas-marquis/kleo-back/internal/transactions Transactionrepository
