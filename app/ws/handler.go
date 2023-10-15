package ws

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

type handler struct {
	websocketService interfaces.WebsocketService
}

func NewHandler(websocketService interfaces.WebsocketService) *handler {
	return &handler{
		websocketService: websocketService,
	}
}

func (h *handler) WebsocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeConnection(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := r.Context().Value(util.UserContextKey).(models.User)
	client := models.NewWebsocketClient(conn, &user)

	h.websocketService.ReadMessageFromClient(client)
}
