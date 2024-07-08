package main

import "fmt"

func commandPokedex(args ...string) error {
	if len(pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}
