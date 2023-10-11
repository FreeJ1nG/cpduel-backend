package app

import "github.com/FreeJ1nG/cpduel-backend/util"

func (s *Server) InjectDependencies() {
	s.router.Use(util.LoggerMiddleware())

}
