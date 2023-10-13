package problem

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	problemService Service
}

type Handler interface {
	GetProblem(w http.ResponseWriter, r *http.Request)
}

func NewHandler(problemService Service) *handler {
	return &handler{
		problemService: problemService,
	}
}

func (h *handler) GetProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, status, err := h.problemService.GetProblem(id)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Printf("unable to encode problem: %s\n", err.Error())
	}
}
