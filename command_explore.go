package main

import "fmt"

func callbackExplore(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationArea("")
	if err != nil {
		return err
	}
	fmt.Println(resp.PokemonEncounters)
	return nil
}
