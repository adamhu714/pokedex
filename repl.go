package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexCLI > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			invalidCommand()
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if !exists {
			invalidCommand()
			continue
		}
		if command.name == "exit" {
			break
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func invalidCommand() {
	fmt.Println("Please input a valid command.")
}

func cleanInput(text string) []string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next page of location areas",
			callback:    callbackMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of location areas",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore",
			description: "",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program.",
			callback:    callbackExit,
		},
	}
}
