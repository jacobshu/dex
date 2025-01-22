package main

import (
	"fmt"
	"os"

	"github.com/jacobshu/dex/pokedex"
)

func CommandExit(cfg *pokedex.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
