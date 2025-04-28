package organizations

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertOrganizationToDB(org models.OrganizationDTO) db.Organization {
	return db.Organization{
		ID:              org.ID,
		Name:            org.Name,
		Image:           shared.StringToNullString(*org.Image),
		PurchasePackage: shared.StringToNullString(*org.PurchasePackage),
	}
}
