package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zigzagalex/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		fmt.Println("Please provide a username.")
		os.Exit(1)
		return nil
	}

	username := cmd.Args[0]

	userparams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	// Create user
	user, err := s.DB.CreateUser(context.Background(), userparams)
	if err != nil {
		if isUniqueViolation(err) {
			fmt.Printf("User with name '%s' already exists.\n", username)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Set current user
	s.Pointer.SetUser(username)

	fmt.Printf("User %v was created\n", username)
	log.Printf("DEBUG: user %+v\n", user)
	return nil
}

func isUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505"
	}
	return false
}
