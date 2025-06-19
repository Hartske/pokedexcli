package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	caught := cfg.pokedex.GetAll()

	if len(caught) < 1 {
		fmt.Println("you haven't caught any Pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range caught {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
