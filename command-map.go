package main

import (
	"log"

	"github.com/adamhu714/pokedex/internal/pokeapi"
)

func commandMapF() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	print(resp)
	return nil
}

func commandMapB() error {
	return nil
}
