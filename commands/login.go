package commands

import (
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		fmt.Println("Please provide a username.")
		os.Exit(1)
		return nil
	}

	username := cmd.Args[0]

	err := s.Pointer.SetUser(username)
	if err != nil {
		fmt.Println("Error while setting current user.")
		return nil
	}

	fmt.Printf("User has been set to: %v\n", username)
	return nil
}
