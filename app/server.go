package app

import (
	"log"

	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	config util.Config
	router *fiber.App
	db     *pgxpool.Pool
}

func MakeServer(config util.Config, db *pgxpool.Pool) Server {
	r := fiber.New()
	return Server{
		config: config,
		router: r,
		db:     db,
	}
}

func (s *Server) RunServer() {
	port := s.config.ServerPort
	if err := s.router.Listen(":" + port); err != nil {
		log.Fatal("Unable to start server", err.Error())
	}
}
