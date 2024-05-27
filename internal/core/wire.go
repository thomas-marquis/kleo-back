package core

import (
	"github.com/google/wire"
	"github.com/thomas-marquis/kleo-back/internal/core/port/service"
	impl "github.com/thomas-marquis/kleo-back/internal/core/service"
)

var (
	diSet = wire.NewSet(
		wire.Bind(new(service.TransactionService), new(*impl.TransactionServiceImpl)),
		impl.NewTransactionServiceImpl,
	)
)
