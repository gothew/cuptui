package dirfs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Directory shortcuts.
const (
	CurrentDirectory  = "."
	PreviousDirectory = ".."
	HomeDirectory     = "~"
	RootDirectory     = "/"
)

// Different types of listings.
const (
	DirectoriesListingType = "directories"
	FilesListingType       = "files"
)

// RenameDirectoryItem renames a directory or files given a source and
// destination.
func RenameDirectoryItem(src, dst string) error {
	err := os.Rename(src, dst)

	return errors.Unwrap(err)
}

// CreateDirectory new directories given a name
func CreateDirectory(name string) error {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			return errors.Unwrap(err)
		}
	}

	return nil
}

// GetDirectoryListing returns a list of files and directories within a given
// directory.
func GetDirectoryListing(dir string, showHidden bool) ([]fs.DirEntry, error) {
	index := 0

	files, err := os.ReadDir(dir)

	if err != nil {
		return nil, errors.Unwrap(err)
	}

	if !showHidden {
		for _, file := range files {
			if !strings.HasPrefix(file.Name(), ".") {
				files[index] = file
				index++
			}
		}

		// Set files to the list that does not include hidden files.
		files = files[:index]
	}

	return files, nil
}

// DeleteDirectory deletes a directory given a name.
func DeleteDirectory(name string) error {
	err := os.RemoveAll(name)

	return errors.Unwrap(err)
}

// GetHomeDirectory returns the users home directory.
func GetHomeDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return home, nil
}

// GetWorkingDirectory returns the current working directory.
func GetWorkingDirectory() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", errors.Unwrap(err)
	}

	return workingDir, nil
}

// DeleteFile deletes a file given a name.
func DeleteFile(name string) error {
	err := os.Remove(name)

	return errors.Unwrap(err)
}

// MoveDirectoryItem moves a file from one place to another.
func MoveDirectoryItem(src, dst string) error {
	err := os.Rename(src, dst)
	return errors.Unwrap(err)
}

// ReadFileContent returns the contents of a file given a name.
func ReadFileContent(name string) (string, error) {
	fileContent, err := os.ReadFile(filepath.Clean(name))

	if err != nil {
		return "", errors.Unwrap(err)
	}

	return string(fileContent), nil
}

// CreateFile creates a file given a name.
func CreateFile(name string) error {
	f, err := os.Create(filepath.Clean(name))
	if err != nil {
		return errors.Unwrap(err)
	}

	if err = f.Close(); err != nil {
		return errors.Unwrap(err)
	}

	return errors.Unwrap(err)
}
