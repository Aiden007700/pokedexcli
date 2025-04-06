package pokeapi

import (
	"encoding/json"
)

// ListExplore -
func (c *Client) ListExplore(location string) (RespShallowExplore, error) {
	url := baseURL + locationAreaPath + location

	dat, err := checkCacheAndFetch(c, url)
	exploreResp := RespShallowExplore{}

	if err != nil {
		return exploreResp, err
	}

	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return exploreResp, err
	}

	return exploreResp, nil
}
