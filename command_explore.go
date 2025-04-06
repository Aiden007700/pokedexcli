package main

import "fmt"

func commandExplore(cfg *config) error {

	for _, location := range cfg.args {
		fmt.Println("Exploring", location+"...")
		respShallowExplore, err := cfg.pokeapiClient.ListExplore(location)
		if err != nil {
			return err
		}
		for i, v := range respShallowExplore.Pokemon_encounters {
			if i == 0 {
				fmt.Println("Found Pokemon:")
			}
			fmt.Println(" -", v.Pokemon.Name)
		}
	}

	return nil
}
