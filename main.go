package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/develeap/dustomize/file"
	"github.com/develeap/dustomize/internal"
)

var (
	// Flags
	fileFlag         string
	folderFlag       string
	configFlag       string
	verboseFlag      bool
	exportFlag       bool
	outputFolderFlag string

	// Files slice
	files []file.File

	// Config file (YAML) reader
	m map[interface{}]interface{}
)

func main() {
	flag.StringVar(&configFlag, "config", "", "Config definition")
	flag.StringVar(&folderFlag, "folder", "", "Folder to parse")
	flag.BoolVar(&verboseFlag, "verbose", false, "Enable to view logs")
	flag.BoolVar(&exportFlag, "export", false, "Enable to export the resulted files")
	flag.StringVar(&outputFolderFlag, "output", "", "Output folder")
	flag.Parse()

	if verboseFlag {
		fmt.Println("Loading folder..")
	}

	if folderFlag != "" {
		err := file.FolderRead(folderFlag, &files, verboseFlag)
		if err != nil {
			panic(err)
		}
	}

	if configFlag != "" {
		var err error = nil
		m, err = internal.ReadConfigFromFile(configFlag)
		if err != nil {
			panic(err)
		}
	}

	if verboseFlag {
		fmt.Println("\nParsing files..")
	}

	for _, f := range files {
		res, err := internal.ParseFile(f.Path, f.Content, m)
		if err != nil {
			panic(err)
		}

		if verboseFlag {
			fmt.Printf("\n%s:\n", f.Name)
			fmt.Println(res)
		}

		if exportFlag {
			if outputFolderFlag != "" {
				err := ioutil.WriteFile((outputFolderFlag + "/" + f.Name), []byte(res), 0644)

				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Please also configure `output` flag to export the results")
			}
		}
	}
}
