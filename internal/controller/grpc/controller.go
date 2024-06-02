package grpc

import (
	"context"
	"strconv"

	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/generated"
	"github.com/thomas-marquis/kleo-back/internal/core/port/service"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionController struct {
	generated.UnimplementedTransactionServiceServer
	transactionService service.TransactionService
}

func NewGrpcAdapter(transactionService service.TransactionService) TransactionController {
	return TransactionController{
		transactionService: transactionService,
	}
}

func (g *TransactionController) SearchTransactions(ctx context.Context, in *generated.SearchTransactionRequest) (*generated.TransactionsListResponse, error) {
	f := mapSearchTransactionRequestToFilter(in)

	page, err := strconv.Atoi(in.PageToken)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, hasNext, err := g.transactionService.Find(f, int32(page), in.PageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mapTransactionListToTransactionsListResponse(res, hasNext), nil
}

func (g *TransactionController) GetTransactionById(ctx context.Context, in *generated.GetTransactionByIdRequest) (*generated.GetTransactionByIdRResponse, error) {
	res, err := g.transactionService.FindById(in.Id)
	if err != nil {
		if err == value.ErrTransactionNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := generated.GetTransactionByIdRResponse{
		Transaction: mapTransactionToMessage(res),
	}
	return &resp, nil
}

var _ generated.TransactionServiceServer = &TransactionController{}
