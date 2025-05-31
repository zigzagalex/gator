package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error getting users from database")
	}

	currentUser := s.Pointer.CurrentUserName

	for _, user := range users {
		line := "* " + user.Name
		if user.Name == currentUser {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}
