package main

import (
	"fmt"
	"os"

	"github.com/zigzagalex/gator/commands"
	"github.com/zigzagalex/gator/internal/config"
)

func main() {
	conf, _ := config.Read()
	conf.DBURL = "postgres://example"

	state := &commands.State{Pointer: conf}

	cmdRegistry, _ := commands.InitCommands()

	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := commands.Command{
		Name: commandName,
		Args: commandArgs,
	}

	err := cmdRegistry.Run(state, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
