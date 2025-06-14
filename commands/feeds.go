package commands

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error when getting feeds from database: %v", err)
	}

	PrettyPrintFeeds(feeds)

	return nil
}
