package filetree

import "github.com/charmbracelet/bubbles/key"

var (
	openDirectoryKey   = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "open directory"))
	createDirectoryKey = key.NewBinding(key.WithKeys("N"), key.WithHelp("N", "create directory"))
	createFileKey      = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "create file"))
	deleteItemKey      = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete item"))
	renameItemKey      = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "rename item"))
	submitInputKey     = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit input value"))
	moveItemKey        = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "move item"))
	escapeKey          = key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit"))
)
