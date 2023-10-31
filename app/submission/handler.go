package submission

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

type handler struct{
	submissionService interfaces.SubmissionService
}

func NewHandler(submissionService interfaces.SubmissionService) *handler {
	return &handler{
		submissionService: submissionService,
	}
}

func (h *handler) GetMySubmissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(util.UserContextKey).(models.User)

	submissions, status, err := h.submissionService.GetSubmissionsOfUser(user.Username)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = util.EncodeResponse(
		w, 
		dto.GetSubmissionsOfUserResponse{
			Submissions: submissions,
		}, 
		status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
