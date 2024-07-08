package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandExplore(args ...string) error {
	location := args[0]
	pokemons, err := pokeapi.GetLocationPokemons(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", location)
	for _, pokemon := range pokemons {
		fmt.Printf(" - %v\n", pokemon)
	}
	
	return nil
}
