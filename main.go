package main

import (
	"fmt"
	"os"
)

func main() {
	// os Args: what did the user type after the program name?
	args := os.Args

	command := args[1]

	if command == "init" {
		fmt.Println("hit")
	}
}
