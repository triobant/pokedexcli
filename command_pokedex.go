package main

import (
    "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    fmt.Println("Your Pokdex:")
    for _, p := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", p.Name)
    }
    return nil
}
