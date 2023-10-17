package interfaces

type ServiceContainer interface {
	GetAuthUtil() AuthUtil
	GetWebsocketUtil() WebsocketUtil
	GetAuthService() AuthService
	GetProblemService() ProblemService
	GetSubmissionService() SubmissionService
	GetWebscrapperService() WebscrapperService
}
