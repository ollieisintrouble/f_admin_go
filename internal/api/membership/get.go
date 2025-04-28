package membership

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"net/http"
)

type UserIDRequestBody struct {
	UserID string `json:"userId"`
}

func handleGetMembership(w http.ResponseWriter, r *http.Request) {
	var req UserIDRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var res []db.Membership

	rows, err := db.DB.Query("SELECT * FROM memberships WHERE user_id = $1", req.UserID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var membership db.Membership
		if err := rows.Scan(&membership.UserID, &membership.OrganizationID, &membership.Role); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		res = append(res, membership)
	}

	if len(res) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No memberships found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, res)
}
