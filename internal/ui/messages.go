package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzagalex/gator/internal/database"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		var cmd tea.Cmd

		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if m.inputMode {
			m.textInput, cmd = m.textInput.Update(msg)

			switch msg.String() {
			case "enter":
				inputValue := m.textInput.Value()
				m.inputMode = false
				m.textInput.Reset()
				return m, createUsersCmd(m.Q, inputValue)
			case "esc":
				m.inputMode = false
				m.textInput.Reset()
				return m, nil
			}
			return m, cmd
		}

		switch m.level {
		case 0: // user list active
			m.userList, cmd = m.userList.Update(msg)
			switch msg.String() {
			case "enter":
				u := m.userList.SelectedItem().(listItem).meta.(database.User)
				return m, fetchFollowedFeedsCmd(m.Q, u.Name)
			case "+":
				m.inputMode = true
				m.textInput.Focus()
				return m, nil
			}

			return m, cmd

		case 1: // feed list active
			m.feedList, cmd = m.feedList.Update(msg)
			switch msg.String() {
			case "enter":
				f := m.feedList.SelectedItem().(listItem).meta.(database.GetFeedFollowsForUserRow)
				if f.FeedID.Valid {
					return m, fetchPostsCmd(m.Q, f.UserID.UUID, f.FeedID.UUID)
				}
			case "esc":
				m.level = 0 // go back to user list
			}
			return m, cmd

		case 2: // post list active
			m.postList, cmd = m.postList.Update(msg)
			switch msg.String() {
			case "enter":
				p := m.postList.SelectedItem().(listItem).meta.(database.Post)
				// Extract user ID from the currently selected user
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				userID := selectedUser.ID
				_ = openBrowser(p.Url) // ignore error for now
				return m, postOpenedPostCmd(m.Q, userID, p.FeedID, p.ID)

			case "esc":
				m.level = 1
			}
			return m, cmd
		}

	case usersFetchedMsg:
		items := make([]list.Item, len(msg.Users))
		for i, u := range msg.Users {
			items[i] = listItem{title: u.Name, meta: u}
		}
		m.userList.SetItems(items)
		m.users = msg.Users
		m.Loading = false
		m.level = 0
		return m, nil

	case feedsFetchedMsg:
		items := make([]list.Item, len(msg.Feeds))
		for i, f := range msg.Feeds {
			items[i] = listItem{title: f.FeedName.String, meta: f}
		}
		m.feedList.SetItems(items)
		m.feeds = msg.Feeds
		m.level = 1
		return m, nil

	case postsFetchedMsg:
		items := make([]list.Item, len(msg.Posts))
		for i, p := range msg.Posts {
			items[i] = listItem{title: p.Title, meta: p}
		}
		m.postList.SetItems(items)
		m.posts = msg.Posts
		m.level = 2
		return m, nil

	case CreateUserMsg:
		if msg.Error != nil {
			m.Err = msg.Error
		} else {
			return m, fetchUsersCmd(m.Q)
		}
	}
	return m, nil
}
