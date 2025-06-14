package commands

import (
	"fmt"
	"time"

	"github.com/zigzagalex/gator/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Parse duration error: %v\n", err)
	}

	fmt.Printf("Collecting feeds every %v", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err = rss.ScrapeFeeds(s.DB)
		if err != nil {
			return fmt.Errorf("Scrape error: %v\n", err)
		}
	}

}
