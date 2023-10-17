package interfaces

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
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

type WebsocketUtil interface {
	SendJSONToClient(client *models.WebsocketClient, data dto.WebsocketMessageResponse) (err error)
	SendErrorToClient(client *models.WebsocketClient, errorMessage string) (err error)
}
