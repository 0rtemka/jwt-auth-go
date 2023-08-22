package handler

import (
	"net/http"
)

func (h *Handler) FindAllUsers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}
}
