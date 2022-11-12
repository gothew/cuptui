package filetree

import "github.com/charmbracelet/lipgloss"

func (b *Bubble) SetSize(width, height int) {
	horizontal, vertical := bubbleStyle.GetFrameSize()

	b.list.Styles.StatusBar.Width(width - horizontal)
	b.list.SetSize(
		width-horizontal-vertical,
		height-vertical-lipgloss.Height(b.input.View())-inputStyle.GetVerticalPadding(),
	)
}

func (b Bubble) GetSelectedItem() Item {
	selectedDir, ok := b.list.SelectedItem().(Item)
	if ok {
		return selectedDir
	}
	return Item{}
}
