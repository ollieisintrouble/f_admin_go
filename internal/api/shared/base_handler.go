package shared

import (
	"context"
	"net/http"
)

type EntityHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
}

func HandleEntity(h EntityHandler, authenticator *SimpleAuthenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		userID, orgID, err := CheckOrg(r, authenticator)
		if err != nil {
			WriteError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserIDContextKey, userID)
		ctx = context.WithValue(ctx, OrgIDContextKey, orgID)
		r = r.WithContext(ctx)

		switch r.Method {
		case http.MethodGet:
			h.Get(w, r)
		case http.MethodPost:
			h.Post(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		case http.MethodPatch:
			h.Patch(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
