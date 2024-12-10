package main

import pokeapi "github.com/Taanviir/pokedex/internal/pokeapi"

type config struct {
	pokeAPIClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeAPIClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
