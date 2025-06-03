package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/zigzagalex/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	url := cmd.Args[0]

	// Get feed
	feed, err := s.DB.GetFeed(context.Background(), url)
	if err != nil {
		fmt.Println("Error getting feed info")
		os.Exit(1)
		return nil
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	s.DB.DeleteFeedFollow(context.Background(), params)

	fmt.Printf("Feed %v is now unfollowed by user %v", feed.Name, user.Name)

	return nil

}
