package mapping

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/generated"
	"github.com/thomas-marquis/kleo-back/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FromMovmentType(m domain.MovmentType) (generated.MovmentType, error) {
	switch m {
	case domain.BaseIncome:
		return generated.MovmentType_BASE_INCOME, nil
	case domain.BaseExpense:
		return generated.MovmentType_BASE_EXPENSE, nil
	case domain.BaseSaving:
		return generated.MovmentType_BASE_SAVING, nil
	case domain.BaseTransfer:
		return generated.MovmentType_BASE_TRANSFER, nil
	default:
		return generated.MovmentType(0), errors.New("invalid movment type")
	}
}

func ToMovmentType(m generated.MovmentType) (domain.MovmentType, error) {
	switch m {
	case generated.MovmentType_BASE_INCOME:
		return domain.BaseIncome, nil
	case generated.MovmentType_BASE_EXPENSE:
		return domain.BaseExpense, nil
	case generated.MovmentType_BASE_SAVING:
		return domain.BaseSaving, nil
	case generated.MovmentType_BASE_TRANSFER:
		return domain.BaseTransfer, nil
	default:
		return domain.MovmentType(""), errors.New("invalid movment type")
	}
}

func FromCategory(c domain.Category) (*generated.Category, error) {
	m := c.SubCategory.MovmentType
	movType, err := FromMovmentType(m)
	if err != nil {
		return &generated.Category{}, status.Errorf(codes.InvalidArgument, "unknown movment type %s", m.String())
	}
	category := &generated.Category{
		Id:          c.ID.String(),
		Label:       c.Label,
		Value:       c.Value,
		Description: c.Description,
		SubCategory: &generated.SubCategory{
			Label:       c.SubCategory.Label,
			Value:       c.SubCategory.Value,
			MovmentType: movType,
		},
	}
	return category, nil
}

func ToCategory(c *generated.Category) (domain.Category, error) {
	sub, err := domain.GetSubCategoryFromValue(c.SubCategory.Value)
	if err != nil {
		return domain.Category{}, nil
	}

	catID, err := uuid.Parse(c.Id)
	if err != nil {
		return domain.Category{}, errors.New("invalid uuid format")
	}
	cat := domain.Category{
		ID:          domain.CategoryId(catID),
		Value:       c.Value,
		SubCategory: sub,
		Label:       c.Label,
		Description: c.Description,
	}
	return cat, nil
}

func FromTransactionToAllocation(t domain.Transaction) map[string]float32 {
	var allocations = make(map[string]float32)
	for userId, ratio := range t.Allocations {
		allocations[userId.String()] = float32(ratio)
	}
	return allocations
}

func ToTransaction(t *generated.Transaction) (domain.Transaction, error) {
	var cat domain.Category
	var err error
	if t.Category != nil {
		cat, err = ToCategory(t.Category)
		if err != nil {
			return domain.Transaction{}, err
		}
	}

	var allocations = make(map[domain.UserId]float64)
	for userId, ratio := range t.Allocations {
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			return domain.Transaction{}, nil
		}
		allocations[domain.UserId(userUUID)] = float64(ratio)
	}

	trID, err := uuid.Parse(t.Id)
	if err != nil {
		return domain.Transaction{}, errors.New("invalid transaction id")
	}

	tr := domain.Transaction{
		ID:          domain.TransactionId(trID),
		Category:    &cat,
		Amount:      float64(t.Amount),
		Label:       t.Label,
		Date:        t.Date.AsTime(),
		Allocations: allocations,
	}
	return tr, nil
}
