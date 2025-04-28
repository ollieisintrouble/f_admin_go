package shared

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"f_admin_go/internal/db"
)

func CheckOrg(r *http.Request, authenticator *SimpleAuthenticator) (string, string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", "", errors.New("unauthorized")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := authenticator.DecodeToken(token)
	if err != nil {
		return "", "", errors.New("unauthorized")
	}

	orgID := r.URL.Query().Get("org_id")
	if orgID == "" {
		return "", "", errors.New("no access allowed")
	}

	var membership db.Membership
	row := db.DB.QueryRow("SELECT * FROM membership WHERE user_id = $1 AND organization_id = $2", userID, orgID)
	if err := row.Scan(&membership.UserID, &membership.OrganizationID, &membership.Role); err != nil {
		if err == sql.ErrNoRows {
			return "", "", errors.New("no access allowed")
		}
		return "", "", errors.New("no access allowed")
	}

	return userID, orgID, nil
}
