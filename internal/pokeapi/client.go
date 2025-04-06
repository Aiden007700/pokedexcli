package pokeapi

import (
	"github.com/aiden007700/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
	Pokedex    map[string]RespShallowPokemon
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokecache.NewCache(5*time.Second, false),
		Pokedex: make(map[string]RespShallowPokemon),
	}
}
