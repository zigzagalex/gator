package commands

import (
	"fmt"
	"sort"
)

func HandlerHelp(cmdRegistry *Commands) func(*State, Command) error {
	return func(s *State, c Command) error {
		fmt.Printf("\nAvailable commands:\n")

		// Sort commands alphabetically
		var names []string
		for name := range cmdRegistry.Handlers {
			names = append(names, name)
		}
		sort.Strings(names)

		// Pretty print each command
		for _, name := range names {
			cmd := cmdRegistry.Handlers[name]
			fmt.Printf("• %s\n", name)
			fmt.Printf("  %s\n", cmd.Description)
			if cmd.Usage != "" {
				fmt.Printf("  usage: %s\n", cmd.Usage)
			}
			fmt.Println()
		}

		fmt.Println("Type 'exit' or 'quit' to leave.")
		return nil
	}
}
