package pool

import "github.com/FreeJ1nG/cpduel-backend/app/interfaces"

type serviceContainer struct {
	authService        interfaces.AuthService
	problemService     interfaces.ProblemService
	webscrapperService interfaces.WebscrapperService
}

func NewServiceContainer(
	authService interfaces.AuthService,
	problemService interfaces.ProblemService,
	webscrapperService interfaces.WebscrapperService,
) *serviceContainer {
	return &serviceContainer{
		authService:        authService,
		problemService:     problemService,
		webscrapperService: webscrapperService,
	}
}

func (sc *serviceContainer) GetAuthService() interfaces.AuthService {
	return sc.authService
}

func (sc *serviceContainer) GetProblemService() interfaces.ProblemService {
	return sc.problemService
}

func (sc *serviceContainer) GetWebscrapperService() interfaces.WebscrapperService {
	return sc.webscrapperService
}
