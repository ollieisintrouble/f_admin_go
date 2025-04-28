package transactions

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handlePatchTransaction(w http.ResponseWriter, r *http.Request) {
	var req models.TransactionDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	transaction := ConvertTransactionToDB(req)

	_, err := db.DB.Exec("UPDATE transactions SET amount = $1, description = $2, method = $3, status = $4, type = $5, recorded_date = $6 WHERE id = $7", transaction.Amount, transaction.Description, transaction.Method, transaction.Status, transaction.Type, transaction.RecordedDate, transaction.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database update error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
