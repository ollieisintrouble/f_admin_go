package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"net/http"
)

func handlePatchAsset(w http.ResponseWriter, r *http.Request) {
	var req db.Asset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.DB.Exec("UPDATE assets SET title = $1, cost = $2, description = $3 WHERE id = $4", req.Title, req.Cost, req.Description, req.Id)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database update error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
