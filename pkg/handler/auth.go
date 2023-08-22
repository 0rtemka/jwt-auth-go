package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	cookieRefreshTokenName = "refresh_token"
	cookieMaxAge           = 3600 * 24 * 14 // 2 weeks
	cookieDomain           = "localhost"
)

func (h *Handler) GenerateTokens() http.HandlerFunc {

	type Response struct {
		AccessToken string `json:"access_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]

		accessToken, refreshToken, err := h.services.GenerateNewPair(userId)
		if err != nil {
			NewErrorResponse("error generating tokens: "+err.Error(), r.RequestURI, http.StatusInternalServerError, w)
			return
		}

		jsonResp := Response{AccessToken: accessToken}

		http.SetCookie(w, &http.Cookie{
			Name:   cookieRefreshTokenName,
			Value:  refreshToken,
			MaxAge: cookieMaxAge,
			Path:   r.RequestURI,
			Domain: cookieDomain,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResp)
	}
}

func (h *Handler) RefreshTokens() http.HandlerFunc {

	type Response struct {
		AccessToken string `json:"access_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		rt, err := r.Cookie(cookieRefreshTokenName)
		if err != nil {
			NewErrorResponse("cookie 'refresh_token' not found", r.RequestURI, http.StatusBadRequest, w)
			return
		}

		userId := mux.Vars(r)["id"]

		accessToken, refreshToken, err := h.services.Refresh(userId, rt.Value)
		if err != nil {
			NewErrorResponse("error refreshing tokens: "+err.Error(), r.RequestURI, http.StatusInternalServerError, w)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   cookieRefreshTokenName,
			Value:  refreshToken,
			MaxAge: cookieMaxAge,
			Path:   r.RequestURI,
			Domain: cookieDomain,
		})

		jsonResp := Response{AccessToken: accessToken}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResp)
	}
}
