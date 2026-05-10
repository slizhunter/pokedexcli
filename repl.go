package main

import (
	"strings"
)

func cleanInput(text string) []string {
	newText := strings.Fields(strings.ToLower(text))
	return newText
}
