package main

import (
	"time"

	"github.com/hartske/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	dex := pokeapi.NewPokedex()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       dex,
	}

	startRepl(cfg)
}
