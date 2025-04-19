package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"net/http"
)

func handleDeleteAsset(w http.ResponseWriter, r *http.Request) {
	var req db.Asset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.DB.Exec("DELETE FROM assets WHERE id = $1", req.Id)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database delete error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
