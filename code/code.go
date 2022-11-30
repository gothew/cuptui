// make library for highlight
package code

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/karchx/cuptui/dirfs"
)

type syntaxMsg string
type errorMessage error

const (
	padding = 1
)

func readFileContentCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		content, err := dirfs.ReadFileContent(filename)
		if err != nil {
			return errorMessage(err)
		}

		return syntaxMsg(content)
	}
}

// Bubble represents the properties of a code bubble.
type Bubble struct {
	Viewport           viewport.Model
	BoderColor         lipgloss.AdaptiveColor
	Boderless          bool
	Active             bool
	Filename           string
	HighlightedContent string
}

// New creates a new instance of code.
func New(active, borderless bool, borderColor lipgloss.AdaptiveColor) Bubble {
	viewPort := viewport.New(0, 0)
	border := lipgloss.NormalBorder()

	if borderless {
		border = lipgloss.HiddenBorder()
	}

	viewPort.Style = lipgloss.NewStyle().
		PaddingLeft(padding).
		PaddingRight(padding).
		Border(border).
		BorderForeground(borderColor)

	return Bubble{
		Viewport:   viewPort,
		Boderless:  borderless,
		Active:     active,
		BoderColor: borderColor,
	}
}

// Init initializes the code bubble.
func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b *Bubble) SetFileName(filename string) tea.Cmd {
	b.Filename = filename

	return readFileContentCmd(filename)
}

func (b *Bubble) SetSize(w, h int) {
	b.Viewport.Width = w - b.Viewport.Style.GetHorizontalFrameSize()
	b.Viewport.Height = h - b.Viewport.Style.GetVerticalFrameSize()

	b.Viewport.SetContent(lipgloss.NewStyle().
		Width(b.Viewport.Width).
		Height(b.Viewport.Height).
    Render(b.HighlightedContent))
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case syntaxMsg:
		b.Filename = ""
		b.HighlightedContent = lipgloss.NewStyle().
			Width(b.Viewport.Width).
			Height(b.Viewport.Height).
			Render(string(msg))

		b.Viewport.SetContent(b.HighlightedContent)

		return b, nil
	case errorMessage:
		b.Filename = ""
		b.HighlightedContent = lipgloss.NewStyle().
			Width(b.Viewport.Width).
			Height(b.Viewport.Height).
			Render("Error: " + msg.Error())

		b.Viewport.SetContent(b.HighlightedContent)

		return b, nil
	}

	if b.Active {
		b.Viewport, cmd = b.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}
	return b, tea.Batch(cmds...)
}

func (b Bubble) View() string {
	border := lipgloss.NormalBorder()

	if b.Boderless {
		border = lipgloss.HiddenBorder()
	}

	b.Viewport.Style = lipgloss.NewStyle().
		PaddingLeft(padding).
		PaddingRight(padding).
		Border(border).
		BorderForeground(b.BoderColor)

	return b.Viewport.View()
}
