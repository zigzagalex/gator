package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case !m.userSelected:
			switch msg.String() {
			case "up":
				if m.userIndex > 0 {
					m.userIndex--
				}
			case "down":
				if m.userIndex < len(m.users)-1 {
					m.userIndex++
				}
			case "enter":
				m.userSelected = true
				userName := m.users[m.userIndex].Name
				return m, fetchFollowedFeedsCmd(m.Q, userName)
			}

		case m.userSelected && !m.feedSelected:
			switch msg.String() {
			case "up":
				if m.feedIndex > 0 {
					m.feedIndex--
				}
			case "down":
				if m.feedIndex < len(m.feeds)-1 {
					m.feedIndex++
				}
			case "enter":
				m.feedSelected = true
				if !m.feeds[m.feedIndex].FeedID.Valid {
					fmt.Printf("FeedID is NULL in feed follows.")
					return m, nil
				}
				feedID := m.feeds[m.feedIndex].FeedID.UUID
				return m, fetchPostsCmd(m.Q, m.users[m.userIndex].ID, feedID)
			case "esc":
				m.userSelected = false
			}

		case m.userSelected && m.feedSelected && !m.postSelected:
			switch msg.String() {
			case "up":
				if m.postIndex > 0 {
					m.postIndex--
				}
			case "down":
				if m.postIndex < len(m.posts)-1 {
					m.postIndex++
				}
			case "enter":
				m.postSelected = true
				post := m.posts[m.postIndex]
				m.Status = fmt.Sprintf("Selected post: %s", post.Title)
			case "esc":
				m.feedSelected = false
			}

		default:
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}

	case usersFetchedMsg:
		m.users = msg.Users
		m.Loading = false
		return m, nil

	case feedsFetchedMsg:
		m.feeds = msg.Feeds
		m.feedIndex = 0
		m.feedSelected = false
		return m, nil

	case postsFetchedMsg:
		m.posts = msg.Posts
		m.postIndex = 0
		m.postSelected = false
		return m, nil
	}

	return m, nil
}
