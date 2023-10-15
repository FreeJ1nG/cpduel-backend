package interfaces

type ServiceContainer interface {
	GetAuthService() AuthService
	GetProblemService() ProblemService
	GetWebscrapperService() WebscrapperService
}

type PoolAlgo interface {
	Run(services ServiceContainer) (err error)
}

type Pool interface {
	SetAlgo(algo PoolAlgo)
	Start() (err error)
}
