package filetree

import (
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
  list list.Model
	input textinput.Model
	state sessionState
}

func New() Bubble {
  listDelegate := list.NewDefaultDelegate()

  listModel := list.New([]list.Item{}, listDelegate, 0,0)
  listModel.Title = "Filetree"
  

	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.Placeholder = "Enter directory name"
	input.CharLimit = 250
	input.Width = 50

	return Bubble{
    list: listModel,
		input: input,
		state: idleState,
	}
}
