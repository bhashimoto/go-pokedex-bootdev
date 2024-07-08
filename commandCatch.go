package main

import (
	"fmt"
	"internal/pokeapi"
	"math/rand"
)

var pokedex map[string]pokeapi.Pokemon = make(map[string]pokeapi.Pokemon)
	
func commandCatch(args ...string) error {
	pokemonName := args[0]
	pokemon, err := pokeapi.GetPokemonData(pokemonName)
	if err != nil {
		return err
	}

	rand := rand.Intn(600)
	fmt.Printf("Throwing a Pokebal at %s...\n", pokemon.Name)
	if rand < pokemon.BaseExperience {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		_, ok := pokedex[pokemon.Name]
		if !ok {
			pokedex[pokemon.Name] = pokemon
		}
	}

	return nil
}


