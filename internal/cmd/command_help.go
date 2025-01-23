package cmd

import (
	"fmt"
	"strings"

	"github.com/jacobshu/dex/internal/pokedex"
)

func CommandHelp(cfg *pokedex.Config) error {
	fmt.Printf("Welcome to the Pokedex!\n\nCommands:\n")

	sepWidth := 8
	allCmds := GetCommands()
	for _, cmd := range allCmds {
		fmt.Printf("    %s:%s%s\n",
			cmd.Name,
			strings.Repeat(" ", sepWidth-len(cmd.Name)),
			cmd.Description)
	}
	fmt.Printf("\n")
	return nil
}
