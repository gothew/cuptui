package filetree

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gothew/cuptui/dirfs"
)

func (b Bubble) Init() tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

  if b.startDir == "" {
    cmd = getDirectoryListingCmd(dirfs.CurrentDirectory, b.showHidden, b.showIcons)
  } else {
    cmd = getDirectoryListingCmd(b.startDir, b.showHidden, b.showIcons)
  }

	cmds = append(cmds, cmd, textinput.Blink)

	return tea.Batch(cmds...)
}
