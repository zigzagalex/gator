package rss

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zigzagalex/gator/internal/database"
)

func ScrapeFeeds(db *database.Queries) error {
	ctx := context.Background()
	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error getting next feed: %w", err)
	}

	err = db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("Error marking feed as fetched: %w", err)
	}

	parsedFeed, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, item := range parsedFeed.Channel.Item {
		// Declaring layout constant
		const layout = time.RFC1123Z
		pub_date, err := time.Parse(layout, item.PubDate)
		if err != nil {
			return fmt.Errorf("Error parsing post pub date: %v", err)
		}
		desc := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}
		post_param := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pub_date,
			FeedID:      feed.ID,
		}

		post, err := db.CreatePost(ctx, post_param)
		if err != nil {
			// Check for "duplicate key" error (PostgreSQL code 23505)
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				// Duplicate URL – skip silently
				continue
			}
			// Other errors
			fmt.Printf("Error creating post (%s): %v\n", item.Link, err)
			continue
		}
		log.Printf("%v\n", post.Title)
	}

	return nil
}
