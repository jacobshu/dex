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
			if cmd == "quit" || cmd == "q" {
				fmt.Print("bye 👋\n")
				return
			}
			fmt.Printf("Your command was: %s\n", cmd)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
