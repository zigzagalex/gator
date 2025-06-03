package rss

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		const layout = "2006-Jan-02"
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

		post, err := db.CreatePost(context.Background(), post_param)
		if err != nil {
			fmt.Printf("Error creating post: %v", err)
		}
		fmt.Printf("%v\n", post.Title)
	}

	return nil
}
