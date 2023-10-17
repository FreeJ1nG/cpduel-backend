package pool

import (
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
)

type pool struct {
	algo     chan interfaces.PoolAlgo
	services interfaces.ServiceContainer
}

func NewPool(services interfaces.ServiceContainer) *pool {
	return &pool{
		algo:     make(chan interfaces.PoolAlgo),
		services: services,
	}
}

func (p *pool) SetAlgo(algo interfaces.PoolAlgo) {
	p.algo <- algo
}

func (p *pool) Start() (err error) {
	for {
		algo := <-p.algo
		err = algo.Run(p.services)
		if err != nil {
			fmt.Printf("pool error: %s\n", err.Error())
		}
	}
}
