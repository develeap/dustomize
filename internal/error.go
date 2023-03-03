// Package internal provides internal helper functions.
package internal

import (
	"errors"
	"fmt"
	"os"
)

var (
	exitCode = 1 // Something went wrong

	ErrNoConfigFile             = errors.New("Values file must be selected..")
	ErrBothFileAndFolderDefined = errors.New("-f and -k flags cannot be used together..")
	ErrNoFileOrFolderDefined    = errors.New("Either one of the flags -f or -k must be set..")
	ErrBadConfig                = errors.New("Config couldn't be read..")
	ErrBadFolder                = errors.New("Folder couldn't be read..")
	ErrBadFile                  = errors.New("File couldn't be read..")
	ErrBadExport                = errors.New("File couldn't be exported..")
)

func Stop(errMessage error) {
	fmt.Println(errMessage)
	os.Exit(exitCode)
}

func StopWithDebug(errMessage error, err error) {
	fmt.Println(errMessage)
	fmt.Println("ERROR:", err)
	os.Exit(exitCode)
}
