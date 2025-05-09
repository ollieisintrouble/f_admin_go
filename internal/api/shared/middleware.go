package shared

import (
	"database/sql"
	"errors"
	"net/http"

	"f_admin_go/internal/db"
)

func CheckOrg(r *http.Request, authenticator *SimpleAuthenticator) (string, string, error) {
	cookie, err := r.Cookie("authToken")
	if err != nil {
		return "", "", errors.New("no auth token found")
	}
	token := cookie.Value
	userID, err := authenticator.DecodeToken(token)
	if err != nil {
		return "", "", errors.New(err.Error())
	}

	orgIDStr := r.URL.Query().Get("org_id")
	if orgIDStr == "" {
		return "", "", errors.New("no access allowed")
	}
	orgID, err := ConvertOrgIDToInt(orgIDStr)
	if err != nil {
		return "", "", errors.New("invalid org ID")
	}

	var membership db.Membership
	row := db.DB.QueryRow("SELECT * FROM membership WHERE user_id = $1 AND organization_id = $2", userID, orgID)
	if err := row.Scan(&membership.UserID, &membership.OrganizationID, &membership.Role); err != nil {
		if err == sql.ErrNoRows {
			return "", "", errors.New("no access allowed")
		}
		return "", "", errors.New("no access allowed")
	}

	return userID, orgIDStr, nil
}
