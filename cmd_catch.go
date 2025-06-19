package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}

	pokemonRes, err := cfg.pokeapiClient.GetPokemon(&args[0])
	if err != nil {
		return err
	}

	name := pokemonRes.Species.Name
	exp := pokemonRes.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	catchRes := cfg.pokeapiClient.CatchPokemon(exp)

	if catchRes {
		cfg.pokedex.Add(name, pokemonRes)
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
