package main

import (
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	spaceRegex := regexp.MustCompile(`\s+`)
	clearString := strings.ToLower(strings.TrimSpace(spaceRegex.ReplaceAllString(text, " ")))

	commands := strings.Split(clearString, " ")

	return commands
}

func main() {

}
