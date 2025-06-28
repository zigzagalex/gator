package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const defaultWidth = 50
const listHeight = 14

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Back   key.Binding
	Add    key.Binding
	Del    key.Binding
	Follow key.Binding
	Quit   key.Binding
	Help   key.Binding
}

func newKeyMap() keyMap {
    km := keyMap{
        Up:     key.NewBinding(key.WithKeys("up", "k"),   key.WithHelp("↑/k", "up")),
        Down:   key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
        Enter:  key.NewBinding(key.WithKeys("enter"),     key.WithHelp("enter", "select")),
        Back:   key.NewBinding(key.WithKeys("esc"),       key.WithHelp("esc", "back")),
        Add:    key.NewBinding(key.WithKeys("+"),         key.WithHelp("+", "add")),
        Del:    key.NewBinding(key.WithKeys("-"),         key.WithHelp("-", "delete")),
        Follow: key.NewBinding(key.WithKeys("="),         key.WithHelp("=", "follow")),
        Quit:   key.NewBinding(key.WithKeys("q"),         key.WithHelp("q", "quit")),
        Help:   key.NewBinding(key.WithKeys("?"),         key.WithHelp("?", "help")),
    }
    return km
}
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	cursorStyleList = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
	baseItemStyle   = lipgloss.NewStyle().PaddingLeft(4)

	unreadItemStyle   = baseItemStyle
	readItemStyle     = baseItemStyle.Foreground(lipgloss.Color("240"))
	selectedItemStyle = baseItemStyle.Bold(true)
)

type listItem struct {
	title  string
	meta   any
	opened bool
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return "" }
func (i listItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, itm list.Item) {
	li, _ := itm.(listItem)

	prefix := "  "
	style := unreadItemStyle

	if li.opened {
		style = readItemStyle
	}
	if index == m.Index() {
		prefix = "> "
		style = selectedItemStyle
	}

	fmt.Fprint(w, style.Render(prefix+li.title))
}
