package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	DisplayName     string
	IsActive        bool
	Transactions    []Transaction    `gorm:"foreignKey:AccountID"`
	RawTransactions []RawTransaction `gorm:"foreignKey:AccountID"`
	Users           []*User          `gorm:"many2many:user_accounts"`
}

type RawTransaction struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Label     string
	Date      time.Time
	Amount    float64
	AccountID uuid.UUID
}

type Allocation struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Rate          float64
	TransactionId uuid.UUID
	UserID        uuid.UUID
}

type Transaction struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Label       string
	Date        time.Time
	Amount      float64
	AccountID   uuid.UUID
	Allocations []Allocation `gorm:"foreignKey:TransactionId"`
}

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserName    string
	Email       string
	Allocations []Allocation `gorm:"foreignKey:UserID"`
	Accounts    []*Account   `gorm:"many2many:user_accounts"`
}
