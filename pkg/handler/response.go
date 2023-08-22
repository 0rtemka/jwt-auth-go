package handler

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Path    string `json:"path"`
	Status  int    `json:"status"`
}

func NewErrorResponse(msg, path string, status int, w http.ResponseWriter) {
	errorResponse := Error{
		Message: msg,
		Path:    path,
		Status:  status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorResponse.Status)
	json.NewEncoder(w).Encode(errorResponse)
}
