package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	url := cmd.Args[0]

	// Get feed
	feed, err := s.DB.GetFeed(context.Background(), url)
	if err != nil {
		fmt.Println("Error getting feed info")
		os.Exit(1)
		return nil
	}
	
	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		fmt.Println("Error creating the follow for the feed")
		os.Exit(1)
		return nil
	}

	fmt.Printf("Feed %v is now followed by user %v", feedFollow[0].FeedName, feedFollow[0].UserName)

	return nil
}
