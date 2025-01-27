package pokedex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseurl = "https://pokeapi.co/api/v2"

func GetLocations(cfg *Config) error {
	fullurl := baseurl + "/location-area/"
	req, err := http.NewRequest(http.MethodGet, fullurl, nil)
	if err != nil {
		return fmt.Errorf("GetLocations: error making request: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("GetLocations: error sending request: %w\n", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("GetLocations: error decoding map response: %w\n", err)
	}

	return nil
}

func GetNextLocations(cfg *Config) error {
	req, err := http.NewRequest(http.MethodGet, cfg.Next, nil)
	if err != nil {
		return fmt.Errorf("GetNextLocations: error creating request: %w\n", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("GetNextLocations: error sending request: %w\n", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("GetNextLocations: error decoding map response: %w\n", err)
	}

	return nil

}

func GetPrevLocations(cfg *Config) error {
	req, err := http.NewRequest(http.MethodGet, cfg.Previous, nil)
	if err != nil {
		return fmt.Errorf("GetPrevLocations: error creating request: %w\n", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("GetPrevLocations: error sending request: %w\n", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("GetPrevLocations: error decoding map response: %w\n", err)
	}

	return nil

}
