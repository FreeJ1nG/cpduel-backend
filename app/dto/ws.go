package dto

import (
	"encoding/json"
	"fmt"
)

type ClientWebsocketMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type WebsocketMessageResponse struct {
}

func (wsm *ClientWebsocketMessage) ParsePayload() (payload interface{}, err error) {
	if err = json.Unmarshal([]byte(wsm.Payload), &payload); err != nil {
		err = fmt.Errorf("unable to parse client message: %s", err.Error())
		return
	}
	return
}
