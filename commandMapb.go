package main

import (
	"fmt"
	"internal/pokeapi"
)


func commandMapb(args ...string) error {
	locations, err := pokeapi.GetPreviousLocations()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
