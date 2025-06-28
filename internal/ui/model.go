package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

type Model struct {
	Q *database.Queries

	userList    list.Model
	feedList    list.Model
	allFeedList list.Model
	postList    list.Model

	level int // 0=user, 1=feed, 2=post

	// User input model
	inputMode bool
	textInput textinput.Model

	// Feed input form model
	form *feedFormModel

	users    []database.User
	feeds    []database.GetFeedFollowsForUserRow
	allFeeds []database.GetFeedsRow
	posts    []database.Post

	// Opened posts
	opened map[uuid.UUID]bool

	Status  string
	Loading bool
	Err     error
}

func (m *Model) Init() tea.Cmd {
	keys := newKeyMap()

	// User Model
	m.userList = list.New(nil, itemDelegate{}, defaultWidth, listHeight)
	m.userList.Title = "Welcome to gator, please select a user:"
	m.userList.Styles.Title = titleStyle
	m.userList.SetShowStatusBar(false)
	m.userList.Styles.PaginationStyle = paginationStyle
	m.userList.Styles.HelpStyle = helpStyle
	m.userList.DisableQuitKeybindings()
	// New User input Model
	ti := textinput.New()
	ti.Placeholder = "Enter new username"
	ti.CharLimit = 50
	ti.Width = 30
	ti.Focus()
	m.textInput = ti
	m.inputMode = false
	// Functions
	m.userList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Add, keys.Del}
	}

	// Feed Model
	m.feedList = list.New(nil, itemDelegate{}, defaultWidth, listHeight)
	m.feedList.Title = "Select a feed:"
	m.feedList.Styles.Title = titleStyle
	m.feedList.SetShowStatusBar(false)
	m.feedList.Styles.PaginationStyle = paginationStyle
	m.feedList.Styles.HelpStyle = helpStyle
	m.feedList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Enter, keys.Back, keys.Quit}
	}
	m.feedList.DisableQuitKeybindings()
	// User functions
	m.feedList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Add, keys.Del, keys.Follow}
	}

	// Add feed form model
	if m.level == 99 && m.form != nil {
		cmds := m.form.updateFocus() // focus the first input
		return tea.Batch(fetchUsersCmd(m.Q), tea.Batch(cmds...))
	}
	// All feed model
	m.allFeedList = list.New(nil, itemDelegate{}, defaultWidth, listHeight)
	m.allFeedList.Title = "Select a feed to follow:"
	m.allFeedList.Styles.Title = titleStyle
	m.allFeedList.SetShowStatusBar(false)
	m.allFeedList.Styles.PaginationStyle = paginationStyle
	m.allFeedList.Styles.HelpStyle = helpStyle
	m.allFeedList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Enter, keys.Back, keys.Quit}
	}
	// Post Model
	m.postList = list.New(nil, itemDelegate{}, defaultWidth, listHeight)
	m.postList.Title = "Select to open post in browser:"
	m.postList.Styles.Title = titleStyle
	m.postList.SetShowStatusBar(false)
	m.postList.Styles.PaginationStyle = paginationStyle
	m.postList.Styles.HelpStyle = helpStyle
	m.postList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Enter, keys.Back, keys.Quit}
	}
	m.postList.DisableQuitKeybindings()

	// Opened Posts
	m.opened = make(map[uuid.UUID]bool)

	return fetchUsersCmd(m.Q)
}
