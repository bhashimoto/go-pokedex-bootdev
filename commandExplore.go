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

	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	
	return nil
}
