package commands

import (
	"context"
	"fmt"
	"os"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("Error when getting feeds from database")
		os.Exit(1)
		return nil
	}

	PrettyPrintFeeds(feeds)

	return nil
}
