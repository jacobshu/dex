package main

import (
	"fmt"
	"github.com/jacobshu/dex/pokedex"
)

func CommandMap(cfg *pokedex.Config) error {
	fmt.Printf("MAP COMMAND:\nPREV%v, NEXT%v\n", cfg.Previous, cfg.Next)
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
