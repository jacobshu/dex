package main

import (
	"github.com/jacobshu/dex/pokedex"
)

type Command struct {
	name        string
	description string
	callback    func(*pokedex.Config) error
}

func getCommands() map[string]Command {
	exit := Command{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	}

	cmds := make(map[string]Command)
	cmds["help"] = Command{
		name:        "help",
		description: "Displays a help message",
		callback:    CommandHelp,
	}
	cmds["exit"] = exit
	cmds["quit"] = exit
	cmds["q"] = exit
	cmds["map"] = Command{
		name:        "map",
		description: "List locations, paginated by 20",
		callback:    CommandMap,
	}
	cmds["mapb"] = Command{
		name:        "map",
		description: "List prior locations, paginated by 20",
		callback:    CommandMapb,
	}

	return cmds
}
