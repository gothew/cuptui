package filetree

import "github.com/charmbracelet/bubbles/key"

var (
	createDirectoryKey = key.NewBinding(key.WithKeys("N"), key.WithHelp("N", "create directory"))
	renameItemKey      = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "rename item"))
	submitInputKey     = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit input value"))
	escapeKey          = key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit"))
)
