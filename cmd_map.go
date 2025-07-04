package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationRes.Next
	cfg.prevLocationsURL = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationRes.Next
	cfg.prevLocationsURL = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
