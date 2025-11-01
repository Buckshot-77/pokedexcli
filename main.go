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

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		terms := cleanInput(input)

		fmt.Printf("Your command was: %s\n", terms[0])
	}
}
