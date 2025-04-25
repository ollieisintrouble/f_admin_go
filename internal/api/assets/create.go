package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleCreateAsset(w http.ResponseWriter, r *http.Request) {
	var req models.AssetDTO
	var res models.AssetDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err := db.DB.QueryRow("INSERT INTO assets (title, cost, description, created_by, status, type, purchase_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", req.Title, req.Cost, req.Description, req.CreatedBy, req.Status, req.Type, req.PurchaseDate).Scan(&res.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, res)
}
