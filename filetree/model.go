package filetree

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type sessionState int

const (
	idleState sessionState = iota
	createDirectoryState
  renameItemState
)

// Bubble represents the properties of a filetree
type Bubble struct {
	state sessionState
  list list.Model
	input textinput.Model
}

// New create a new instance of a filetree.
func New() Bubble {
  listDelegate := list.NewDefaultDelegate()

  listModel := list.New([]list.Item{}, listDelegate, 0,0)
  listModel.Title = "Filetree"
  listModel.DisableQuitKeybindings()
  listModel.AdditionalShortHelpKeys = func() []key.Binding {
    return []key.Binding {
      createDirectoryKey,
      escapeKey,
      renameItemKey,
      submitInputKey,
    }
  }
  listModel.AdditionalFullHelpKeys = func() []key.Binding {
    return []key.Binding {
      createDirectoryKey,
      escapeKey,
      renameItemKey,
      submitInputKey,
    }
  }

	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.Placeholder = "Enter file name"
	input.CharLimit = 250
	input.Width = 50

	return Bubble{
    list: listModel,
		input: input,
		state: idleState,
	}
}
