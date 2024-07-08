package pokeapi

import (
	"encoding/json"
	"errors"
	"internal/pokecache"
	"io"
	"log"
	"net/http"
	"time"
)

var cache pokecache.Cache = pokecache.NewCache(5 * time.Minute)
const baseURL string = "https://pokeapi.co/api/v2/"


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

func GetLocationPokemons(area string) ([]string, error) {
	url := baseURL + "location-area/" + area
	body, found := cache.Get(url)	
	if !found{
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return []string{}, err
		}

		data, err := io.ReadAll(res.Body)
		body = data
		res.Body.Close()
		cache.Add(url, body)
		if err != nil {
			log.Fatal(err)
			return []string{}, err
		}
	}

	details := locationAreaDetails{}
	err := json.Unmarshal(body, &details)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	pokemons := []string{}

	for _, pokemon := range details.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}
	return pokemons, nil

}

func getLocations(url string, c *config) ([]string, error) {
	body, found := cache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		data, err := io.ReadAll(res.Body)
		body = data
		res.Body.Close()
		cache.Add(url, body)
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
			return []string{}, errors.New("invalid status code")
		}
		if err != nil {
			log.Fatal(err)
			return []string{}, err
		}
	}


	locs := locations{}
	err := json.Unmarshal(body, &locs)
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

type locationAreaDetails struct {
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
