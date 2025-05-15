package users

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleGetUser(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.DB.Query("SELECT id, username, full_name, email, phone, image, created_at FROM users WHERE id IN (SELECT user_id FROM membership WHERE organization_id = $1)", orgID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	var responses []models.UserDTO
	for rows.Next() {
		var user db.User
		if err := rows.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.Phone, &user.Image, &user.CreatedAt); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		var role string
		err = db.DB.QueryRow("SELECT role FROM membership WHERE user_id = $1 AND organization_id = $2", user.ID, orgID).Scan(&role)
		if err != nil {
			role = "Undefined"
		}
		responses = append(responses, ConvertUserFromDB(user, role))
	}

	if len(responses) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No transactions found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, responses)
}
