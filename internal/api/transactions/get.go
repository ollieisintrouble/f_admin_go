package transactions

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetTransaction(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.DB.Query("SELECT id, amount, description, method, created_by, created_at, updated_at, status, type, recorded_date FROM transactions WHERE organization = $1", orgID)
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
		responses = append(responses, ConvertTransactionFromDB(transaction))
	}

	if len(responses) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No transactions found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, responses)
}
