package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/slizhunter/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Second)

func GetLocations(pageURL *string) (Location, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// Checks if the requested URL is already in the cache and returns the cached data if it exists
	if data, exists := cache.Get(url); exists {
		//fmt.Printf("Cache hit for URL: %s\n", url) // Logs a message indicating a cache hit
		locationsRes := Location{}
		err := json.Unmarshal(data, &locationsRes) // Unmarshals the cached data into a Location struct and checks for errors
		if err != nil {
			return Location{}, err
		}
		return locationsRes, nil
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
	cache.Add(url, body) // Adds the response data to the cache with the URL as the key
	//fmt.Printf("Cached data added for: %s\n", url) // Logs a message indicating that data has been added to the cache

	// Unmarshals the JSON response into a map and checks for errors
	locationsRes := Location{} // Define a struct to hold the location area data
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return Location{}, err
	}

	return locationsRes, nil
}
