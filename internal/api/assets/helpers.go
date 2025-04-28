package assets

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertAssetFromDB(t db.Asset) models.AssetDTO {
	return models.AssetDTO{
		ID:           t.ID,
		Title:        t.Title,
		Cost:         t.Cost,
		Description:  shared.NullStringPtr(t.Description),
		CreatedBy:    shared.NullStringPtr(t.CreatedBy),
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		Status:       shared.NullStringPtr(t.Status),
		Type:         shared.NullStringPtr(t.Type),
		PurchaseDate: shared.NullTimePtr(t.PurchaseDate),
	}
}

func ConvertAssetToDB(t models.AssetDTO) db.Asset {
	return db.Asset{
		ID:           t.ID,
		Title:        t.Title,
		Cost:         t.Cost,
		Description:  shared.StringToNullString(*t.Description),
		CreatedBy:    shared.StringToNullString(*t.CreatedBy),
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		Status:       shared.StringToNullString(*t.Status),
		Type:         shared.StringToNullString(*t.Type),
		PurchaseDate: shared.TimeToNullTime(t.PurchaseDate),
	}
}
