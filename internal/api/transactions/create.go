package transactions

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req models.TransactionDTO
	var res models.TransactionDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err := db.DB.QueryRow("INSERT INTO transactionns (amount, description, method, created_by, status, type, recorded_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", req.Amount, req.Description, req.Method, req.CreatedBy, req.Status, req.Type, req.RecordedDate).Scan(&res.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, res)
}
