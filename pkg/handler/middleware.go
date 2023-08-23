package handler

import (
	"net/http"
	"strings"
)

const (
	AuthHeader = "Authorization"
)

func (h *Handler) VerifyToken(nextHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(AuthHeader)
		if header == "" {
			NewErrorResponse("Authorization header not specified", r.RequestURI, http.StatusBadRequest, w)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || len(headerParts) == 2 && headerParts[0] != "Bearer" {
			NewErrorResponse("invalid Authorization header", r.RequestURI, http.StatusBadRequest, w)
			return
		}

		_, err := h.services.Auth.ParseToken(headerParts[1])
		if err != nil {
			NewErrorResponse("invalid access token", r.RequestURI, http.StatusBadRequest, w)
			return
		}

		nextHandler(w, r)
	}
}
