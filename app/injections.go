package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/auth"
	"github.com/FreeJ1nG/cpduel-backend/app/pool"
	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/app/ws"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

func (s *Server) InjectDependencies(ctx context.Context) {
	s.router.Use(util.LoggerMiddleware)

	// Repositories
	authRepo := auth.NewRepository(s.db)
	problemRepo := problem.NewRepository(s.db)

	routeProtector := util.NewRouteProtector(authRepo)

	// Services
	authService := auth.NewService(authRepo)
	webscrapperService := webscrapper.NewService(ctx)
	problemService := problem.NewService(ctx, problemRepo, webscrapperService)

	serviceContainer := pool.NewServiceContainer(authService, problemService, webscrapperService)
	pool := pool.NewPool(serviceContainer)
	websocketService := ws.NewService(pool)

	// Controllers
	problemHandler := problem.NewHandler(problemService)
	websocketHandler := ws.NewHandler(websocketService)
	authHandler := auth.NewHandler(authService)

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
	}).Methods("GET")

	s.router.HandleFunc("/websocket", websocketHandler.WebsocketConnectionHandler)

	authRouter := s.router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.AuthenticateUser).Methods("POST")
	authRouter.HandleFunc("/register", authHandler.RegisterUser).Methods("POST")
	authRouter.HandleFunc("/me", routeProtector.Wrapper(authHandler.GetCurrentUser)).Methods("GET")

	problemRouter := s.router.PathPrefix("/problems").Subrouter()
	problemRouter.HandleFunc("/{id}", problemHandler.GetProblem).Methods("GET")

}
