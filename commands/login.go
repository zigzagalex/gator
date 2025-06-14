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
		return fmt.Errorf("No user found with name %s: %v\n", username, err)
	}

	err = s.Pointer.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error while setting current user: %v", err)
	}

	fmt.Printf("Logged in as: %v\n", username)
	return nil
}
