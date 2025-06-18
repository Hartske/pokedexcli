package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		fmt.Println("Error getting locations")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb() error {
	return nil
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
