package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/auth"
	"github.com/FreeJ1nG/cpduel-backend/app/pool"
	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/submission"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/app/ws"
	"github.com/FreeJ1nG/cpduel-backend/util"
)

func (s *Server) InjectDependencies(ctx context.Context) {
	s.router.Use(util.LoggerMiddleware)

	// Utils
	authUtil := auth.NewUtil()
	websocketUtil := ws.NewUtil()

	// Repositories
	authRepo := auth.NewRepository(s.db)
	problemRepo := problem.NewRepository(s.db)
	submissionRepo := submission.NewRepository(s.db)

	routeProtector := util.NewRouteProtector(authUtil, authRepo)

	// Services
	authService := auth.NewService(authUtil, authRepo)
	webscrapperService := webscrapper.NewService(ctx)
	problemService := problem.NewService(ctx, problemRepo, webscrapperService)
	submissionService := submission.NewService(submissionRepo, webscrapperService)

	serviceContainer := pool.NewServiceContainer(authUtil, websocketUtil, authService, problemService, submissionService, webscrapperService)
	pool := pool.NewPool(serviceContainer)
	go pool.Start()

	websocketService := ws.NewService(ctx, pool)

	// Controllers
	problemHandler := problem.NewHandler(problemService)
	websocketHandler := ws.NewHandler(websocketService)
	authHandler := auth.NewHandler(authService)
	submissionHandler := submission.NewHandler(submissionService)

	s.router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
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

	submissionRouter := s.router.PathPrefix("/submission").Subrouter()
	submissionRouter.HandleFunc("/my", routeProtector.Wrapper(submissionHandler.GetMySubmissions)).Methods("GET")
}
