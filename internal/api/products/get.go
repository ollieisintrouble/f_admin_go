package products

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.DB.Query("SELECT id, product_name, description, product_url, created_by, created_at, updated_at, status, type, launch_date, metrics_url, logo FROM products WHERE organization = $1", orgID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	var responses []models.ProductDTO
	for rows.Next() {
		var product db.Product
		if err := rows.Scan(&product.ID, &product.ProductName, &product.Description, &product.ProductURL, &product.CreatedBy, &product.CreatedAt, &product.UpdatedAt, &product.Status, &product.Type, &product.LaunchDate, &product.MetricsURL, &product.Logo); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		responses = append(responses, ConvertProductFromDB(product))
	}

	if len(responses) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No product found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, responses)
}
