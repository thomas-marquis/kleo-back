package domain

import (
	"errors"

	"github.com/google/uuid"
)

type MovmentType string

func (m *MovmentType) String() string {
	return string(*m)
}

var (
	BaseIncome   MovmentType = "base_income"
	BaseExpense  MovmentType = "base_expense"
	BaseTransfer MovmentType = "base_transfer"
	BaseSaving   MovmentType = "base_saving"
)

type SubCategory struct {
	Value       string
	MovmentType MovmentType
	Label       string
}

var (
	RequiredFixedExpense    = SubCategory{"required_fixed_expense", BaseExpense, "Dépense obligatoire fixe"}
	RequiredVariableExpense = SubCategory{"required_variable_expense", BaseExpense, "Dépense obligatoire variable"}
	NonRequiredExpense      = SubCategory{"non_required_expense", BaseExpense, "Dépense facultative"}
	Income                  = SubCategory{"income", BaseIncome, "Revenu"}
	Investment              = SubCategory{"investment", BaseSaving, "Investissement"}
	Transfer                = SubCategory{"transfer", BaseTransfer, "Transfert"}
)

func GetSubCategoryFromValue(value string) (SubCategory, error) {
	switch value {
	case "required_fixed_expense":
		return RequiredFixedExpense, nil
	case "required_variable_expense":
		return RequiredVariableExpense, nil
	case "non_required_expense":
		return NonRequiredExpense, nil
	case "income":
		return Income, nil
	case "investment":
		return Investment, nil
	case "transfer":
		return Transfer, nil
	default:
		return SubCategory{}, errors.New("invalid sub category value")
	}
}

type CategoryId uuid.UUID

func (i *CategoryId) String() string {
	return uuid.UUID(*i).String()
}

type Category struct {
	ID          CategoryId
	Value       string
	Label       string
	Description string
	SubCategory SubCategory
}

func NewCategory(value, label, description string, subCategory SubCategory) *Category {
	return &Category{
		ID:          CategoryId(uuid.New()),
		Value:       value,
		Label:       label,
		Description: description,
		SubCategory: subCategory,
	}
}
