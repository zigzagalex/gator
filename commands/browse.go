package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/zigzagalex/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	const defaultPostLimit = 3
	n_posts := defaultPostLimit
	if len(cmd.Args) > 0 && strings.TrimSpace(cmd.Args[0]) != "" {
		num, err := strconv.Atoi(strings.TrimSpace(cmd.Args[0]))
		if err != nil {
			return fmt.Errorf("Invalid number of posts: %w", err)
		}
		n_posts = num
	}

	posts, err := s.DB.GetPosts(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting posts for %v: %v", user.Name, err)
	}

	for i, post := range posts {
		if i >= n_posts {
			break
		}
		fmt.Printf("* %s | %s\n  ↳ %s\n", post.PublishedAt.Format("2006-01-02"), post.Title, post.Url)
	}
	return nil

}
