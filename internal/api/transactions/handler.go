package transactions

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	handleGetTransaction(w, r)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	handleCreateTransaction(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDeleteTransaction(w, r)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	handlePatchTransaction(w, r)
}
