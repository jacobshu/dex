package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			err := cmds[cmd].callback()
			if err != nil {
				fmt.Print(fmt.Errorf("error in command %s: %w", cmd, err))
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
