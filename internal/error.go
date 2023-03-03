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
	ErrBadConfigMerge           = errors.New("Couldn't merges configs..")

	ErrBadURLHTTPResposne = errors.New("URL response was not OK (200)..")
	ErrBadURLConfig       = errors.New("URL config couldn't be read..")
	ErrBadURLRequest      = errors.New("URL couldn't be reached..")
	ErrBadURLResponse     = errors.New("URL response couldn't be read..")

	ErrNoConfigDefined   = errors.New("No config was found..")
	ErrNoTemplateDefined = errors.New("No template was found..")
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
