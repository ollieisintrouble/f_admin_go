package api

import (
	"database/sql"
	"encoding/json"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func getAsset(w http.ResponseWriter, r *http.Request) {
	var req models.GetAssetRequest
	var res db.Asset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	switch findMany := req.FindMany; findMany {
	case true:
		start := req.StartDate
		end := req.EndDate
		rows, err := db.DB.Query("SELECT * FROM assets WHERE created_at BETWEEN $1 AND $2", start, end)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var assets []db.Asset
		for rows.Next() {
			var asset db.Asset
			if err := rows.Scan(&asset.Id, &asset.Title, &asset.Description, &asset.CreatedAt); err != nil {
				http.Error(w, "Database scan error", http.StatusInternalServerError)
				return
			}
			assets = append(assets, asset)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Database rows error", http.StatusInternalServerError)
			return
		}
		if len(assets) == 0 {
			http.Error(w, "No assets found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(assets); err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case false:
		assetId := req.AssetId
		row := db.DB.QueryRow("SELECT * FROM assets WHERE id = $1", assetId)
		if err := row.Scan(&res.Id, &res.Title, &res.Description, &res.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Asset not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		}
	}
}

func createAsset(w http.ResponseWriter, r *http.Request) {
	var req db.Asset
	var res db.Asset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow("INSERT INTO assets (title, description, created_by) VALUES ($1, $2, $3) RETURNING id", req.Title, req.Description, req.CreatedBy).Scan(&res.Id)
	if err != nil {
		http.Error(w, "Database insert error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// TODO: PUT request for asset

// TODO: DELETE request for asset
