package core

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc"
	"github.com/thomas-marquis/kleo-back/internal/core/port/repository"
	"github.com/thomas-marquis/kleo-back/internal/core/port/service"
	implSvc "github.com/thomas-marquis/kleo-back/internal/core/service"
	implRep "github.com/thomas-marquis/kleo-back/internal/infra/repository"
)

var (
	diSet = wire.NewSet(
		wire.Bind(new(service.TransactionService), new(*implSvc.TransactionServiceImpl)),
		implSvc.NewTransactionServiceImpl,
		wire.Bind(new(repository.TransactionRepository), new(*implRep.SQLTransactionRepository)),
		implRep.NewSQLTransactionRepository,
	)
)

func InjectGrpc(bd *sql.DB) grpc.TransactionController {
	wire.Build(diSet)

	return grpc.TransactionController{}
}
