package ws

import (
	"fmt"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
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
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	client := models.NewWebsocketClient(conn, nil)

	err = h.websocketService.ReadMessageFromClient(client)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
