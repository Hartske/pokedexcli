package main

import (
	"errors"
	"fmt"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}

	pokemonRes, err := cfg.pokeapiClient.GetPokemon(&args[0])
	if err != nil {
		return errors.New("Pokemon doesn't exist")
	}

	name := pokemonRes.Species.Name
	exp := pokemonRes.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	catchRes := cfg.pokeapiClient.CatchPokemon(exp)

	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".\n")
	time.Sleep(1 * time.Second)
	fmt.Println("")

	if catchRes {
		cfg.pokedex.Add(name, pokemonRes)
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You can now inspect this pokemon!")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
