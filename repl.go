package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Pokedex CLI help menu.",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application gracefully.",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Lists out locations in the pokemon world.",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous page of location areas.",
			callback:    mapBackCommand,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "List out all the pokemon roaming in the area.",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Random chance to catch the pokemon and add to the pokedex.",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details of a caught pokemon.",
			callback:    inspectCommand,
		},
	}
}

func startRepl(cfg *config) {
	const PROMPT = "> "
	scanner := bufio.NewScanner(os.Stdin)

	// prompt loop
	for {
		fmt.Print(PROMPT)

		scanner.Scan()
		text := scanner.Text()

		tokens := cleanInput(text)
		if len(tokens) == 0 {
			continue
		}

		commandEntered := tokens[0]
		args := []string{}
		if len(tokens) > 1 {
			args = tokens[1:]
		}

		command, exists := getCommands()[commandEntered]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	lowercased := strings.ToLower(input)
	words := strings.Fields(lowercased)
	return words
}
