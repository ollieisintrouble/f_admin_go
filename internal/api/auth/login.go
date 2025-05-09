package auth

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/models"
)

func Login(w http.ResponseWriter, r *http.Request, authenticator *shared.SimpleAuthenticator) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := FindUserByUsername(creds.Username)
	if err != nil {
		shared.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
		shared.WriteError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := authenticator.GenerateToken(user.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}
