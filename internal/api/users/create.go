package users

import (
	"errors"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

// User factory defined temporarily for register, will refine later
func CreateUser(w http.ResponseWriter, r models.RegisterForm) (db.User, error) {
	result := ConvertUserToDB(r)

	err := db.DB.QueryRow("INSERT INTO users (username, password_hash, full_name, email, phone, image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", result.Username, result.PasswordHash, result.FullName, result.Email, result.Phone, result.Image).Scan(&result.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return db.User{}, errors.New("database error")
	}

	return result, nil
}
