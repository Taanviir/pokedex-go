package main

import (
	"fmt"
	"os"
)

func helpCommand(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("This is a Pokedex built into the CLI.")
	fmt.Println("Available Commands:")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func exitCommand(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config) error {
	res, err := cfg.pokeAPIClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}

func mapBackCommand(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("no previous page")
	}

	res, err := cfg.pokeAPIClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}
