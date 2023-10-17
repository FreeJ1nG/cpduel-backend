package strategies

import (
	"context"
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/chromedp/chromedp"
)

type submitProblemStrategy struct {
	ctx     context.Context
	client  *models.WebsocketClient
	payload dto.SubmitCodePayload
}

func NewSubmitProblemStrategy(ctx context.Context, client *models.WebsocketClient, payload dto.SubmitCodePayload) *submitProblemStrategy {
	return &submitProblemStrategy{
		ctx:     ctx,
		client:  client,
		payload: payload,
	}
}

func (s *submitProblemStrategy) Run(services interfaces.ServiceContainer) (err error) {
	submissionService := services.GetSubmissionService()
	webscrapperService := services.GetWebscrapperService()
	websocketUtil := services.GetWebsocketUtil()
	submission, err := submissionService.MakeSubmission(
		s.client.User.Username,
		s.payload.ProblemId,
		s.payload.SourceCode,
		s.payload.LanguageId,
	)
	if err != nil {
		return
	}

	ctx, cancel := chromedp.NewContext(s.ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	var completed bool = false
	go func() {
		err = webscrapperService.SubmitCode(ctx, &submission, s.payload.ProblemId, s.payload.SourceCode, s.payload.LanguageId)
		completed = true
	}()

	previousVerdict := submission.Verdict
	for {
		if previousVerdict != submission.Verdict {
			websocketUtil.SendJSONToClient(s.client, dto.WebsocketMessageResponse{
				Type: dto.SubmissionStateMessageType,
				Data: &dto.ResponseData{
					Submission: submission,
				},
			})
		}
		previousVerdict = submission.Verdict
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			websocketUtil.SendErrorToClient(s.client, err.Error())
			return
		}
		if completed {
			return
		}
	}
}
