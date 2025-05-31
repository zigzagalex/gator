package commands

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Pointer.CurrentUserName)
	if err != nil {
		fmt.Printf("Error getting follows for user %v", s.Pointer.CurrentUserName)
	}

	PrettyPrintFollows(follows)

	return nil
}
