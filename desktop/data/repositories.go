package data

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/generated"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/mapping"
	"github.com/thomas-marquis/kleo-back/internal/domain"
	grpc_base "google.golang.org/grpc"
)

type KlepRepository struct {
	remoteClient generated.KeloAppServiceClient
}

func NewKleoRepository(addr string) (*KlepRepository, error) {
	conn, err := grpc_base.NewClient(addr)
	if err != nil {
		return nil, ErrDataTechnical
	}

	client := generated.NewKeloAppServiceClient(conn)
	return &KlepRepository{client}, nil
}

func (r *KlepRepository) ListTransactions(ctx context.Context, userId domain.UserId, limit, offset int) ([]domain.Transaction, error) {
	if false {
		req := &generated.ListTransactionsByUserRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			UserId: userId.String(),
		}
		resp, err := r.remoteClient.ListTransactionsByUser(ctx, req)
		if err != nil {
			return nil, err // TODO better handling
		}

		transactions := make([]domain.Transaction, len(resp.Transactions))
		for _, t := range resp.Transactions {
			tr, err := mapping.ToTransaction(t)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, tr)
		}
		return transactions, nil
	} else {
		return getFakeTransactions(), nil
	}
}

func getFakeTransactions() []domain.Transaction {
	// Define categories
	var (
		Groceries = domain.NewCategory("courses", "Courses", "Alimentation maison et autres courses", domain.RequiredVariableExpense)
		Salary    = domain.NewCategory("salaire", "Salaire", "Revenus réguliers", domain.Income)
		Saving    = domain.NewCategory("epargne", "Épargne", "Épargne et investissements", domain.Investment)
		Transfer  = domain.NewCategory("transfert", "Transfert", "Transfert d'argent entre comptes", domain.Transfer)
	)

	// Generate example transactions
	transactions := []domain.Transaction{
		{ID: domain.TransactionId(uuid.New()), Amount: 1200.50, Label: "Salary", Date: time.Now().AddDate(0, -1, -5), Allocations: generateAllocations(2), Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: -45.75, Label: "Groceries", Date: time.Now().AddDate(0, 0, -3), Allocations: generateAllocations(1), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: -100.00, Label: "Electricity Bill", Date: time.Now().AddDate(0, 0, -2), Allocations: nil, Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 500.00, Label: "Freelance Work", Date: time.Now().AddDate(0, -1, -10), Allocations: generateAllocations(3), Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: -75.50, Label: "Dining Out", Date: time.Now().AddDate(0, 0, -8), Allocations: generateAllocations(2), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 250.00, Label: "Savings Deposit", Date: time.Now().AddDate(0, -2, 0), Allocations: nil, Category: Saving},
		{ID: domain.TransactionId(uuid.New()), Amount: -60.00, Label: "Internet Bill", Date: time.Now().AddDate(0, 0, -12), Allocations: nil, Category: nil},
		{ID: domain.TransactionId(uuid.New()), Amount: -150.00, Label: "Clothing", Date: time.Now().AddDate(0, -1, -18), Allocations: generateAllocations(1), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 300.00, Label: "Sold Old Furniture", Date: time.Now().AddDate(0, -2, -5), Allocations: nil, Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: -20.00, Label: "Gym Membership", Date: time.Now().AddDate(0, 0, -15), Allocations: generateAllocations(2), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: -500.00, Label: "Rent Payment", Date: time.Now().AddDate(0, -1, 0), Allocations: generateAllocations(3), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 100.00, Label: "Gift Received", Date: time.Now().AddDate(0, 0, -20), Allocations: nil, Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: -35.00, Label: "Books Purchase", Date: time.Now().AddDate(0, -1, -3), Allocations: generateAllocations(1), Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 150.00, Label: "Stock Dividends", Date: time.Now().AddDate(0, -3, 0), Allocations: generateAllocations(2), Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: 200.00, Label: "Savings Interest", Date: time.Now().AddDate(0, -2, -2), Allocations: nil, Category: Saving},
		{ID: domain.TransactionId(uuid.New()), Amount: -45.00, Label: "Transportation", Date: time.Now().AddDate(0, 0, -1), Allocations: nil, Category: nil},
		{ID: domain.TransactionId(uuid.New()), Amount: -250.00, Label: "Vacation Fund", Date: time.Now().AddDate(0, -4, 0), Allocations: generateAllocations(3), Category: Saving},
		{ID: domain.TransactionId(uuid.New()), Amount: 120.00, Label: "Sold Gadgets", Date: time.Now().AddDate(0, -1, -12), Allocations: generateAllocations(2), Category: Groceries},
		{ID: domain.TransactionId(uuid.New()), Amount: -15.00, Label: "Coffee Shop", Date: time.Now().AddDate(0, 0, -6), Allocations: nil, Category: Salary},
		{ID: domain.TransactionId(uuid.New()), Amount: 500.00, Label: "Account Transfer", Date: time.Now().AddDate(0, -1, -7), Allocations: nil, Category: Transfer},
	}

	return transactions
}

// Helper function to generate random allocations that sum up to 1
func generateAllocations(n int) map[domain.UserId]float64 {
	allocations := make(map[domain.UserId]float64)
	total := 1.0

	for i := 0; i < n; i++ {
		if i == n-1 {
			allocations[domain.UserId(uuid.New())] = total // Ensure sum of allocations equals 1
		} else {
			ratio := rand.Float64() * total
			allocations[domain.UserId(uuid.New())] = ratio
			total -= ratio
		}
	}

	return allocations
}
