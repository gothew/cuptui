package filetree

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/cuptui/dirfs"
)

type errorMessage error

// createDirectoryCmd creates a directory based on the name provided.
func createDirectoryCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := dirfs.CreateDirectory(name); err != nil {
			return errorMessage(err)
		}
		return nil
	}
}