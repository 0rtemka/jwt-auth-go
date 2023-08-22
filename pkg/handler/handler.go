package handler

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", h.FindAllUsers()).Methods("GET")

	router.HandleFunc("/auth/sign-in", h.GenerateTokens()).Methods("POST")
	router.HandleFunc("/auth/refresh", h.RefreshTokens()).Methods("POST")

	router.HandleFunc("/", h.HomePage()).Methods("GET")

	return router
}

func (h *Handler) HomePage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Secured Route")
	}
}
