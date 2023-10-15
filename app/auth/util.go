package auth

import (
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user models.User) (signedToken string, err error) {
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "cpduel-api",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			Subject:   user.Username,
		},
	)
	signedToken, err = token.SignedString([]byte(viper.GetString("JWT_SECRET_KEY")))
	return
}

func HashPassword(password string) (passwordHash string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	passwordHash = string(hashedPassword)
	return
}
