package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
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

	var req models.RegisterForm
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := ConvertUserToDB(req)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Encryption failed")
	}
	role := req.Role

	tx, err := db.DB.Begin()
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Failed to start transaction")
		return
	}
	defer tx.Rollback()

	err = tx.QueryRow(`
		INSERT INTO users (username, password_hash, full_name, email, phone, image) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`, user.Username, user.PasswordHash, user.FullName, user.Email, user.Phone, user.Image).Scan(&user.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	_, err = tx.Exec(`
		INSERT INTO membership (user_id, organization_id, role) 
		VALUES ($1, $2, $3)
	`, user.ID, orgID, role)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Failed to assign membership")
		return
	}

	if err = tx.Commit(); err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Failed to commit transaction")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UserRegister(w http.ResponseWriter, r models.RegisterForm) (db.User, error) {
	result, err := ConvertUserToDB(r)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Encryption failed")
	}

	err = db.DB.QueryRow("INSERT INTO users (username, password_hash, full_name, email, phone, image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", result.Username, result.PasswordHash, result.FullName, result.Email, result.Phone, result.Image).Scan(&result.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return db.User{}, errors.New("database error")
	}

	return result, nil
}
