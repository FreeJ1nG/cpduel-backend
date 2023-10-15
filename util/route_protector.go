package util

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type ContextKey string

var UserContextKey = ContextKey("user")

type routeProtector struct {
	authRepo interfaces.AuthRepository
}

func NewRouteProtector(authRepo interfaces.AuthRepository) *routeProtector {
	return &routeProtector{
		authRepo: authRepo,
	}
}

func extractJwtToken(r *http.Request) (jwtToken string, err error) {
	authorization := r.Header.Get("Authorization")
	authSplit := strings.Split(authorization, " ")
	if len(authSplit) != 2 {
		err = fmt.Errorf("invalid authorization header format")
		return
	}
	prefix := authSplit[0]
	tokenString := authSplit[1]
	if prefix != "Bearer" {
		err = fmt.Errorf("JWT token not found on authorization header")
		return
	}
	jwtToken = tokenString
	return
}

func convertJwtStringToToken(tokenString string) (token *jwt.Token, err error) {
	if err != nil {
		err = fmt.Errorf("unable to extract jwt token: %s", err.Error())
		return
	}

	token, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Method)
		}
		return []byte(viper.GetString("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		err = fmt.Errorf("unable to parse token: %s", err.Error())
		return
	}
	return
}

func (rp *routeProtector) Wrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractJwtToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := convertJwtStringToToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "unable to get token claims", http.StatusInternalServerError)
			return
		}

		username := claims["sub"].(string)
		user, err := rp.authRepo.GetUserByUsername(username)
		if err != nil {
			http.Error(w, fmt.Sprintf("user with username of %s not found", username), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)
		f(w, r.WithContext(ctx))
	}
}
