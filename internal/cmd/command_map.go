package cmd

import (
	"fmt"
	"github.com/jacobshu/dex/internal/pokedex"
)

func CommandMap(cfg *pokedex.Config) error {
	if cfg.Next == "" {
		pokedex.GetLocations(cfg)
	} else {
		pokedex.GetNextLocations(cfg)
	}

	for _, area := range cfg.Areas {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
