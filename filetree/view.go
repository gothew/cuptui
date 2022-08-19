package filetree

import (
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) View() string {
	var inputView string

	switch b.state {
	case createDirectoryState:
		inputView = b.input.View()
	default:
		inputView = ""
	}

	return bubbleStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, inputStyle.Render(inputView)),
	)
}
