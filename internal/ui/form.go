package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type feedFormModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	onSubmit   func(name, url string) tea.Msg
}

func newFeedFormModel(onSubmit func(name, url string) tea.Msg) feedFormModel {
	m := feedFormModel{
		inputs:     make([]textinput.Model, 2),
		focusIndex: 0,
		onSubmit:   onSubmit,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "string"
			t.Prompt = "Feed Name: "
			t.CharLimit = 128
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "string"
			t.Prompt = "Feed RSS URL: "
			t.CharLimit = 256
		}
		m.inputs[i] = t
	}
	return m
}
