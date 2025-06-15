package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if !m.userSelected {
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
				userID := m.users[m.userIndex].ID
				return m, fetchPostsCmd(m.Q, userID)
			}
		} else {
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}

	case usersFetchedMsg:
		m.users = msg.Users
		m.Loading = false
		return m, nil
	}
	return m, nil
}
