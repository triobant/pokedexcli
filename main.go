package main

import (
    "time"

    "github.com/triobant/pokedexcli/internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
    cfg := &config{
        pokeapiClient: pokeClient,
        caughtPokemon: map[string]pokeapi.Pokemon{},
    }

    startRepl(cfg)
}
