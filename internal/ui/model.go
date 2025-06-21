package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzagalex/gator/internal/database"
)

type Model struct {
	Q *database.Queries

	users        []database.User
	userIndex    int
	userSelected bool

	feeds        []database.GetFeedFollowsForUserRow
	feedIndex    int
	feedSelected bool

	posts        []database.Post
	postIndex    int
	postSelected bool

	Status  string
	Loading bool
	Err     error
}

func (m Model) Init() tea.Cmd {
	return fetchUsersCmd(m.Q)
}
