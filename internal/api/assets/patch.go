package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handlePatchAsset(w http.ResponseWriter, r *http.Request) {
	var req models.AssetDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	transaction := ConvertAssetToDB(req)

	_, err := db.DB.Exec("UPDATE assets SET title = $1, cost = $2, description = $3, status = $4, type = $5, purchase_date = $6 WHERE id = $7", transaction.Title, transaction.Cost, transaction.Description, transaction.Status, transaction.Type, transaction.PurchaseDate, transaction.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database update error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
