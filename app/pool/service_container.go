package pool

import "github.com/FreeJ1nG/cpduel-backend/app/interfaces"

type serviceContainer struct {
	authUtil           interfaces.AuthUtil
	websocketUtil      interfaces.WebsocketUtil
	authService        interfaces.AuthService
	problemService     interfaces.ProblemService
	submissionService  interfaces.SubmissionService
	webscrapperService interfaces.WebscrapperService
}

func NewServiceContainer(
	authUtil interfaces.AuthUtil,
	websocketUtil interfaces.WebsocketUtil,
	authService interfaces.AuthService,
	problemService interfaces.ProblemService,
	submissionService interfaces.SubmissionService,
	webscrapperService interfaces.WebscrapperService,
) *serviceContainer {
	return &serviceContainer{
		authUtil:           authUtil,
		websocketUtil:      websocketUtil,
		authService:        authService,
		problemService:     problemService,
		submissionService:  submissionService,
		webscrapperService: webscrapperService,
	}
}

func (sc *serviceContainer) GetAuthUtil() interfaces.AuthUtil {
	return sc.authUtil
}

func (sc *serviceContainer) GetWebsocketUtil() interfaces.WebsocketUtil {
	return sc.websocketUtil
}

func (sc *serviceContainer) GetAuthService() interfaces.AuthService {
	return sc.authService
}

func (sc *serviceContainer) GetProblemService() interfaces.ProblemService {
	return sc.problemService
}

func (sc *serviceContainer) GetSubmissionService() interfaces.SubmissionService {
	return sc.submissionService
}

func (sc *serviceContainer) GetWebscrapperService() interfaces.WebscrapperService {
	return sc.webscrapperService
}
