package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//Sets up ability to scan for user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//Prints the program name to the console for user interaction
		fmt.Print("Pokedex > ")
		//Scans for user input and saves it for Text() to use
		scanner.Scan()
		//Cleans the input
		cleanedInput := cleanInput(scanner.Text())

		// If the cleaned input is empty, skip to the next iteration of the loop
		if len(cleanedInput) == 0 {
			continue
		}

		//Checks if the first word of the cleaned input matches a command and executes it if it does
		command, exists := getCommands()[cleanedInput[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	newText := strings.Fields(strings.ToLower(text))
	return newText
}

// Struct to define and describe command types
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Map of supported commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
