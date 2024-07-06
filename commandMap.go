package main

import (
	"fmt"
	"internal/pokeapi"
	"log"
)

func commandMap() error {
	locations, err := pokeapi.GetNextLocations()
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
