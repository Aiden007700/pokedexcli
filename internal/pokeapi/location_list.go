package pokeapi

import (
	"encoding/json"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, err := checkCacheAndFetch(c, url)
	locationsResp := RespShallowLocations{}

	if err != nil {
		return locationsResp, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return locationsResp, err
	}

	return locationsResp, nil
}
