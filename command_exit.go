package main

import (
	"fmt"
	"os"
)

func CommandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
