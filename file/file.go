// Package file provides populating file struct from reading folders/files.
package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Name    string
	Path    string
	Content string
}

// Read a folder and populate the Files struct
func FolderRead(folder string, files *[]File) error {
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			newFile := File{
				Name:    info.Name(),
				Path:    path,
				Content: string(data),
			}

			*files = append(*files, newFile)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Read a file and populate the Files struct
func FileRead(file string, files *[]File) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	newFile := File{
		Name:    filepath.Base(file),
		Path:    file,
		Content: string(data),
	}

	*files = append(*files, newFile)

	return nil
}
