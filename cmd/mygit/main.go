package main

import (
	"fmt"
	"os"

	"github.com/AmirAbaris/mygit/cmd/internal/command"
)

func main() {
	// os Args: what did the user type after the program name?
	args := os.Args

	inputCommand := args[1]

	switch inputCommand {
	case "init":
		err := command.Init()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("unknown command")

	}

}
