package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"net/http"
)

func handleCreateAsset(w http.ResponseWriter, r *http.Request) {
	var req db.Asset
	var res db.Asset

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err := db.DB.QueryRow("INSERT INTO assets (title, description, created_by) VALUES ($1, $2, $3) RETURNING id", req.Title, req.Description, req.CreatedBy).Scan(&res.Id)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, res)
}
