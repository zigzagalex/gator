package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder
	if !m.userSelected {
		b.WriteString("Select a user:\n\n")
		for i, user := range m.users {
			cursor := " "
			if i == m.userIndex {
				cursor = ">"
			}
			fmt.Fprintf(&b, "%s %s\n", cursor, user.Name)
		}
		return b.String()
	}
	return b.String()
}
