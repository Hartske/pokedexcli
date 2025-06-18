package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	locationRes, err := cfg.pokeapiClient.GetLocation(&args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", locationRes.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
