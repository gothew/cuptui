package filetree

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	idleState sessionState = iota
	createFileState
	createDirectoryState
	deleteItemState
	moveItemState
	renameItemState
)

type itemToMove struct {
	shortName string
	path      string
}

// Bubble represents the properties of a filetree
type Bubble struct {
	state         sessionState
	list          list.Model
	input         textinput.Model
	showHidden    bool
	showIcons     bool
	width         int
	height        int
	startDir      string
	selectionPath string
	itemToMove    itemToMove
	delegate      list.DefaultDelegate
}

// New create a new instance of a filetree.
func New(
	startDir, selectionPath string,
	borderColor, selectedItemColor, titleBackgroundColor, titleForegroundColor lipgloss.AdaptiveColor) Bubble {
	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.Copy().
		Foreground(selectedItemColor).
		BorderLeftForeground(selectedItemColor)
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()

	listModel := list.New([]list.Item{}, listDelegate, 0, 0)
	listModel.Title = "Filetree"
	listModel.Styles.Title = listModel.Styles.Title.Copy().
		Bold(true).
		Italic(true).
		Background(titleBackgroundColor).
		Foreground(titleForegroundColor)
	listModel.DisableQuitKeybindings()
	listModel.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			openDirectoryKey,
			createDirectoryKey,
			createFileKey,
			deleteItemKey,
			escapeKey,
			renameItemKey,
			moveItemKey,
			submitInputKey,
		}
	}
	listModel.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			openDirectoryKey,
			createDirectoryKey,
			createFileKey,
			deleteItemKey,
			escapeKey,
			renameItemKey,
			moveItemKey,
			submitInputKey,
		}
	}

	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.Placeholder = "Enter file name"
	input.CharLimit = 250
	input.Width = 50

	bubbleStyle = bubbleStyle.Copy().BorderForeground(borderColor)

	return Bubble{
		list:          listModel,
		input:         input,
		showHidden:    true,
		showIcons:     false,
		state:         idleState,
		startDir:      startDir,
		selectionPath: selectionPath,
		delegate:      listDelegate,
	}
}
