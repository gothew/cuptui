package filetree

import (
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) View() string {
	var inputView string

	switch b.state {
	case idleState:
		inputView = ""
	case createFileState, createDirectoryState, renameItemState:
		inputView = b.input.View()
	default:
		inputView = ""
	}

	return bubbleStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, b.list.View(), inputStyle.Render(inputView)),
	)
}
