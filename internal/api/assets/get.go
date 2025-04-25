package assets

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetAsset(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM assets")
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
		response = append(response, ConvertAsset(asset))
	}

	if len(response) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No assets found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, response)
}
