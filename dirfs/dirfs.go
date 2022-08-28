package dirfs

import (
	"errors"
	"os"
)

// RenameDirectoryItem renames a directory or files given a source and
// destination.
func RenameDirectoryItem(src, dst string) error {
  err := os.Rename(src, dst)
  return err
}

// CreateDirectory new directories given a name
func CreateDirectory(name string) error {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
