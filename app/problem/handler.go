package problem

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/gorilla/mux"
)

type handler struct {
	problemService interfaces.ProblemService
}

func NewHandler(problemService interfaces.ProblemService) *handler {
	return &handler{
		problemService: problemService,
	}
}

func (h *handler) GetProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	res, status, err := h.problemService.GetProblem(id)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = util.EncodeResponse(w, res, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
