package main

import "github.com/thomas-marquis/kleo-back/cmd"

//go:generate protoc --proto_path=./proto --go_out=./internal/controller/grpc/generated --go_opt=paths=source_relative --go-grpc_out=./internal/controller/grpc/generated --go-grpc_opt=paths=source_relative transaction.proto

func main() {
	cmd.Execute()
}
