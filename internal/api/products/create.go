package products

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleCreateProduct(w http.ResponseWriter, r *http.Request) {
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

	var req models.ProductDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	product := ConvertProductToDB(req)

	_, err = db.DB.Exec("INSERT INTO products (product_name, description, product_url, created_by, created_at, updated_at, organization, status, type, launch_date, metrics_url, logo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", &product.ProductName, &product.Description, &product.ProductURL, &product.CreatedBy, &product.CreatedAt, &product.UpdatedAt, orgID, &product.Status, &product.Type, &product.LaunchDate, &product.MetricsURL, &product.Logo)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
