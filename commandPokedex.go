package main

import "fmt"

func commandPokedex(cfg *config) error {
	for k := range cfg.pokeapiClient.Pokedex {
		fmt.Println(" - ", k)
	}

	return nil
}
