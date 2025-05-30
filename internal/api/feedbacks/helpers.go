package feedbacks

import (
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
)

func ConvertFeedbackFromDB(f db.Feedback) models.FeedbackDTO {
	return models.FeedbackDTO{
		Content:   f.Content,
		Email:     shared.NullStringPtr(f.Email),
		Product:   shared.NullStringPtr(f.Product),
		CreatedAt: f.CreatedAt,
	}
}

func ConvertFeedbackToDB(f models.FeedbackDTO) db.Feedback {
	return db.Feedback{
		Content: f.Content,
		Email:   shared.StringToNullString(f.Email),
		Product: shared.StringToNullString(f.Product),
	}
}
