package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) FindAllUsers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.services.User.FindAll()
		if err != nil {
			NewErrorResponse("error finding users: "+err.Error(), r.RequestURI, http.StatusInternalServerError, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
