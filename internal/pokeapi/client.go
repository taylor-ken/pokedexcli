package pokeapi

import (
	"net/http"
	"time"

	"github.com/taylor-ken/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      cache,
	}
}
