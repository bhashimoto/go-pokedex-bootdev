package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

	
var locationsConfig = config{
	next: "https://pokeapi.co/api/v2/location-area",
	previous: nil,
}

func GetNextLocations() ([]string, error) {
	return getLocations(locationsConfig.next, &locationsConfig)
}

func GetPreviousLocations() ([]string, error) {
	if locationsConfig.previous == nil {
		return []string{"already at the top of the list"}, nil
	}
	return getLocations(*locationsConfig.previous, &locationsConfig)
}

func getLocations(url string, c *config) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		return []string{}, errors.New("invalid status code")
	}
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	locs := locations{}
	err = json.Unmarshal(body, &locs)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}


	c.previous = locs.Previous
	c.next = locs.Next

	locations := []string{}
	for _, location := range locs.Results {
		locations = append(locations, location.Name)
	}

	return locations, nil
}

type config struct {
	next		string
	previous	*string
}

type locations struct {
	Count		int	`json:"count"`
	Next		string	`json:"next"`
	Previous	*string	`json:"previous"`
	Results		[]struct {
		Name	string	`json:"name"`
		URL	string	`json:"url"`
	}
}
