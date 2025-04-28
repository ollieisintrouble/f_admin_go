package auth

import (
	"database/sql"
	"errors"
	"f_admin_go/internal/db"
)

type UserQuery struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func FindUserByUsername(u string) (UserQuery, error) {
	var user UserQuery
	row := db.DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", u)
	if err := row.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return UserQuery{}, errors.New("user not found")
		}
		return UserQuery{}, errors.New("database forbidden")
	}
	return user, nil
}
