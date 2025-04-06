package main

import (
	"fmt"
)

func commandCatch(cfg *config) error {
	if len(cfg.args) == 0 || cfg.args[0] == "" {
		return fmt.Errorf("No pokemon passed, please pass a pokemon: Pokedex > catch charmander")
	}
	pokemonName := cfg.args[0]

	fmt.Printf("Throwing a Pokeball at %v... ", pokemonName)
	if pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName); err != nil {
		return err
	} else {
		fmt.Printf("Caught %v, with a base lv of %v!!!\n", pokemon.Name, pokemon.BaseExperience)
	}
	return nil
}
