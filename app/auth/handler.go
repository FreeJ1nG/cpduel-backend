package auth

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

type handler struct {
	authService interfaces.AuthService
}

func NewHandler(authService interfaces.AuthService) *handler {
	return &handler{
		authService: authService,
	}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := util.ParseRequestBody[dto.RegisterRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, status, err := h.authService.RegisterUser(body.Username, body.FullName, body.Password)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = util.EncodeResponse(w, res, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := util.ParseRequestBody[dto.LoginRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, status, err := h.authService.AuthenticateUser(body.Username, body.Password)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = util.EncodeResponse(w, res, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(util.UserContextKey).(models.User)

	err := util.EncodeResponse(
		w,
		dto.GetCurrentUserResponse{
			Username: user.Username,
			FullName: user.FullName,
		},
		http.StatusOK,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
