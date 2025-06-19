package pokeapi

import (
	"sync"
)

type Pokedex struct {
	caught map[string]Pokemon
	mux    *sync.Mutex
}

func NewPokedex() Pokedex {
	p := Pokedex{
		caught: make(map[string]Pokemon),
		mux:    &sync.Mutex{},
	}
	return p
}

func (p *Pokedex) Add(key string, mon Pokemon) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.caught[key] = mon
}
