package main

import (
	"time"

	"github.com/taylor-ken/pokedexcli/internal/pokeapi"
	"github.com/taylor-ken/pokedexcli/internal/pokecache"
)

func main() {
	myCache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(5*time.Second, myCache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
