package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/golang-jwt/jwt/v4"
)

type ContextKey string

var UserContextKey = ContextKey("user")

type routeProtector struct {
	authUtil interfaces.AuthUtil
	authRepo interfaces.AuthRepository
}

func NewRouteProtector(authUtil interfaces.AuthUtil, authRepo interfaces.AuthRepository) *routeProtector {
	return &routeProtector{
		authUtil: authUtil,
		authRepo: authRepo,
	}
}

func (rp *routeProtector) Wrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := rp.authUtil.ExtractJwtToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := rp.authUtil.ConvertJwtStringToToken(tokenString)
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
