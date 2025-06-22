package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *feedFormModel) Update(msg tea.Msg) (feedFormModel, tea.Cmd, bool) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.focusIndex == len(m.inputs) {
				// Form submitted
				name := strings.TrimSpace(m.inputs[0].Value())
				url := strings.TrimSpace(m.inputs[1].Value())
				cmd := func() tea.Msg {
					return m.onSubmit(name, url)
				}
				return *m, cmd, true
			}

			m.focusIndex++
			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			}
			cmds := m.updateFocus()
			return *m, tea.Batch(cmds...), false

			// Set focus to next input
		case "tab", "shift+tab", "up", "down":
			s := msg.String()

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

		case "esc":
			return *m, tea.Quit, false
		}
	}

	cmd := m.updateInputs(msg)
	return *m, cmd, false
}

func (m *feedFormModel) updateFocus() []tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		if i == m.focusIndex {
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
		} else {
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = noStyle
			m.inputs[i].TextStyle = noStyle
		}
	}
	return cmds
}

func (m *feedFormModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}
