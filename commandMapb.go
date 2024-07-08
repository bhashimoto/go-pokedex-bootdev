package main

import (
	"fmt"
	"internal/pokeapi"
	"log"
)


func commandMapb(args ...string) error {
	locations, err := pokeapi.GetPreviousLocations()
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
