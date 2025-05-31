package auth

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/api/users"
	"f_admin_go/internal/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request, authenticator *shared.SimpleAuthenticator) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req models.RegisterForm

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := users.UserRegister(w, req)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Unable to create account")
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
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	})

	w.WriteHeader(http.StatusOK)
}
