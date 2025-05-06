package http

import (
	"encoding/json"
	"net/http"
	"todo-app-project/internal/usecase/auth"
)

type AuthHandler struct {
	authUseCase *auth.AuthUseCase // AuthUseCase arayüzü
}

func NewAuthHandler(authUseCase *auth.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase, // AuthUseCase'i alırız
	}
}

type LoginRequest struct {
	Username string `json:"username"` // Kullanıcı adı alanı
	Password string `json:"password"` // Şifre alanı
}

type LoginResponse struct {
	Token string `json:"token"` // Token alanı
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest) // hata varsa döneriz
		return
	}
	token, err := h.authUseCase.Login(req.Username, req.Password) // AuthUseCase'den token alırız
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized) // hata varsa döneriz
		return
	}
	res := LoginResponse{Token: token}                 // tokeni alırız
	w.Header().Set("Content-Type", "application/json") // içerik tipini ayarlarız
	json.NewEncoder(w).Encode(res)                     // yanıtı döneriz
}
