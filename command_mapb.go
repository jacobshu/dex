package main

import (
	"fmt"
	"github.com/jacobshu/dex/pokedex"
)

func CommandMapb(cfg *pokedex.Config) error {
	if cfg.Previous == "" {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	pokedex.GetPrevLocations(cfg)
	for _, area := range cfg.Areas {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
