package feedbacks

import (
	"database/sql"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	var res []models.FeedbackDTO
	productQuery := r.URL.Query().Get("product")
	var rows *sql.Rows
	var err error

	if productQuery == "" {
		rows, err = db.DB.Query("SELECT content, email, product, created_at FROM feedbacks ORDER BY created_at DESC")
	} else {
		rows, err = db.DB.Query("SELECT content, email, product, created_at FROM feedbacks WHERE product = $1 ORDER BY created_at DESC", productQuery)
	}
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var feedback db.Feedback
		if err := rows.Scan(&feedback.Content, &feedback.Email, &feedback.Product, &feedback.CreatedAt); err != nil {
			shared.WriteError(w, http.StatusInternalServerError, "Database scan error")
			return
		}
		res = append(res, ConvertFeedbackFromDB(feedback))
	}

	if len(res) == 0 {
		shared.WriteError(w, http.StatusNotFound, "No feedback found")
		return
	}

	shared.WriteJSON(w, http.StatusOK, res)
}
