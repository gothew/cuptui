package filetree

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Init() tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmds = append(cmds, cmd, textinput.Blink)

	return tea.Batch(cmds...)
}
