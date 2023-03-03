/*
Copyright © 2023 Erez Mizrahi <erez.mizrahi@develeap.com>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/develeap/dustomize/file"
	"github.com/develeap/dustomize/internal"
	"github.com/spf13/cobra"
)

var (
	files []file.File

	// Config file (YAML) reader
	m map[interface{}]interface{}

	configFlag     string
	filesFlag      []string
	folderFlag     string
	outputPathFlag string
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse your templates",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: parse,
}

func init() {
	RootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVarP(&configFlag, "config", "c", "", "Values file to read from.")
	parseCmd.Flags().StringSliceVarP(&filesFlag, "file", "f", []string{}, "Template files to parse.")
	parseCmd.Flags().StringVarP(&folderFlag, "folder", "k", "", "Your templates directory to parse.")
	parseCmd.Flags().StringVarP(&outputPathFlag, "output", "o", "", "Export parsing to target folder.")
}

func parse(cmd *cobra.Command, args []string) {
	var err error

	if configFlag == "" {
		panic("Values file must be selected..")
	}

	if len(filesFlag) != 0 && folderFlag != "" {
		fmt.Println("-f and -k flags cannot be used together..")
		os.Exit(1)
	}

	if len(filesFlag) == 0 && folderFlag == "" {
		panic("Either one of the flags -f or -k must be set..")
	}

	m, err = internal.ReadConfigFromFile(configFlag)
	if err != nil {
		panic(err)
	}

	if folderFlag != "" {
		err := file.FolderRead(folderFlag, &files, false)
		if err != nil {
			panic(err)
		}
	}

	if len(filesFlag) > 0 {
		for _, v := range filesFlag {
			file.FileRead(v, &files, false)
		}
	}

	for _, f := range files {
		res, err := internal.ParseFile(f.Path, f.Content, m)
		if err != nil {
			panic(err)
		}

		// Do not export, just print
		if outputPathFlag == "" {
			fmt.Println(res)
			return
		}

		err = ioutil.WriteFile((outputPathFlag + "/" + f.Name), []byte(res), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
