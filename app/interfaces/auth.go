package interfaces

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	RegisterUser(username string, fullName string, password string) (res dto.RegisterResponse, status int, err error)
	AuthenticateUser(username string, password string) (res dto.LoginResponse, status int, err error)
	GetUserByUsername(username string) (user models.User, status int, err error)
}

type AuthRepository interface {
	CreateUser(username string, fullName string, passwordHash string) (user models.User, err error)
	GetUserByUsername(username string) (user models.User, err error)
}

type AuthHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	AuthenticateUser(w http.ResponseWriter, r *http.Request)
	GetCurrentUser(w http.ResponseWriter, r *http.Request)
}

type AuthUtil interface {
	GenerateToken(user models.User) (signedToken string, err error)
	HashPassword(password string) (passwordHash string, err error)
	ExtractJwtToken(r *http.Request) (jwtToken string, err error)
	ConvertJwtStringToToken(tokenString string) (token *jwt.Token, err error)
}
