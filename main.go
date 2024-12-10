package main

import (
	"time"

	pokeapi "github.com/Taanviir/pokedex/internal/pokeapi"
)

type config struct {
	pokeAPIClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeAPIClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
