package handler

import (
	"encoding/json"
	"github.com/secret-santa-rubot/api/internal/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Структура запроса
type registerUserRequest struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Структура ответа
type registerUserResponse struct {
	Status string `json:"status"`
	UserID string `json:"userId"`
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest

	// Читаем JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем сервис
	userID, err := h.userService.RegisterUser(r.Context(), req.UserID, req.Name, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := registerUserResponse{Status: "ok", UserID: userID}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
