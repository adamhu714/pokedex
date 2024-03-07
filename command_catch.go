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
	fmt.Printf("base experience %v, roll %v, threshold %v\n", pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	cfg.caughtPokemon[fmt.Sprint(pokemon.ID)] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
