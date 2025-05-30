package feedbacks

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case http.MethodPost:
		CreateFeedback(w, r)
	case http.MethodGet:
		GetFeedbacks(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
