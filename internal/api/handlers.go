package api

import (
	"net/http"
)

func HandleAssetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case http.MethodGet:
		getAsset(w, r)
	case http.MethodPost:
		createAsset(w, r)
	// case http.MethodPut:
	// 	updateAsset(w, r)
	case http.MethodDelete:
		deleteAsset(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
