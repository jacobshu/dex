package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jacobshu/dex/internal/cmd"
	"github.com/jacobshu/dex/internal/pokedex"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := pokedex.Config{}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			sanitized := cleanInput(input)
			if len(sanitized) == 0 {
				continue
			}
			com := sanitized[0]
			err := cmd.RunCommand(com, &config)
			if err != nil {
				fmt.Print(fmt.Errorf("error in command %s: %w", com, err))
			} else {
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
