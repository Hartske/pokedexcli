package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}
	locationRes, err := cfg.pokeapiClient.GetLocation(&args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", locationRes.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
