package handler

import (
	"net/http"
)

func (h *Handler) GenerateTokens() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handler) RefreshTokens() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}
}
