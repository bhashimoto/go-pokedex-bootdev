package main

import (
	"bufio"
	"fmt"
	"os"
)



func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := scanner.Text()
		cmd, ok := commands[input]
		if !ok {
			fmt.Println("\nInvalid command")
			continue
		}
		
		cmd.callback()
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
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
	}
}
