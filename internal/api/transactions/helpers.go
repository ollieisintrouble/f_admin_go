package transactions

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertTransactionFromDB(t db.Transaction) models.TransactionDTO {
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

func ConvertTransactionToDB(t models.TransactionDTO) db.Transaction {
	return db.Transaction{
		ID:           t.ID,
		Amount:       t.Amount,
		Description:  shared.StringToNullString(*t.Description),
		Method:       shared.StringToNullString(*t.Method),
		CreatedBy:    shared.StringToNullString(*t.CreatedBy),
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		Status:       shared.StringToNullString(*t.Status),
		Type:         shared.StringToNullString(*t.Type),
		RecordedDate: shared.TimeToNullTime(t.RecordedDate),
	}
}
