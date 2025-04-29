package users

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ConvertUserToDB(r models.RegisterForm) (db.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, errors.New("unauthorized")
	}

	return db.User{
		Username:     r.Username,
		PasswordHash: string(hash),
		FullName:     shared.StringToNullString(*r.FullName),
		Email:        shared.StringToNullString(*r.Email),
		Phone:        shared.StringToNullString(*r.Phone),
		Image:        shared.StringToNullString(*r.Image),
	}, nil
}

func UpdateUserDB(u models.UserDTO) db.User {
	return db.User{
		Username: u.Username,
		FullName: shared.StringToNullString(*u.FullName),
		Email:    shared.StringToNullString(*u.Email),
		Phone:    shared.StringToNullString(*u.Phone),
		Image:    shared.StringToNullString(*u.Image),
	}
}

func ConvertUserFromDB(u db.User) models.UserDTO {
	return models.UserDTO{
		Username:  u.Username,
		FullName:  shared.NullStringPtr(u.FullName),
		Email:     shared.NullStringPtr(u.Email),
		Phone:     shared.NullStringPtr(u.Phone),
		Image:     shared.NullStringPtr(u.Image),
		CreatedAt: u.CreatedAt,
	}
}
