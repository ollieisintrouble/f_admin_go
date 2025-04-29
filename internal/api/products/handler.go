package products

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	handleGetProduct(w, r)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	handleCreateProduct(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDeleteProduct(w, r)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	handlePatchProduct(w, r)
}
