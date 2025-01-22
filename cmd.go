package main

type Command struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]Command {
	cmds := make(map[string]Command)
	cmds["help"] = Command{
		name:        "help",
		description: "Displays a help message",
		callback:    CommandHelp,
	}
	cmds["exit"] = Command{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	}

	return cmds
}
