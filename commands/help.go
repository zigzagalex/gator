package commands

import "fmt"

func HandlerHelp(cmdRegistry *Commands) func(*State, Command) error {
	return func(s *State, c Command) error {
		fmt.Println("** Available commands **")
		for name := range cmdRegistry.Handlers {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println("Type a command followed by arguments if needed.")
		return nil
	}
}
