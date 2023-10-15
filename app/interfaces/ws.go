package interfaces

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type WebsocketService interface {
	ReadMessageFromClient(client *models.WebsocketClient) (err error)
}

type WebsocketRepository interface {
}

type WebsocketHandler interface {
	WebsocketConnectionHandler(w http.ResponseWriter, r *http.Request)
}
