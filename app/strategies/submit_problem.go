package strategies

import (
	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type submitProblemStrategy struct {
	Client  *models.WebsocketClient
	Payload dto.SubmitCodePayload
}

func NewSubmitProblemStrategy(client *models.WebsocketClient, payload dto.SubmitCodePayload) *submitProblemStrategy {
	return &submitProblemStrategy{
		Client:  client,
		Payload: payload,
	}
}

func (strategy *submitProblemStrategy) Run(services interfaces.ServiceContainer) (err error) {
	// TODO: FINISH THIS METHOD
	// problemService := services.GetProblemService()
	// webscrapperService := services.GetWebscrapperService()
	// problemNumber, problemCode := util.ParseProblemId(strategy.Payload.ProblemId)
	return
}
