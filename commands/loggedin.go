package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/zigzagalex/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		username := s.Pointer.CurrentUserName
		if username == "" {
			fmt.Println("Error: No user is currently logged in.")
			return errors.New("Must be logged in")
		}
		user, err := s.DB.GetUser(context.Background(), username)
		if err != nil {
			fmt.Printf("Error fetching user '%s': %v\n", username, err)
			return err
		}

		return handler(s, cmd, user)
	}

}
