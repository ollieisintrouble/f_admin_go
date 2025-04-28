package users

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func ConvertUserToDB(r models.RegisterForm) db.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return db.User{
		Username:     r.Username,
		PasswordHash: string(hash),
		FullName:     shared.StringToNullString(*r.FullName),
		Email:        shared.StringToNullString(*r.Email),
		Phone:        shared.StringToNullString(*r.Phone),
	}
}
