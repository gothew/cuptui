package filetree

import "io/fs"

const fileIconWidth = 2

// Item represents a list item.
type Item struct {
	title            string
	desc             string
	fileName         string
	shortName        string
	extension        string
	currentDirectory string
	isDirectory      bool
	showIcons        bool
	fileInfo         fs.FileInfo
}

// Title returns the title of the list item.
func (i Item) Title() string {
	return i.title
}

// FileName returns the file name of the list item.
func (i Item) FileName() string { return i.fileName }

// FileExtension returns the file extension of the list item.
func (i Item) FileExtension() string { return i.extension }

// IsDirectory returns true if the list item is a directory.
func (i Item) IsDirectory() bool { return i.isDirectory }

// Description returns the description of the list item.
func (i Item) Description() string { return i.desc }

// FilterValue returns the current filter value.
func (i Item) FilterValue() string { return i.title }

// ShortName returns the short name of the selected item.
func (i Item) ShortName() string { return i.shortName }

// CurrentDirectory returns the current directory of the tree.
func (i Item) CurrentDirectory() string { return i.currentDirectory }
