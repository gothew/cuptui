package filetree

func (b Bubble) GetSelectedItem() Item {
	selectedDir, ok := b.list.SelectedItem().(Item)
	if ok {
		return selectedDir
	}
	return Item{}
}
