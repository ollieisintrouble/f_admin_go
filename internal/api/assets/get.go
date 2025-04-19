package assets

import (
	"database/sql"
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetAsset(w http.ResponseWriter, r *http.Request) {
	var req models.GetAssetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.FindMany {
		rows, err := db.DB.Query("SELECT * FROM assets WHERE created_at BETWEEN $1 AND $2", req.StartDate, req.EndDate)
		if err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database query error")
			return
		}
		defer rows.Close()

		var assets []db.Asset
		for rows.Next() {
			var asset db.Asset
			if err := rows.Scan(&asset.Id, &asset.Title, &asset.Description, &asset.Cost, &asset.CreatedBy, &asset.CreatedAt, &asset.UpdatedAt); err != nil {
				shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
				return
			}
			assets = append(assets, asset)
		}

		if len(assets) == 0 {
			shared.WriteError(w, http.StatusNotFound, "No assets found")
			return
		}

		shared.WriteJSON(w, http.StatusOK, assets)
		return
	}

	var res db.Asset
	row := db.DB.QueryRow("SELECT * FROM assets WHERE id = $1", req.AssetId)
	if err := row.Scan(&res.Id, &res.Title, &res.Description, &res.Cost, &res.CreatedBy, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			shared.WriteError(w, http.StatusNotFound, "Asset not found")
			return
		}
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	shared.WriteJSON(w, http.StatusOK, res)
}
