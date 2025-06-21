package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	if !m.userSelected {
		b.WriteString("Select a user:\n\n")
		for i, user := range m.users {
			cursor := " "
			if i == m.userIndex {
				cursor = ">"
			}
			fmt.Fprintf(&b, "%s %s\n", cursor, user.Name)
		}
		return b.String()
	}

	if !m.feedSelected {
		b.WriteString("Select a feed:\n\n")
		for i, feed := range m.feeds {
			cursor := " "
			if i == m.feedIndex {
				cursor = ">"
			}
			fmt.Fprintf(&b, "%s %s\n", cursor, feed.FeedName.String)
		}
		return b.String()
	}

	if m.feedSelected && !m.postSelected {
		b.WriteString("Select a post:\n\n")
		if len(m.posts) == 0 {
			b.WriteString("No posts available for this feed.\n")
		}
		for i, post := range m.posts {
			cursor := " "
			if i == m.postIndex {
				cursor = ">"
			}

			fmt.Fprintf(&b, "%s %s\n", cursor, post.Title)
		}
		return b.String()
	}

	// Final view after selecting feed
	b.WriteString("✅ Feed selected:\n\n")
	b.WriteString(fmt.Sprintf("👉 %s\n", m.Status))
	b.WriteString("\nPress q to quit.\n")
	return b.String()
}
