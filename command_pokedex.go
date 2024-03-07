package main

import (
	"fmt"
	"sort"
	"strconv"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Pokemon in pokedex:")
	keys := make([]int, 0)
	for k := range cfg.caughtPokemon {
		if i, err := strconv.Atoi(k); err == nil {
			keys = append(keys, i)
		}
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d: %s\n", k, cfg.caughtPokemon[fmt.Sprint(k)].Name)
	}
	fmt.Println()
	return nil
}
