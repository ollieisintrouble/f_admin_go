package assets

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetAsset(w http.ResponseWriter, r *http.Request) {
	orgIDStr, ok := shared.GetOrgID(r.Context())
	if !ok {
		shared.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	orgID, err := shared.ConvertOrgIDToInt(orgIDStr)
	if err != nil {
		shared.WriteError(w, http.StatusUnauthorized, "Invalid org ID")
		return
	}

	rows, err := db.DB.Query("SELECT id, title, cost, description, created_by, created_at, updated_at, status, type, purchase_date FROM assets WHERE organization = $1", orgID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	var response []models.AssetDTO
	for rows.Next() {
		var asset db.Asset
		if err := rows.Scan(&asset.ID, &asset.Title, &asset.Cost, &asset.Description, &asset.CreatedBy, &asset.CreatedAt, &asset.UpdatedAt, &asset.Status, &asset.Type, &asset.PurchaseDate); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		response = append(response, ConvertAssetFromDB(asset))
	}

	if len(response) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No assets found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, response)
}
