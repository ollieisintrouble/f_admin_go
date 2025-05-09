package transactions

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"fmt"
	"net/http"
)

func handleCreateTransaction(w http.ResponseWriter, r *http.Request) {
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

	var req models.TransactionForm

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error decoding JSON:", err)
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	transaction := ConvertTransactionToCreate(req)

	_, err = db.DB.Exec("INSERT INTO transactions (amount, description, method, created_by, status, type, recorded_date, organization) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", transaction.Amount, transaction.Description, transaction.Method, transaction.CreatedBy, transaction.Status, transaction.Type, transaction.RecordedDate, orgID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
