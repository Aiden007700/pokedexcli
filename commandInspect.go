package main

import "fmt"

func commandInspect(cfg *config) error {
	if len(cfg.args) == 0 || cfg.args[0] == "" {
		return fmt.Errorf("No pokemon passed, please pass a pokemon: Pokedex > inspect charmander \n")
	}
	pokemonName := cfg.args[0]

	if pokemon, ok := cfg.pokeapiClient.Pokedex[pokemonName]; !ok {
		return fmt.Errorf("No pokemon with that name found in Pokedex")
	} else {
		fmt.Println(pokemon)
		return nil
	}
}
