package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

func helpCommand(cfg *config, args ...string) error {
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

func exitCommand(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config, args ...string) error {
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

func mapBackCommand(cfg *config, args ...string) error {
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

func exploreCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]
	location, err := cfg.pokeAPIClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - '%s'\n", pokemon.Pokemon.Name)
	}

	return nil
}

func catchCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeAPIClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 50
	randNum := rand.IntN(pokemon.BaseExperience)
	if randNum > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return errors.New("")
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}

func inspectCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("this pokemon has not been caught")
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Println("Name:" + pokemon.Name)
	fmt.Println("Weight:" + string(pokemon.Weight))
	fmt.Println("Height:" + string(pokemon.Height))
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf("  - %s\n", stat.Type.Name)
	}

	return nil
}

func pokedexCommand(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
