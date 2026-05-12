package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/slizhunter/pokedexcli/internal/pokecache"
)

func GetLocations(cache pokecache.Cache, pageURL *string) (Location, error) {
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

	// Unmarshals the JSON response into a map and checks for errors
	locationsRes := Location{} // Define a struct to hold the location area data
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return Location{}, err
	}

	cache.Add(url, body) // Adds the response data to the cache with the URL as the key
	//fmt.Printf("Cached data added for: %s\n", url) // Logs a message indicating that data has been added to the cache

	return locationsRes, nil
}

func GetLocationData(cache pokecache.Cache, pageURL *string) (LocationData, error) {
	// Checks if the requested URL is already in the cache and returns the cached data if it exists
	if data, exists := cache.Get(*pageURL); exists {
		//fmt.Printf("Cache hit for URL: %s\n", *pageURL) // Logs a message indicating a cache hit
		locationDataRes := LocationData{}
		err := json.Unmarshal(data, &locationDataRes) // Unmarshals the cached data into a Location struct and checks for errors
		if err != nil {
			return LocationData{}, err
		}
		return locationDataRes, nil
	}

	// Makes a GET request to the PokeAPI to retrieve location area data
	res, err := http.Get(*pageURL)
	if err != nil {
		return LocationData{}, fmt.Errorf("Error fetching location data: %v", err)
	}
	defer res.Body.Close()

	// Reads the response body and checks for errors
	body, err := io.ReadAll(res.Body)
	// Checks if the response status code indicates an error and logs it if so
	if err != nil {
		return LocationData{}, fmt.Errorf("Error reading response body: %v", err)
	}
	cache.Add(*pageURL, body) // Adds the response data to the cache with the URL as the key
	//fmt.Printf("Cached data added for: %s\n", *pageURL) // Logs a message indicating that data has been added to the cache

	// Unmarshals the JSON response into a map and checks for errors
	locationDataRes := LocationData{} // Define a struct to hold the location area data
	err = json.Unmarshal(body, &locationDataRes)
	if err != nil {
		return LocationData{}, fmt.Errorf("Error unmarshaling location data: %v", err)
	}

	return locationDataRes, nil
}

type LocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
