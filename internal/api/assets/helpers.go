package assets

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertAsset(t db.Asset) models.AssetDTO {
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
