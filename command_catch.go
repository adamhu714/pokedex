package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide one pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	threshold := 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
