package pokeapi

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"time"
)

var (
	rng               = rand.New(rand.NewSource(time.Now().UnixNano()))
	ErrPokemonEscaped = errors.New("pokemon broke free and escaped")
)

func (c *Client) FetchPokemon(pokemon string) (RespShallowPokemon, error) {
	url := baseURL + pokemonPath + pokemon
	dat, err := checkCacheAndFetch(c, url)
	pokemonResp := RespShallowPokemon{}

	if err != nil {
		return pokemonResp, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return pokemonResp, err
	}

	return pokemonResp, nil
}

func (c *Client) CatchPokemon(pokemonName string) (RespShallowPokemon, error) {
	pokemon, err := c.FetchPokemon(pokemonName)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	catchProbability := calculateCatchProbability(pokemon.BaseExperience)

	if rng.Float64() <= catchProbability {
		c.Pokedex[pokemon.Name] = pokemon
		return pokemon, nil
	}
	return RespShallowPokemon{}, ErrPokemonEscaped
}

func calculateCatchProbability(baseExp int) float64 {
	// These constants determine the catch probability curve
	const (
		minProbability = 0.2  // Minimum catch probability (20%)
		maxProbability = 0.8  // Maximum catch probability (80%)
		midExpPoint    = 150  // Base experience that gives 50% catch rate
		steepness      = 0.01 // Controls how quickly probability changes with base exp
	)

	// Logistic function that maps base experience to catch probability
	// Higher base exp = lower probability, with floor of 20%
	probability := maxProbability - ((maxProbability - minProbability) /
		(1 + math.Exp(-steepness*(float64(baseExp)-midExpPoint))))

	return probability
}
