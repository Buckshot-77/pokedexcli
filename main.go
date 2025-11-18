package main

import (
	"time"

	"github.com/Buckshot-77/pokedexcli/internal/commands"
	"github.com/Buckshot-77/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
