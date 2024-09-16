package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/application"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/generated"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/mapping"
	"github.com/thomas-marquis/kleo-back/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type KleoServer struct {
	*generated.UnimplementedKeloAppServiceServer
	transactionSvc *application.TransactionService
}

var _ generated.KeloAppServiceServer = &KleoServer{}

func NewKleoServer() *KleoServer {
	return &KleoServer{}
}

func (s *KleoServer) ListTransactionsByUser(ctx context.Context, req *generated.ListTransactionsByUserRequest) (*generated.TransactionsListResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id format, uuid expected: %v", err)
	}

	transactions, err := s.transactionSvc.ListUserTransactions(ctx, domain.UserId(userId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "an error occurred when fetching transactions: %v", err)
	}

	var transactionsList = make([]*generated.Transaction, len(transactions))
	for _, transaction := range transactions {
		allocations := mapping.FromTransactionToAllocation(*transaction)

		var category *generated.Category
		if transaction.Category != nil {
			category, err = mapping.FromCategory(*transaction.Category)
			if err != nil {
				return nil, err
			}
		}

		transactionsList = append(transactionsList, &generated.Transaction{
			Id:          transaction.ID.String(),
			Amount:      float32(transaction.Amount),
			Label:       transaction.Label,
			Date:        timestamppb.New(transaction.Date),
			Allocations: allocations,
			Category:    category,
		})
	}

	response := &generated.TransactionsListResponse{
		Transactions: transactionsList,
	}

	return response, nil
}

func (s *KleoServer) GetUserById(ctx context.Context, req *generated.GetByIdRequest) (*generated.GetUserResponse, error) {
	return nil, nil
}
