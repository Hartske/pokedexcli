package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(mon *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *mon

	if val, ok := c.cache.Get(url); ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(val, &pokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokenmnRes := Pokemon{}
	err = json.Unmarshal(dat, &pokenmnRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokenmnRes, nil
}
