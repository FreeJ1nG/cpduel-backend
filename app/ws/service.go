package ws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/app/strategies"
)

type service struct {
	ctx  context.Context
	pool interfaces.Pool
}

func NewService(ctx context.Context, pool interfaces.Pool) *service {
	return &service{
		ctx:  ctx,
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
		fmt.Println("Websocket message:\n", string(p))

		var clientMessage dto.ClientWebsocketMessage
		if err = json.Unmarshal(p, &clientMessage); err != nil {
			err = fmt.Errorf("unable to parse client message: %s", err.Error())
			return err
		}

		payload := clientMessage.Payload

		var algo interfaces.PoolAlgo
		if clientMessage.Type == dto.HandshakeMessageType {
			algo = strategies.NewHandshakeStrategy(
				client,
				dto.HandshakePayload{JwtToken: payload.JwtToken},
			)
		} else if clientMessage.Type == dto.SubmitMessageType {
			algo = strategies.NewSubmitProblemStrategy(
				s.ctx,
				client,
				dto.SubmitCodePayload{
					ProblemId:  payload.ProblemId,
					LanguageId: payload.LanguageId,
					SourceCode: payload.SourceCode,
				},
			)
		}

		s.pool.SetAlgo(algo)
	}
}
