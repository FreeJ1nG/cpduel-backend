package models

import (
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	Conn *websocket.Conn
	User *User
}

func NewWebsocketClient(conn *websocket.Conn, user *User) *WebsocketClient {
	return &WebsocketClient{
		Conn: conn,
		User: user,
	}
}
