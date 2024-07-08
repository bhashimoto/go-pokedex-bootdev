package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")

		cmd, ok := commands[input[0]]
		if !ok {
			fmt.Println("\nInvalid command")
			continue
		}
		
		cmd.callback(input[1:]...)
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(args ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:		"help",
			description:	"Displays a help message",
			callback:	commandHelp,
			
		},
		"exit": {
			name:		"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
		"map": {
			name:		"map",
			description:	"gets the next 20 location areas available",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description:	"gets the next 20 location areas available",
			callback:	commandMapb,
		},
		"explore": {
			name:		"explore",
			description:	"gets the pokemon found in the specified area",
			callback:	commandExplore,
		},
	}
}
