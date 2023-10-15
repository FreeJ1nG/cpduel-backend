package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func upgradeConnection(w http.ResponseWriter, r *http.Request) (conn *websocket.Conn, err error) {
	conn, err = Upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = fmt.Errorf("unable to upgrade connection: %s", err.Error())
		return
	}
	return
}
