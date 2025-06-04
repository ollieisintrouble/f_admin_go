package auth

import (
	"f_admin_go/internal/api/organizations"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/api/users"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
	"strings"
)

func ValidateToken(w http.ResponseWriter, r *http.Request, authenticator *shared.SimpleAuthenticator) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		shared.WriteError(w, http.StatusBadRequest, "Invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := authenticator.DecodeToken(token)
	if err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid token")
		return
	}

	var user db.User
	var res models.TokenValidationResponse

	row := db.DB.QueryRow("SELECT username, full_name, email, phone, image, created_at FROM users WHERE id = $1", userID)
	if err := row.Scan(&user.Username, &user.FullName, &user.Email, &user.Phone, &user.Image, &user.CreatedAt); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid token")
		return
	}
	res.User = users.ConvertUserFromDB(user, "")

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
		res.Memberships = append(res.Memberships, membership)
	}
	if len(res.Memberships) == 0 {
		shared.WriteError(w, http.StatusNotAcceptable, "No membership found")
		return
	}

	for _, membership := range res.Memberships {
		var org db.Organization
		row := db.DB.QueryRow("SELECT * FROM organizations WHERE id = $1 AND purchase_package IS NOT NULL AND purchase_package IS DISTINCT FROM 'None'", membership.OrganizationID)
		if err := row.Scan(&org.ID, &org.Name, &org.Image, &org.CreatedAt, &org.PurchasePackage); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		res.Organizations = append(res.Organizations, organizations.ConvertOrganizationFromDB(org))
	}
	if len(res.Organizations) == 0 {
		shared.WriteError(w, http.StatusPaymentRequired, "No paid organization found")
		return
	}
	shared.WriteJSON(w, http.StatusOK, res)
}
