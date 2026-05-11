package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := config{Next: nil, Previous: nil} // Initializes the config struct with empty URLs

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
			err := command.callback(&config)
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
	callback    func(*config) error // Function to execute when the command is called
}

// Struct to contain the Next and Previous URLs to paginate through location areas
type config struct {
	Next     *string // URL for the next page of location areas
	Previous *string // URL for the previous page of location areas
}

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 map locations, subsequent calls will display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 map locations",
			callback:    commandMapb,
		},
	}
}
