package products

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handlePatchProduct(w http.ResponseWriter, r *http.Request) {
	var req models.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	product := ConvertProductToDB(req)

	_, err := db.DB.Exec("UPDATE products SET product_name = $1, description = $2, product_url = $3, status = $4, type = $5, launch_date = $6, metrics_url = $7, logo = $8 WHERE id = $9", product.ProductName, product.Description, product.ProductURL, product.Status, product.Type, product.LaunchDate, product.MetricsURL, product.Logo, product.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database update error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
