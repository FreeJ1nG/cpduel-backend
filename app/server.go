package app

import (
	"log"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

type Server struct {
	config     util.Config
	router     *mux.Router
	db         *pgxpool.Pool
	httpServer *http.Server
}

func MakeServer(config util.Config, db *pgxpool.Pool) Server {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	return Server{
		config: config,
		router: r,
		db:     db,
		httpServer: &http.Server{
			Addr:    ":" + config.ServerPort,
			Handler: cors.AllowAll().Handler(r),
		},
	}
}

func (s *Server) RunServer() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal("unable to start server: ", err.Error())
	}
}
