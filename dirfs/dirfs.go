package dirfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
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

// CopyFile copies a file given a name.
func CopyFile(name string) error {
	var splitName []string
	var output string

	srcFile, err := os.Open(filepath.Clean(name))
	if err != nil {
		return errors.Unwrap(err)
	}

	defer func() {
		err = srcFile.Close()
	}()

	fileExtension := filepath.Ext(name)
	splitFileName := strings.Split(name, "/")
	fileName := splitFileName[len(splitFileName)-1]
	switch {
	case strings.HasPrefix(fileName, ".") && fileExtension != "" && fileExtension == fileName:
		output = fmt.Sprintf("%s_%d", fileName, time.Now().Unix())
	case strings.HasPrefix(fileName, ".") && fileExtension != "" && fileExtension != fileName:
		splitName = strings.Split(fileName, ".")
		output = fmt.Sprintf(".%s_%d.%s", splitName[1], time.Now().Unix(), splitName[2])
	case fileExtension != "":
		splitName = strings.Split(fileName, ".")
		output = fmt.Sprintf("%s_%d.%s", splitName[0], time.Now().Unix(), splitName[1])
	default:
		output = fmt.Sprintf("%s_%d", fileName, time.Now().Unix())
	}

	destFile, err := os.Create(filepath.Clean(output))
	if err != nil {
		return errors.Unwrap(err)
	}

	defer func() {
		err = destFile.Close()
	}()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.Unwrap(err)
	}

	err = destFile.Sync()
	if err != nil {
		return errors.Unwrap(err)
	}

	return errors.Unwrap(err)
}
