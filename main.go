package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	spaceRegex := regexp.MustCompile(`\s+`)
	clearString := strings.ToLower(strings.TrimSpace(spaceRegex.ReplaceAllString(text, " ")))

	terms := strings.Split(clearString, " ")

	return terms
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break // EOF or error
		}

		input := scanner.Text()
		terms := cleanInput(input)

		if len(terms) == 0 {
			continue
		}

		commandName := terms[0]

		if command, exists := CommandsMap[commandName]; exists {
			if err := command.callback(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command '%s'. Type 'help' for available commands.\n", commandName)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
