package strategies

import (
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/golang-jwt/jwt/v4"
)

type handshakeStrategy struct {
	Client  *models.WebsocketClient
	Payload dto.HandshakePayload
}

func NewHandshakeStrategy(client *models.WebsocketClient, payload dto.HandshakePayload) *handshakeStrategy {
	return &handshakeStrategy{
		Client:  client,
		Payload: payload,
	}
}

func (s *handshakeStrategy) Run(services interfaces.ServiceContainer) (err error) {
	authUtil := services.GetAuthUtil()
	authService := services.GetAuthService()

	token, err := authUtil.ConvertJwtStringToToken(s.Payload.JwtToken)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("unable to get token claims")
		return
	}

	username := claims["sub"].(string)
	user, _, err := authService.GetUserByUsername(username)
	if err != nil {
		return
	}

	s.Client.User = &user
	return
}
