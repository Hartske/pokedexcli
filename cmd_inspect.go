package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	mon, err := cfg.pokedex.Get(args[0])
	if err != true {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", mon.Name)
	fmt.Printf("Height: %v\n", mon.Height)
	fmt.Printf("Weight: %v\n", mon.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %v\n", mon.Stats[0].BaseStat)
	fmt.Printf("  -attack: %v\n", mon.Stats[1].BaseStat)
	fmt.Printf("  -defense: %v\n", mon.Stats[2].BaseStat)
	fmt.Printf("  -special-attack: %v\n", mon.Stats[3].BaseStat)
	fmt.Printf("  -special-defense: %v\n", mon.Stats[4].BaseStat)
	fmt.Printf("  -speed: %v\n", mon.Stats[5].BaseStat)
	fmt.Println("Types:")
	for i := 0; i < len(mon.Types); i++ {
		fmt.Printf("  - %s\n", mon.Types[i].Type.Name)
	}

	return nil
}
