package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: addfeed <name> <url>")
	}
	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("Feed could not be added: %v\n", err)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), followParams)

	fmt.Printf("Feed added and followed: %v\n", feed.Name)

	return nil
}
