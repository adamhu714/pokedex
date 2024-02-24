package main

import (
	"fmt"
	"log"

	"github.com/adamhu714/pokedex/internal/pokeapi"
)

func commandMapF() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func commandMapB() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
