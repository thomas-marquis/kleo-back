package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/port/repository"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
	"github.com/thomas-marquis/kleo-back/internal/infra/utils"
)

type SQLTransactionRepository struct {
	db *sql.DB
}

func NewSQLTransactionRepository(db *sql.DB) *SQLTransactionRepository {
	return &SQLTransactionRepository{
		db: db,
	}
}

func (r *SQLTransactionRepository) FindTransactionsByFilter(filter value.Filter, size, offset int32) ([]entity.Transaction, error) {
	builder := newSearchSQLQueryBuilder()
	builder.WithStartDate(filter.StartDate)
	builder.WithEndDate(filter.EndDate)
	// builder.WithCategories(filter.AllCategories)
	// builder.WithCategoryTypes(filter.AllCategoryTypes)
	// builder.WithAccounts(filter.AllAccounts)
	// builder.WithUser(filter.User)
	builder.WithOffsetAndLimit(offset, size)
	sqlQuery, args := builder.Build()

	tr, err := utils.SelectMany(
		r.db,
		sqlQuery,
		func(row *sql.Rows) (entity.Transaction, error) {
			var tr entity.Transaction
			// var cat entity.Category
			if err := row.Scan(&tr.Id, &tr.Label, &tr.Date, &tr.Amount); err != nil {
				// if err := row.Scan(&tr.Id, &tr.Label, &tr.Date, &tr.Amount, &cat.Id, &cat.Label, &cat.Type.Value); err != nil {
				return entity.Transaction{}, err
			}
			// tr.Category = cat
			return tr, nil
		},
		args...,
	)

	if err != nil {
		return nil, err
	}

	return tr, nil
}

func (r *SQLTransactionRepository) FindTransactionById(id string) (entity.Transaction, error) {
	tr, err := utils.SelectOneOrEmpty(
		r.db,
		`select t.id, t.label, t.date, t.amount, c.id, c.label, c.type, a.id, a.display_name
		from transactions t 
		left join categories c on c.id = t.category_id 
		left join accounts a on a.id = t.account_id
		where t.id = ?`,
		func(r *sql.Row) (entity.Transaction, error) {
			var tr entity.Transaction
			// var cat entity.Category
			// var acc entity.BankAccount

			if err := r.Scan(&tr.Id, &tr.Label, &tr.Date, &tr.Amount); err != nil {
				// if err := r.Scan(&tr.Id, &tr.Label, &tr.Date, &tr.Amount, &cat.Id, &cat.Label, &cat.Type.Value, &acc.Id, &acc.DisplayName); err != nil {
				return entity.Transaction{}, err
			}
			// tr.Category = cat
			// tr.Account = acc
			return tr, nil
		},
		entity.Transaction{},
		id)
	if err != nil {
		return entity.Transaction{}, err
	}
	if tr.Id == "" {
		return entity.Transaction{}, value.ErrTransactionNotFound
	}

	return tr, nil
}

var _ repository.TransactionRepository = &SQLTransactionRepository{}

type searchSQLQueryBuilder struct {
	query    string
	args     []interface{}
	maxItems int32
	offset   int32
}

func newSearchSQLQueryBuilder() *searchSQLQueryBuilder {
	var baseQuery = `select t.id, t.label, t.date, t.amount, c.id, c.label, c.type
from transactions t
inner join categories c on c.id = t.category_id
inner join accounts_users_bindings au on au.account_id = t.account_id
where `

	return &searchSQLQueryBuilder{
		query:    baseQuery,
		args:     make([]interface{}, 0),
		maxItems: -1,
	}
}

func (b *searchSQLQueryBuilder) WithStartDate(date time.Time) *searchSQLQueryBuilder {
	if date.IsZero() {
		return b
	}

	if len(b.args) > 0 {
		b.query += "and "
	}
	b.query += "t.date >= ? "
	b.args = append(b.args, date)

	return b
}

func (b *searchSQLQueryBuilder) WithEndDate(date time.Time) *searchSQLQueryBuilder {
	if date.IsZero() {
		return b
	}

	if len(b.args) > 0 {
		b.query += "and "
	}
	b.query += "t.date <= ? "
	b.args = append(b.args, date)

	return b
}

// func (b *searchSQLQueryBuilder) WithCategories(categories []entities.Category) *searchSQLQueryBuilder {
// 	if len(categories) == 0 {
// 		return b
// 	}
//
// 	if len(b.args) > 0 {
// 		b.query += "and "
// 	}
// 	b.query += "t.category_id in ("
//
// 	for i, c := range categories {
// 		b.query += "?"
// 		b.args = append(b.args, c.Id)
// 		if i < len(categories)-1 {
// 			b.query += ", "
// 		}
// 	}
// 	b.query += ") "
//
// 	return b
// }
//
// func (b *searchSQLQueryBuilder) WithCategoryTypes(types []values.CategoryType) *searchSQLQueryBuilder {
// 	if len(types) == 0 {
// 		return b
// 	}
//
// 	if len(b.args) > 0 {
// 		b.query += "and "
// 	}
// 	b.query += "c.type in ("
// 	for i, t := range types {
// 		b.query += "?"
// 		b.args = append(b.args, t.Value)
// 		if i < len(types)-1 {
// 			b.query += ", "
// 		}
// 	}
// 	b.query += ") "
//
// 	return b
// }
//
// func (b *searchSQLQueryBuilder) WithAccounts(accounts []entities.BankAccount) *searchSQLQueryBuilder {
// 	if len(accounts) == 0 {
// 		return b
// 	}
//
// 	if len(b.args) > 0 {
// 		b.query += "and "
// 	}
// 	b.query += "t.account_id in ("
// 	for i, a := range accounts {
// 		b.query += "?"
// 		b.args = append(b.args, a.Id)
// 		if i < len(accounts)-1 {
// 			b.query += ", "
// 		}
// 	}
// 	b.query += ") "
//
// 	return b
// }
//
// func (b *searchSQLQueryBuilder) WithUser(user entities.User) *searchSQLQueryBuilder {
// 	if user.Id == "" {
// 		return b
// 	}
//
// 	if len(b.args) > 0 {
// 		b.query += "and "
// 	}
// 	b.query += "au.user_id = ? "
// 	b.args = append(b.args, user.Id)
//
// 	return b
// }

func (b *searchSQLQueryBuilder) WithMaxItems(maxItems int32) *searchSQLQueryBuilder {
	if maxItems == 0 {
		return b
	}

	b.maxItems = maxItems

	return b
}

func (b *searchSQLQueryBuilder) WithOffsetAndLimit(offset, limit int32) *searchSQLQueryBuilder {
	if limit == 0 {
		return b
	}
	if offset == 0 {
		return b.WithMaxItems(limit)
	}

	b.maxItems = limit
	b.offset = offset

	return b
}

func (b *searchSQLQueryBuilder) Build() (string, []interface{}) {
	b.query = strings.TrimSpace(b.query)
	b.query += "\norder by t.date desc"
	if b.maxItems > 0 && b.offset > 0 {
		b.query += fmt.Sprintf("\noffset %d rows\nfetch first %d rows only", b.offset, b.maxItems)
	} else if b.maxItems > 0 {
		b.query += fmt.Sprintf("\nlimit %d", b.maxItems)
	}
	return b.query, b.args
}
