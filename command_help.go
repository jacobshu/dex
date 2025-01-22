package main

import (
	"fmt"

	"github.com/jacobshu/dex/pokedex"
)

func CommandHelp(cfg *pokedex.Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
