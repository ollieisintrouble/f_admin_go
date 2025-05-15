package membership

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"net/http"
)

func handleGetMembership(w http.ResponseWriter, r *http.Request) {
	userID, ok := shared.GetUserID(r.Context())
	if !ok {
		shared.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var res []db.Membership

	rows, err := db.DB.Query("SELECT * FROM membership WHERE user_id = $1", userID)
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
