package commands

import (
	"context"
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

	// Check if user exists in database
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("No user found with name: %s\n", username)
		os.Exit(1)
	}

	err1 := s.Pointer.SetUser(username)
	if err1 != nil {
		fmt.Println("Error while setting current user.")
		return nil
	}

	fmt.Printf("Logged in as: %v\n", username)
	return nil
}
