package ws

import (
	"encoding/json"
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/app/strategies"
)

type service struct {
	pool interfaces.Pool
}

func NewService(pool interfaces.Pool) *service {
	return &service{
		pool: pool,
	}
}

func (s *service) ReadMessageFromClient(client *models.WebsocketClient) (err error) {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			err = fmt.Errorf("unable to read message from client: %s", err.Error())
			return err
		}

		var clientMessage dto.ClientWebsocketMessage
		if err = json.Unmarshal(p, &clientMessage); err != nil {
			err = fmt.Errorf("unable to parse client message: %s", err.Error())
			return err
		}

		payload, err := clientMessage.ParsePayload()
		if err != nil {
			err = fmt.Errorf("unable to parse payload: %s", err.Error())
			return err
		}

		if clientMessage.Type == "submit" {
			algo := strategies.NewSubmitProblemStrategy(client, payload.(dto.SubmitCodePayload))
			s.pool.SetAlgo(algo)
		}
	}
}
