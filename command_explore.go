package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide one location area")
	}

	locationAreaName := args[0]

	locationAreaResp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println(locationAreaResp.PokemonEncounters)
	return nil
}
