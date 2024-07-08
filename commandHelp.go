package main

import "fmt"

func commandHelp(args ...string) error {
	fmt.Printf("Usage:\n")
	for _, v := range getCommands() {
		fmt.Printf("\t%v\t\t%v\n",
			v.name,
			v.description)
	} 
	return nil
}
