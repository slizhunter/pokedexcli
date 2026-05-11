package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	mapArea, err := GetLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = mapArea.Next
	config.Previous = mapArea.Previous

	// Iterates through the "results" field of the map and prints the name of each location area
	for _, value := range mapArea.Results {
		fmt.Println(value.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	}

	mapArea, err := GetLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = mapArea.Next
	config.Previous = mapArea.Previous

	// Iterates through the "results" field of the map and prints the name of each location area
	for _, value := range mapArea.Results {
		fmt.Println(value.Name)
	}

	return nil
}
