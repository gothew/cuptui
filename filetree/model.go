package filetree

import "github.com/charmbracelet/bubbles/textinput"

type sessionState int

const (
	idleState sessionState = iota
	createDirectoryState
)

// Bubble represents the properties of a filetree
type Bubble struct {
	input textinput.Model
	state sessionState
}

func New() Bubble {
	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.Placeholder = "Enter directory name"
	input.CharLimit = 250
	input.Width = 50

	return Bubble{
		input: input,
		state: idleState,
	}
}
