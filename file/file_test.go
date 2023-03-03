// Package file provides populating file struct from reading folders/files.
package file

import (
	"io/ioutil"
	"testing"
)

func TestFolderRead(t *testing.T) {
	var filesSlice []File
	testFolder := "../tests/file"
	files, err := ioutil.ReadDir(testFolder)

	if err != nil {
		t.Errorf("folder %s couldn't be read! (%s)", testFolder, err)
	}

	numOfFilesInTestFolder := len(files)

	t.Run("TestFolderRead", func(t *testing.T) {
		err := FolderRead(testFolder, &filesSlice, false)

		if err != nil {
			t.Errorf("folder %s couldn't be read! (%s)", testFolder, err)
		}

		if len(files) != numOfFilesInTestFolder {
			t.Errorf("got %d, want %d", len(files), numOfFilesInTestFolder)
		}
	})
}
