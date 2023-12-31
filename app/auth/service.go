package auth

import (
	"fmt"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type service struct {
	authUtil interfaces.AuthUtil
	authRepo interfaces.AuthRepository
}

func NewService(authUtil interfaces.AuthUtil, authRepo interfaces.AuthRepository) *service {
	return &service{
		authUtil: authUtil,
		authRepo: authRepo,
	}
}

func (s *service) RegisterUser(username string, fullName string, password string) (res dto.RegisterResponse, status int, err error) {
	status = http.StatusOK

	passwordHash, err := s.authUtil.HashPassword(password)
	if err != nil {
		err = fmt.Errorf("failed to hash password: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}

	user, err := s.authRepo.CreateUser(username, fullName, passwordHash)
	if err != nil {
		err = fmt.Errorf("unable to create user: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}

	signedToken, err := s.authUtil.GenerateToken(user)
	if err != nil {
		err = fmt.Errorf("unable to generate token: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}

	res = dto.RegisterResponse{
		Token: signedToken,
	}
	return
}

func (s *service) AuthenticateUser(username string, password string) (res dto.LoginResponse, status int, err error) {
	status = http.StatusOK

	user, err := s.authRepo.GetUserByUsername(username)
	if err != nil {
		err = fmt.Errorf("unable to authenticate user: user with username of %s not found", username)
		status = http.StatusNotFound
		return
	}

	if !user.ValidatePasswordHash(password) {
		err = fmt.Errorf("unable to authenticate user: invalid password")
		status = http.StatusUnauthorized
		return
	}

	signedToken, err := s.authUtil.GenerateToken(user)
	if err != nil {
		err = fmt.Errorf("unable to generate token: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}

	res = dto.LoginResponse{
		Token: signedToken,
	}
	return
}

func (s *service) GetUserByUsername(username string) (user models.User, status int, err error) {
	status = http.StatusOK
	user, err = s.authRepo.GetUserByUsername(username)
	if err != nil {
		err = fmt.Errorf("unable to find user: %s", err.Error())
		status = http.StatusNotFound
		return
	}
	return
}
