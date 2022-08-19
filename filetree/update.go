package filetree

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, createDirectoryKey):
			if !b.input.Focused() {
				b.input.Focus()
				b.input.Placeholder = "Enter name of new directory"
				b.state = createDirectoryState

				return b, textinput.Blink
			}
		case key.Matches(msg, escapeKey):
			if b.input.Focused() {
				b.input.Reset()
				b.input.Blur()
			}
		case key.Matches(msg, submitInputKey):
			switch b.state {
			case createDirectoryState:
				cmds = append(cmds, tea.Sequentially(
					createDirectoryCmd(b.input.Value()),
				))
			}
			b.state = idleState
			b.input.Blur()
			b.input.Reset()
		}
	}

	switch b.state {
	case createDirectoryState:
		b.input, cmd = b.input.Update(msg)
		cmds = append(cmds, cmd)
	}
	return b, tea.Batch(cmds...)
}
