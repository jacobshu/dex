package cmd

import (
	"fmt"
	"slices"

	"github.com/jacobshu/dex/internal/pokedex"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*pokedex.Config) error
	Alias       []string
}

func GetCommands() []Command {
	cmds := []Command{
		{
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
			Alias:       []string{"q", "quit"},
		},
		{
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
			Alias:       []string{"h"},
		},
		{
			Name:        "map",
			Description: "List next locations, paginated by 20",
			Callback:    CommandMap,
			Alias:       []string{"m"},
		},
		{
			Name:        "mapb",
			Description: "List previous locations, paginated by 20",
			Callback:    CommandMapb,
			Alias:       []string{"mb"},
		},
	}

	return cmds
}

func RunCommand(cmd string, cfg *pokedex.Config) error {
	cmds := GetCommands()

	cmdExecuted := false
	for _, command := range cmds {
		if command.Name == cmd || slices.Contains(command.Alias, cmd) {
			cmdExecuted = true
			err := command.Callback(cfg)
			if err != nil {
				return fmt.Errorf("error in command %s: %w\n", cmd, err)
			}
		}
	}
	if !cmdExecuted {
		return fmt.Errorf("unrecognized command: %s\n", cmd)
	}
	return nil
}
