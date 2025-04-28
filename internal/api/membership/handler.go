package membership

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	handleGetMembership(w, r)
}

// func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
// 	handleCreateMembership(w, r)
// }

// func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
// 	handleDeleteMembership(w, r)
// }

// func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
// 	handlePatchMembership(w, r)
// }
