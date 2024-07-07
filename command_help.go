package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	keys := []string{
		"catch",
		"inspect",
		"pokedex",
		"map",
		"mapb",
		"explore",
		"help",
		"exit",
	}
	for _, k := range keys {
		fmt.Printf("%s: %s\n", commands[k].name, commands[k].description)
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("For more information, please refer to the documentation at:")
	fmt.Println("https://github.com/adamhu714/pokedex")
	fmt.Println()

	return nil
}
