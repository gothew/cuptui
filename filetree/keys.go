package filetree

import "github.com/charmbracelet/bubbles/key"

var (
	openDirectoryKey   = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "open directory"))
	createFileKey      = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "create file"))
	submitInputKey     = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit input value"))
	createDirectoryKey = key.NewBinding(key.WithKeys("N"), key.WithHelp("N", "create directory"))
	deleteItemKey      = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete item"))
	copyItemKey        = key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "copy item"))
	renameItemKey      = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "rename item"))
	moveItemKey        = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "move item"))
	escapeKey          = key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit"))
)
