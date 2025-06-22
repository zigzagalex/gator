package ui

import "fmt"

func (m Model) View() string {
	if m.Loading {
		return "Loading…"
	}
	if m.inputMode {
		return fmt.Sprintf(
			"\nCreate a new user:\n\n%s\n\n(enter to submit, esc to cancel)",
			m.textInput.View(),
		)
	}
	switch m.level {
	case 0:
		return m.userList.View()
	case 1:
		return m.feedList.View()
	case 2:
		return m.postList.View()
	}
	return "No data"
}
