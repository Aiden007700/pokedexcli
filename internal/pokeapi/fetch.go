package pokeapi

import (
	"io"
	"net/http"
)

func checkCacheAndFetch(c *Client, url string) ([]byte, error) {
	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		c.cache.Add(url, dat)
	}
	return dat, nil
}
