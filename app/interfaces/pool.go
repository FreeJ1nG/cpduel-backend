package interfaces

type PoolAlgo interface {
	Run(services ServiceContainer) (err error)
}

type Pool interface {
	SetAlgo(algo PoolAlgo)
	Start() (err error)
}
