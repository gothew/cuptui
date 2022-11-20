package filetree

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/cuptui/dirfs"
)

type getDirectoryListingMsg []list.Item
type errorMessage error

// getDirectoryListingCmd updates the directory listing based on the name of
// the directory provided.
func getDirectoryListingCmd(directoryName string, showHidden, showIcons bool) tea.Cmd {
	return func() tea.Msg {
		var err error
		var items []list.Item

		if directoryName == dirfs.HomeDirectory {
			directoryName, err = dirfs.GetHomeDirectory()
			if err != nil {
				return errorMessage(err)
			}
		}

		directoryInfo, err := os.Stat(directoryName)
		if err != nil {
			return errorMessage(err)
		}

		if !directoryInfo.IsDir() {
			return nil
		}

		files, err := dirfs.GetDirectoryListing(directoryName, showHidden)
		if err != nil {
			return errorMessage(err)
		}

		err = os.Chdir(directoryName)
		if err != nil {
			return errorMessage(err)
		}

		workingDirectory, err := dirfs.GetWorkingDirectory()
		if err != nil {
			return errorMessage(err)
		}

		items = append(items, Item{
			title:            dirfs.PreviousDirectory,
			desc:             "",
			shortName:        dirfs.PreviousDirectory,
			fileName:         filepath.Join(workingDirectory, dirfs.PreviousDirectory),
			extension:        "",
			isDirectory:      directoryInfo.IsDir(),
			currentDirectory: workingDirectory,
			fileInfo:         nil,
			showIcons:        false,
		})

		for _, file := range files {
			fileInfo, err := file.Info()
			if err != nil {
				continue
			}

			status := fmt.Sprintf("%s %s %s",
				fileInfo.ModTime().Format("2006-01-02 15:04:05"),
				fileInfo.Mode().String(),
				ConvertBytesToSizeString(fileInfo.Size()))

			items = append(items, Item{
				title:            file.Name(),
				desc:             status,
				shortName:        file.Name(),
				fileName:         filepath.Join(workingDirectory, file.Name()),
				extension:        filepath.Ext(fileInfo.Name()),
				isDirectory:      fileInfo.IsDir(),
				currentDirectory: workingDirectory,
				fileInfo:         fileInfo,
				showIcons:        showIcons,
			})
		}
		return getDirectoryListingMsg(items)
	}
}

func createFileCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := dirfs.CreateFile(name); err != nil {
			return errorMessage(err)
		}

		return nil
	}
}

// createDirectoryCmd creates a directory based on the name provided.
func createDirectoryCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := dirfs.CreateDirectory(name); err != nil {
			return errorMessage(err)
		}
		return nil
	}
}

// deleteItemCmd deletes a directory based on the name provided.
func deleteItemCmd(name string) tea.Cmd {
	return func() tea.Msg {
		fileInfo, err := os.Lstat(name)
		if err != nil {
			return errorMessage(err)
		}

		if fileInfo.IsDir() {
			if err := dirfs.DeleteDirectory(name); err != nil {
				return errorMessage(err)
			}
		} else {
			if err := dirfs.DeleteFile(name); err != nil {
				return errorMessage(err)
			}
		}

		return nil
	}
}

// renameItemCmd renames a file or directory based on the name and value
// provided.
func renameItemCmd(name, value string) tea.Cmd {
	return func() tea.Msg {
		if err := dirfs.RenameDirectoryItem(name, value); err != nil {
			return errorMessage(err)
		}
		return nil
	}
}
