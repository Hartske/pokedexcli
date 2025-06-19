package pokeapi

import (
	"sort"
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

func (p *Pokedex) Get(key string) (Pokemon, bool) {
	p.mux.Lock()
	defer p.mux.Unlock()
	mon, ok := p.caught[key]
	return mon, ok
}

func (p *Pokedex) GetAll() []Pokemon {
	p.mux.Lock()
	defer p.mux.Unlock()

	var pokemonList []Pokemon
	for _, pokemon := range p.caught {
		pokemonList = append(pokemonList, pokemon)
	}
	sort.Slice(pokemonList, func(i, j int) bool {
		return pokemonList[i].ID < pokemonList[j].ID
	})

	return pokemonList
}
