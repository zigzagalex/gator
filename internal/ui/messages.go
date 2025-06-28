package ui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzagalex/gator/internal/database"
)

// Update function that changes the state of the model depending on the recieved messages
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.level == 99 && m.form != nil {
		form, cmd, done := m.form.Update(msg)
		m.form = &form
		if done {
			m.form = nil
			m.level = 1
			return m, cmd
		}
		return m, cmd
	}

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
			prevFiltering := m.userList.FilterState() == list.Filtering
			m.userList, cmd = m.userList.Update(msg)
			nowFiltering := m.userList.FilterState() == list.Filtering

			// Still filtering: block everything
			if nowFiltering {
				return m, cmd
			}

			// Just finished filtering with a key that wasn't enter → early return
			if prevFiltering && !nowFiltering && msg.String() != "enter" {
				return m, cmd
			}

			switch msg.String() {
			case "enter":
				selectedItem, ok := getFilteredSelectedItem(m.userList)
				if !ok {
					return m, nil
				}
				u := selectedItem.(listItem).meta.(database.User)
				return m, fetchFollowedFeedsCmd(m.Q, u.Name)

			case "+": // add user
				m.inputMode = true
				m.textInput.Focus()
				return m, nil
			case "-": // delete user
				u := m.userList.SelectedItem().(listItem).meta.(database.User)
				return m, deleteUserCmd(m.Q, u.ID)
			case "esc":
				return m, tea.Quit
			}

			return m, cmd

		case 1: // feed list active
			prevFiltering := m.feedList.FilterState() == list.Filtering
			m.feedList, cmd = m.feedList.Update(msg)
			nowFiltering := m.feedList.FilterState() == list.Filtering

			if nowFiltering {
				return m, cmd
			}
			if prevFiltering && !nowFiltering && msg.String() != "enter" {
				return m, cmd
			}

			switch msg.String() {
			case "enter":
				selectedItem, ok := getFilteredSelectedItem(m.feedList)
				if !ok {
					return m, nil
				}
				f := selectedItem.(listItem).meta.(database.GetFeedFollowsForUserRow)
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				if f.FeedID.Valid {
					cmdPosts := fetchPostsCmd(m.Q, f.UserID.UUID, f.FeedID.UUID)
					cmdOpened := fetchOpenedPostsCmd(m.Q, selectedUser.ID)
					return m, tea.Batch(cmdPosts, cmdOpened)
				}

			case "esc":
				m.level = 0 // go back to user list

			case "+": // add feed by activating form for Feedname and URL
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				f := newFeedFormModel(func(name, url string) tea.Msg {
					return createFeedAndFollowCmd(m.Q, selectedUser.ID, name, url)()
				})
				m.form = &f
				m.level = 99
				cmds := m.form.updateFocus()
				return m, tea.Batch(cmds...)

			case "-": // unfollow the selected feed
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				f := m.feedList.SelectedItem().(listItem).meta.(database.GetFeedFollowsForUserRow)
				if f.FeedID.Valid {
					return m, unfollowFeedCmd(m.Q, selectedUser.ID, f.FeedID.UUID)
				}
			case "=": // follow feeds that already exist on the database
				return m, fetchAllFeedsCmd(m.Q)
			}
			return m, cmd

		case 2: // post list active
			prevFiltering := m.postList.FilterState() == list.Filtering
			m.postList, cmd = m.postList.Update(msg)
			nowFiltering := m.postList.FilterState() == list.Filtering

			if nowFiltering {
				return m, cmd
			}
			if prevFiltering && !nowFiltering && msg.String() != "enter" {
				return m, cmd
			}

			switch msg.String() {
			case "enter":
				selectedItem, ok := getFilteredSelectedItem(m.postList)
				if !ok {
					return m, nil
				}
				p := selectedItem.(listItem).meta.(database.Post)
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				userID := selectedUser.ID
				_ = openBrowser(p.Url)
				cmdLog := postOpenedPostCmd(m.Q, userID, p.FeedID, p.ID)
				cmdRef := fetchOpenedPostsCmd(m.Q, userID)
				return m, tea.Batch(cmdLog, cmdRef)
			case "esc":
				m.level = 1 // return to user level
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				return m, fetchFollowedFeedsCmd(m.Q, selectedUser.Name)
			}
			return m, cmd

		case 3: // posts list active
			m.allFeedList, cmd = m.allFeedList.Update(msg)
			if m.allFeedList.FilterState() == list.Filtering {
				return m, cmd // Enter will be used to apply/cancel filter
			}
			switch msg.String() {
			case "enter":
				f := m.allFeedList.SelectedItem().(listItem).meta.(database.GetFeedsRow)
				selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
				return m, followFeedCmd(m.Q, selectedUser.ID, f.ID)
			case "esc":
				m.level = 0
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
		selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)

		items := make([]list.Item, len(msg.Feeds))
		for i, f := range msg.Feeds {
			title := f.FeedName.String

			if f.FeedID.Valid {
				unreadCount, err := m.Q.GetUnreadCount(context.TODO(), database.GetUnreadCountParams{
					UserID: selectedUser.ID,
					FeedID: f.FeedID.UUID,
				})
				if err == nil && unreadCount > 0 {
					title += fmt.Sprintf(" [%d unread]", unreadCount)
				}
			}

			items[i] = listItem{
				title: title,
				meta:  f,
			}
		}
		m.feedList.SetItems(items)
		m.feeds = msg.Feeds
		m.level = 1
		return m, nil

	case postsFetchedMsg:
		items := make([]list.Item, len(msg.Posts))
		for i, p := range msg.Posts {
			items[i] = listItem{
				title:  p.Title,
				meta:   p,
				opened: m.opened[p.ID],
			}
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

	case CreateFeedAndFollowMsg:
		m.level = 1
		m.form = nil
		if msg.Error != nil {
			m.Status = fmt.Sprintf("Failed to create feed: %v", msg.Error)
			return m, nil
		}
		m.Status = "Feed created successfully!"
		selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
		cmd := fetchFollowedFeedsCmd(m.Q, selectedUser.Name)
		return m, cmd

	case allFeedsFetchedMsg:
		if msg.Error != nil {
			m.Status = fmt.Sprintf("Error: %v", msg.Error)
			return m, nil
		}
		items := make([]list.Item, len(msg.Feeds))
		for i, f := range msg.Feeds {
			items[i] = listItem{title: f.Name, meta: f}
		}
		m.allFeedList.SetItems(items)
		m.allFeeds = msg.Feeds
		m.level = 3
		return m, nil

	case followFeedMsg:
		if msg.Error != nil {
			m.Status = fmt.Sprintf("Failed to follow feed: %v", msg.Error)
			return m, nil
		}
		m.Status = "Feed followed successfully!"
		selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
		m.level = 1
		return m, fetchFollowedFeedsCmd(m.Q, selectedUser.Name)

	case unfollowFeedMsg:
		if msg.Error != nil {
			m.Status = fmt.Sprintf("Failed to unfollow feed: %v", msg.Error)
			return m, nil
		}
		m.Status = "Feed unfollowed."
		selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
		m.level = 1
		return m, fetchFollowedFeedsCmd(m.Q, selectedUser.Name)

	case deleteUserMsg:
		if msg.Error != nil {
			m.Status = fmt.Sprintf("Failed to delete user.")
			return m, nil
		}
		m.Status = "User deleted"
		return m, fetchUsersCmd(m.Q)

	case OpenedPostMsg:
		selectedUser := m.userList.SelectedItem().(listItem).meta.(database.User)
		return m, fetchOpenedPostsCmd(m.Q, selectedUser.ID)

	case openedFetchedMsg:
		if msg.Err != nil {
			m.Status = "Error: " + msg.Err.Error()
			return m, nil
		}
		m.opened = msg.Map

		if m.level == 2 && len(m.posts) > 0 {
			items := make([]list.Item, len(m.posts))
			for i, p := range m.posts {
				items[i] = listItem{
					title:  p.Title,
					meta:   p,
					opened: m.opened[p.ID], // ← refreshed flag
				}
			}
			m.postList.SetItems(items)
		}
		return m, nil
	}
	return m, nil
}
