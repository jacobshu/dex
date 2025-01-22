package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jacobshu/dex/pokedex"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			sanitized := cleanInput(input)
			if len(sanitized) == 0 {
				continue
			}
			cmd := sanitized[0]
			cmds := getCommands()
			config := pokedex.Config{}
			if command, ok := cmds[cmd]; ok {
				err := command.callback(&config)
				if err != nil {
					fmt.Print(fmt.Errorf("error in command %s: %w", cmd, err))
				}
			} else {
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
