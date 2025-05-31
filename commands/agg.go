package commands

import (
	"context"
	"fmt"

	"github.com/zigzagalex/gator/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	_, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		fmt.Println("Error fetching the rss Feed.")
	}

	return nil
}
