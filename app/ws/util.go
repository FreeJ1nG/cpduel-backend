package ws

import (
	"fmt"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/gorilla/websocket"
)

type util struct {
}

func NewUtil() *util {
	return &util{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgradeConnection(w http.ResponseWriter, r *http.Request) (conn *websocket.Conn, err error) {
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = fmt.Errorf("unable to upgrade connection: %s", err.Error())
		return
	}
	return
}

func (u *util) SendJSONToClient(client *models.WebsocketClient, data dto.WebsocketMessageResponse) (err error) {
	err = client.Conn.WriteJSON(data)
	if err != nil {
		err = fmt.Errorf("unable to send json message to client: %s", err.Error())
		return
	}
	return
}

func (u *util) SendErrorToClient(client *models.WebsocketClient, errorMessage string) (err error) {
	data := dto.WebsocketMessageResponse{
		Type:  dto.ErrorMessageType,
		Error: errorMessage,
	}
	err = client.Conn.WriteJSON(data)
	if err != nil {
		err = fmt.Errorf("unable to send error to client: %s", err.Error())
		return
	}
	return
}
