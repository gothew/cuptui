package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/cuptui/filetree"
)

type Bubble struct {
	filetree filetree.Bubble
}

func New() Bubble {
	filetreeModel := filetree.New()

	return Bubble{
		filetree: filetreeModel,
	}
}

func (b Bubble) Init() tea.Cmd {
	return b.filetree.Init()
}

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
    case "q":
      cmds = append(cmds, tea.Quit)
		}
	}

	b.filetree, cmd = b.filetree.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}

func (b Bubble) View() string {
	return b.filetree.View()
}

func main() {
	b := New()
	p := tea.NewProgram(b, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
