package products

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertProductFromDB(p db.Product) models.ProductDTO {
	return models.ProductDTO{
		ID:          p.ID,
		ProductName: p.ProductName,
		Description: shared.NullStringPtr(p.Description),
		ProductURL:  shared.NullStringPtr(p.ProductURL),
		CreatedBy:   shared.NullStringPtr(p.CreatedBy),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Status:      shared.NullStringPtr(p.Status),
		Type:        shared.NullStringPtr(p.Type),
		LaunchDate:  shared.NullTimePtr(p.LaunchDate),
		MetricsURL:  shared.NullStringPtr(p.MetricsURL),
		Logo:        shared.NullStringPtr(p.Logo),
	}
}

func ConvertProductToDB(p models.ProductDTO) db.Product {
	return db.Product{
		ID:          p.ID,
		ProductName: p.ProductName,
		Description: shared.StringToNullString(*p.Description),
		ProductURL:  shared.StringToNullString(*p.ProductURL),
		CreatedBy:   shared.StringToNullString(*p.CreatedBy),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Status:      shared.StringToNullString(*p.Status),
		Type:        shared.StringToNullString(*p.Type),
		LaunchDate:  shared.TimeToNullTime(p.LaunchDate),
		MetricsURL:  shared.StringToNullString(*p.MetricsURL),
		Logo:        shared.StringToNullString(*p.Logo),
	}
}
