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

	if m.level == 99 && m.form != nil {
		return m.form.View()
	}
	switch {
	case m.level == 0:
		return m.userList.View() + "\n" + m.Status
	case m.level == 1:
		return m.feedList.View() + "\n\n" + m.Status
	case m.level == 2:
		return m.postList.View()
	case m.level == 3:
		return m.allFeedList.View()
	}
	return "No data"
}
