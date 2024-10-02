package main

import (
    "errors"
    "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    fmt.Println("Your Pokdex:")
    for _, p := range cfg.caughtPokemon {
        fmt.Prinf(" - %s\n", p.Name)
    }
    return nil
}
