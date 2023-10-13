package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

func (s *Server) InjectDependencies(ctx context.Context) {
	s.router.Use(util.LoggerMiddleware)

	// Repositories
	problemRepo := problem.NewRepository(s.db)

	// Services
	webscrapperService := webscrapper.NewService(ctx)
	problemService := problem.NewService(ctx, problemRepo, webscrapperService)

	// Controllers
	problemHandler := problem.NewHandler(problemService)

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
	}).Methods("GET")

	problemRouter := s.router.PathPrefix("/problems").Subrouter()
	problemRouter.HandleFunc("/{id}", problemHandler.GetProblem).Methods("GET")
}
