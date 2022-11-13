package filetree

import "github.com/charmbracelet/bubbles/key"

var (
  openDirectoryKey = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "open directory"))
	createDirectoryKey = key.NewBinding(key.WithKeys("N"), key.WithHelp("N", "create directory"))
	createFileKey      = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "create file"))
	renameItemKey      = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "rename item"))
	submitInputKey     = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit input value"))
	escapeKey          = key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit"))
)
