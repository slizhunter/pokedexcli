package main

import (
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	// Check if the user provided a location name
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location name. Syntax: explore [location name]")
	}

	baseURL := "https://pokeapi.co/api/v2/location-area/"
	locationName := args[0]
	locationURL := baseURL + locationName
	locationData, err := GetLocationData(config.Cache, &locationURL)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	// Iterates through the "results" field of the map and prints the name of each location area
	for _, value := range locationData.PokemonEncounters {
		fmt.Printf(" - %s\n", value.Pokemon.Name)
	}

	return nil
}
