package main

import (
	"time"

	pokeapi "github.com/Taanviir/pokedex/internal/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
