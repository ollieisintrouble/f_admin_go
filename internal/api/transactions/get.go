package transactions

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetTransaction(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM transactions")
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	var responses []models.TransactionDTO
	for rows.Next() {
		var transaction db.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Description, &transaction.Method, &transaction.CreatedBy, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.Status, &transaction.Type, &transaction.RecordedDate); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		responses = append(responses, ConvertTransaction(transaction))
	}

	if len(responses) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No transactions found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, responses)
}
