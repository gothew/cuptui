package filetree

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/cuptui/dirfs"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    b.width = msg.Width
    b.height = msg.Height
	case getDirectoryListingMsg:
		if msg != nil {
			cmd = b.list.SetItems(msg)
			cmds = append(cmds, cmd)
		}

	case tea.KeyMsg:

		switch {
		case key.Matches(msg, createDirectoryKey):
			if !b.input.Focused() {
				b.input.Focus()
				b.input.Placeholder = "Enter name of new directory"
				b.state = createDirectoryState

				return b, textinput.Blink
			}
		case key.Matches(msg, renameItemKey):
			if !b.input.Focused() {
				b.input.Focus()
				b.input.Placeholder = "Enter new name"
				b.state = renameItemState

				return b, textinput.Blink
			}
		case key.Matches(msg, escapeKey):
			b.state = idleState

			if b.input.Focused() {
				b.input.Reset()
				b.input.Blur()
			}
		case key.Matches(msg, submitInputKey):
			selectedItem := b.GetSelectedItem()

			switch b.state {
			case idleState:
				return b, nil
			case createDirectoryState:
				statusCmd := b.list.NewStatusMessage(
					statusMessageInfoStyle("Successfully created directory"),
				)

				cmds = append(cmds, statusCmd, tea.Sequentially(
					createDirectoryCmd(b.input.Value()),
					getDirectoryListingCmd(dirfs.CurrentDirectory, b.showHidden, b.showIcons),
				))
			case renameItemState:
				statusCmd := b.list.NewStatusMessage(
					statusMessageInfoStyle("Successfully renamed"),
				)

				cmds = append(cmds, statusCmd, tea.Sequentially(
					renameItemCmd(selectedItem.fileName, b.input.Value()),
					getDirectoryListingCmd(dirfs.CurrentDirectory, b.showHidden, b.showIcons),
				))
			}

			b.state = idleState
			b.input.Blur()
			b.input.Reset()
		}
	}

	switch b.state {
	case idleState:
		b.list, cmd = b.list.Update(msg)
	case createDirectoryState, renameItemState:
		b.input, cmd = b.input.Update(msg)
		cmds = append(cmds, cmd)
	}
	return b, tea.Batch(cmds...)
}
