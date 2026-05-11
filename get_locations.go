package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocations(pageURL *string) (Location, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// Makes a GET request to the PokeAPI to retrieve location area data
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	// Reads the response body and checks for errors
	body, err := io.ReadAll(res.Body)
	// Checks if the response status code indicates an error and logs it if so
	if err != nil {
		return Location{}, err
	}

	// Unmarshals the JSON response into a map and checks for errors
	locationsRes := Location{} // Define a struct to hold the location area data
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return Location{}, err
	}

	return locationsRes, nil
}
