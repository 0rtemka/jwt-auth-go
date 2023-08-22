package handler

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"test/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
