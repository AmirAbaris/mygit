package mygit

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
		command.Init()
	default:
		fmt.Println("unknown command")
	}

}
