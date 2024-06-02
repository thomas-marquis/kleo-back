package grpc

import (
	"github.com/thomas-marquis/kleo-back/internal/controller/grpc/generated"
	"github.com/thomas-marquis/kleo-back/internal/core/entity"
	"github.com/thomas-marquis/kleo-back/internal/core/value"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapSearchTransactionRequestToFilter(req *generated.SearchTransactionRequest) value.Filter {
	f := value.Filter{
		StartDate: req.Filter.StartDate.AsTime(),
		EndDate:   req.Filter.EndDate.AsTime(),
		MaxItems:  req.Filter.MaxItems,
		// User:      mapUserToEntity(req.Filter.User),
	}

	// f.AllCategories = make([]entities.Category, 0, len(req.Filter.CategoryIds))
	// for _, cid := range req.Filter.CategoryIds {
	// 	f.AllCategories = append(f.AllCategories, entities.Category{Id: cid})
	// }

	// f.AllCategoryTypes = make([]values.CategoryType, 0, len(req.Filter.CategoryIds))
	// for _, ct := range req.Filter.CategoryTypes {
	// 	val := ct.String()
	// 	if len(val) < 4 {
	// 		continue
	// 	}
	// 	val = val[4:] // remove "CAT_" prefix
	// 	ctype, ok := values.CategoryTypeFromValue(val)
	// 	if !ok {
	// 		continue
	// 	}
	// 	f.AllCategoryTypes = append(f.AllCategoryTypes, ctype)
	// }
	//
	// f.AllAccounts = make([]entities.BankAccount, 0, len(req.Filter.BankAccountIds))
	// for _, aid := range req.Filter.BankAccountIds {
	// 	f.AllAccounts = append(f.AllAccounts, entities.BankAccount{Id: aid})
	// }

	return f
}

func mapTransactionListToTransactionsListResponse(transactions []entity.Transaction, hasNext bool) *generated.TransactionsListResponse {
	nextPageToken := ""
	if hasNext {
		nextPageToken = "1"
	}

	msg := &generated.TransactionsListResponse{
		Transactions:  make([]*generated.Transaction, 0, len(transactions)),
		NextPageToken: nextPageToken,
	}

	for _, t := range transactions {
		msg.Transactions = append(msg.Transactions, mapTransactionToMessage(t))
	}

	return msg
}

// func mapUserToEntity(msg *generated.MsgUser) entities.User {
// 	return entities.User{
// 		Id:       msg.Id,
// 		UserName: msg.Username,
// 	}
// }

func mapTransactionToMessage(t entity.Transaction) *generated.Transaction {
	msg := &generated.Transaction{
		Id:     t.Id,
		Date:   timestamppb.New(t.Date),
		Amount: t.Amount,
		Label:  t.Label,
		// Category: mapCategoryToMessage(t.Category),
	}

	// msg.Tags = make([]*generated.MsgTag, 0, len(t.Tags))
	// for _, tag := range t.Tags {
	// 	msg.Tags = append(msg.Tags, mapTagToMessage(tag))
	// }

	return msg
}

// func mapCategoryToMessage(c entity.Category) *generated.MsgCategory {
// 	ctype := fmt.Sprintf("CAT_%s", c.Type.Value)
//
// 	return &generated.MsgCategory{
// 		Id:          c.Id,
// 		Label:       c.Label,
// 		Description: c.Description,
// 		Type:        generated.MsgCategoryType(generated.MsgCategoryType_value[ctype]),
// 	}
// }

// func mapTagToMessage(t entities.Tag) *generated.MsgTag {
// 	return &generated.MsgTag{
// 		Id:          t.Id,
// 		Label:       t.Label,
// 		Description: t.Description,
// 	}
// }
