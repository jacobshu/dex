package cmd

import (
	"fmt"
	"os"

	"github.com/jacobshu/dex/internal/pokedex"
)

func CommandExit(cfg *pokedex.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
