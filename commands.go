package main

import (
	"fmt"
	"os"
)

func helpCommand() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("This is a Pokedex built into the CLI.")
	fmt.Println("Available Commands:")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println("")
}

func exitCommand() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}
