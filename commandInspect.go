package main

import (
	"fmt"
)


func commandInspect(args ...string) error {
	pokemon, found := pokedex[args[0]]
	if !found {
		fmt.Println("You have not caught that pokemon")
	} else {
		hp, _ := pokemon.GetStat("hp")
		attack, _ := pokemon.GetStat("attack")
		defense, _ := pokemon.GetStat("defense")
		spAttack, _ := pokemon.GetStat("special-attack")
		spDefense, _ := pokemon.GetStat("special-defense")
		speed, _ := pokemon.GetStat("speed")


		fmt.Printf(
`Name: %s
Height: %d
Weight: %d
Stats:
	-hp: %v
	-attack: %v
	-defense: %v
	-special-attack: %v
	-special-defense: %v
	-speed: %v
`,		
			pokemon.Name,
			pokemon.Height,
			pokemon.Weight,
			hp,
			attack,
			defense,
			spAttack,
			spDefense,
			speed,
		)
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("\t- %v\n", pokemonType.Type.Name)
		}

	}
	return nil
}
