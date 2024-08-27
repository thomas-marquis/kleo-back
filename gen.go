package main

import _ "go.uber.org/mock/gomock"

//go:generate protoc --proto_path=./proto --go_out=./internal/controller/grpc/generated --go_opt=paths=source_relative --go-grpc_out=./internal/controller/grpc/generated --go-grpc_opt=paths=source_relative transaction.proto
//go:generate mockgen -package mocks_domain -destination mocks/domain_repositories.go github.com/thomas-marquis/kleo-back/internal/domain TransactionRepository
