package feedbacks

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var req models.FeedbackDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	feedback := ConvertFeedbackToDB(req)
	var res int64

	err := db.DB.QueryRow("INSERT INTO feedbacks (content, email, product, created_at) VALUES ($1, $2, $3) RETURNING id", &feedback.Content, &feedback.Email, &feedback.Product).Scan(&res)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database insert error")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, res)
}
