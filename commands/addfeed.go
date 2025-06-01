package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {

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
		fmt.Println("Feed could not be added")
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

	follow, err := s.DB.CreateFeedFollow(context.Background(), followParams)

	fmt.Print(feed)
	fmt.Print(follow)

	return nil
}
