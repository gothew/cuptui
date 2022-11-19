package filetree

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	thousand    = 1000
	ten         = 10
	fivePercent = 0.0499
)

func ConvertBytesToSizeString(size int64) string {
	if size < thousand {
		return fmt.Sprintf("%dB", size)
	}

	suffix := []string{
		"K", //kilo
		"M", //mega
		"G", //giga
		"T", //teta
		"P", //peat
		"E", //exa
		"Z", //zeta
		"Y", // yotta
	}

	curr := float64(size) / thousand
	for _, s := range suffix {
		if curr < ten {
			return fmt.Sprintf("%.1f%s", curr-fivePercent, s)
		} else if curr < thousand {
			return fmt.Sprintf("%d%s", int(curr), s)
		}
	}

	return ""
}

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
