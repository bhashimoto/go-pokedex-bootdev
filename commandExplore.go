package main

import (
	"fmt"
	"internal/pokeapi"
	"log"
)

func commandExplore(args ...string) error {
	location := args[0]
	pokemons, err := pokeapi.GetLocationPokemons(location)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Exploring %s\n", location)
	for _, pokemon := range pokemons {
		fmt.Printf(" - %v\n", pokemon)
	}
	
	return nil
}
