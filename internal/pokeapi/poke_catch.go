package pokeapi

import (
	"math/rand"
)

func (c *Client) CatchPokemon(exp int) bool {
	maxThresh := 350
	threshold := maxThresh - exp
	roll := rand.Intn(maxThresh)
	if roll <= threshold {
		return true
	}

	return false
}
