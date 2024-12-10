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
	callback    func()
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
	}
}

func startRepl() {
	const PROMPT = "> "
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

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
		command, ok := commands[commandEntered]
		if !ok {
			fmt.Println("Unknown command.")
			continue
		}

		command.callback()
	}
}

func cleanInput(input string) []string {
	lowercased := strings.ToLower(input)
	words := strings.Fields(lowercased)
	return words
}
