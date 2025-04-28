package organizations

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func HandleCreateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var req models.CreateOrginzationForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	org := ConvertOrganizationToDB(req.Organization)
	var res int64

	err := db.DB.QueryRow("INSERT INTO organizations (name, image) VALUES ($1, $2) RETURNING id", org.Name, org.Image).Scan(&res)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	err = db.DB.QueryRow("INSERT INTO membership (user_id, organization_id, role) VALUES ($1, $2, $3)", req.UserID, res, "admin").Scan(&res)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Error inserting membership")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, res)
}
