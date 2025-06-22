package ui

import "strings"

func (m feedFormModel) View() string {
	var b strings.Builder
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}
	button := blurredButton
	if m.focusIndex == len(m.inputs) {
		button = focusedButton
	}
	b.WriteString("\n\n" + button + "\n")
	return b.String()
}
