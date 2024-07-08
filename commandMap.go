package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandMap(args ...string) error {
	locations, err := pokeapi.GetNextLocations()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
