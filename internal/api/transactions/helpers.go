package transactions

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertTransaction(t db.Transaction) models.TransactionDTO {
	return models.TransactionDTO{
		ID:           t.ID,
		Amount:       t.Amount,
		Description:  shared.NullStringPtr(t.Description),
		Method:       shared.NullStringPtr(t.Method),
		CreatedBy:    shared.NullStringPtr(t.CreatedBy),
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		Status:       shared.NullStringPtr(t.Status),
		Type:         shared.NullStringPtr(t.Type),
		RecordedDate: shared.NullTimePtr(t.RecordedDate),
	}
}

// COMMENT OUT BECAUSE UNUSED
// func ConvertTransactionList(list []db.Transactions) []models.TransactionResponse {
// 	out := make([]models.TransactionResponse, len(list))
// 	for i, t := range list {
// 		out[i] = ConvertTransaction(t)
// 	}
// 	return out
// }
