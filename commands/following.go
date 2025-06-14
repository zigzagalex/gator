package commands

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Pointer.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting follows for user %v: %v", s.Pointer.CurrentUserName, err)
	}

	PrettyPrintFollows(follows)

	return nil
}
