package http

import (
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/domain/DTO"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserService user.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req DTO.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.CreateUser(req.Name, req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
