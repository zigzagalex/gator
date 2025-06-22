package commands

import (
	"fmt"
	"time"

	"log"

	"github.com/zigzagalex/gator/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		log.Printf("parse duration error: %v", err)
		return fmt.Errorf("parse duration error: %w", err)
	}

	log.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = rss.ScrapeFeeds(s.DB)
		if err != nil {
			log.Printf("Scrape error: %v\n", err)
		}
	}
}
