package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleCreateAsset(w http.ResponseWriter, r *http.Request) {
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

	var req models.AssetDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	asset := ConvertAssetToDB(req)
	var res int64

	err = db.DB.QueryRow("INSERT INTO assets (title, cost, description, created_by, status, type, purchase_date, organization) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", asset.Title, asset.Cost, asset.Description, asset.CreatedBy, asset.Status, asset.Type, asset.PurchaseDate, orgID).Scan(&res)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	// description := fmt.Sprintf("Procurement of %s", asset.Title)
	// _, err = db.DB.Exec("INSERT INTO transactionns (amount, description, created_by, status, type, recorded_date, organization) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", asset.Cost, description, asset.CreatedBy, "Paid", "Asset", asset.PurchaseDate, orgID)
	// if err != nil {
	// 	shared.WriteError(w, http.StatusCreated, "Extra insertion of transaction failed")
	// }

	shared.WriteJSON(w, http.StatusCreated, res)
}
