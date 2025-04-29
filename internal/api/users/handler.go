package users

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	handleGetUser(w, r)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	handleCreateUser(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDeleteUser(w, r)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	handlePatchUser(w, r)
}
