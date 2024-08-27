package domain

import "github.com/google/uuid"

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
)

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
