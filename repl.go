package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	newInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		newInput.Scan()

		words := cleanInput(newInput.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	clean := strings.Fields(input)
	return clean
}
