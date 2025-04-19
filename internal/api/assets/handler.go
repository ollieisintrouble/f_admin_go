package assets

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	handleGetAsset(w, r)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	handleCreateAsset(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDeleteAsset(w, r)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	handlePatchAsset(w, r)
}
