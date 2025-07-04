package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := Locations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return Locations{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locationsRes := Locations{}
	err = json.Unmarshal(dat, &locationsRes)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, dat)
	return locationsRes, nil
}
