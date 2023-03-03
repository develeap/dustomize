/*
Copyright © 2023 Erez Mizrahi <erez.mizrahi@develeap.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/imdario/mergo"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/develeap/dustomize/config"
	"github.com/develeap/dustomize/file"
	"github.com/develeap/dustomize/internal"
)

var (
	readWritePermission fs.FileMode = 0644
	files               []file.File

	cfg config.Config

	// Config file (YAML) reader
	appendedValues map[interface{}]interface{}

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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := config.Read(&cfg)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
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

	if len(cfg.Import.FromFile) == 0 && len(cfg.Import.FromURL) == 0 && cfg.Import.FromText != "" {
		internal.Stop(internal.ErrNoConfigDefined)
	}

	if len(cfg.Export) == 0 {
		internal.Stop(internal.ErrNoTemplateDefined)
	}

	for _, configFile := range cfg.Import.FromFile {
		m, err := internal.ReadConfigFromFile(configFile)

		if err != nil {
			internal.Stop(internal.ErrBadConfig)
		}

		if err := mergo.Merge(&appendedValues, m, mergo.WithOverride); err != nil {
			internal.Stop(internal.ErrBadConfigMerge)
		}
	}

	if len(cfg.Import.FromText) != 0 {
		m, err := internal.ReadConfigFromText(cfg.Import.FromText)

		if err != nil {
			internal.Stop(internal.ErrBadConfig)
		}

		if err := mergo.Merge(&appendedValues, m, mergo.WithOverride); err != nil {
			internal.Stop(internal.ErrBadConfigMerge)
		}
	}

	for _, url := range cfg.Import.FromURL {
		m, err := internal.ReadConfigFromUrl(url)

		if err != nil {
			internal.Stop(internal.ErrBadURLConfig)
		}

		if err := mergo.Merge(&appendedValues, m, mergo.WithOverride); err != nil {
			internal.Stop(internal.ErrBadConfigMerge)
		}
	}

	for _, targetFile := range cfg.Export {
		file.FileRead(targetFile.Template, &files)
	}

	for _, f := range files {
		res, err := internal.ParseFile(f.Path, f.Content, appendedValues)
		if err != nil {
			internal.StopWithDebug(internal.ErrBadFile, err)
		}

		// TODO: find a faster & native approach
		_, index, ok := lo.FindIndexOf(cfg.Export, func(i config.ExportItem) bool {
			return i.Template == f.Path
		})

		if ok {
			err = ioutil.WriteFile(cfg.Export[index].Target, []byte(res), readWritePermission)
			if err != nil {
				internal.StopWithDebug(internal.ErrBadExport, err)
			}
		}
	}

	if cfg.Options.DisplayValues {
		d, err := yaml.Marshal(&appendedValues)
		if err != nil {
			panic("General error")
		}
		fmt.Println("Values:")
		fmt.Printf("\n%s", d)
	}

	// if cfg != config.Config {

	// }

	// if configFlag == "" {
	// 	internal.Stop(internal.ErrNoConfigFile)
	// }

	// if len(filesFlag) != 0 && folderFlag != "" {
	// 	internal.Stop(internal.ErrBothFileAndFolderDefined)
	// }

	// if len(filesFlag) == 0 && folderFlag == "" {
	// 	internal.Stop(internal.ErrNoFileOrFolderDefined)
	// }

	// m, err = internal.ReadConfigFromFile(configFlag)
	// if err != nil {
	// 	internal.StopWithDebug(internal.ErrBadConfig, err)
	// }

	// if folderFlag != "" {
	// 	err := file.FolderRead(folderFlag, &files)
	// 	if err != nil {
	// 		internal.StopWithDebug(internal.ErrBadFolder, err)
	// 	}
	// }

	// if len(filesFlag) > 0 {
	// 	for _, v := range filesFlag {
	// 		file.FileRead(v, &files)
	// 	}
	// }

	// for _, f := range files {
	// 	res, err := internal.ParseFile(f.Path, f.Content, m)
	// 	if err != nil {
	// 		internal.StopWithDebug(internal.ErrBadFile, err)
	// 	}

	// 	// Do not export, just print
	// 	if outputPathFlag == "" {
	// 		fmt.Println(res)
	// 		return
	// 	}

	// 	err = ioutil.WriteFile((outputPathFlag + "/" + f.Name), []byte(res), readWritePermission)
	// 	if err != nil {
	// 		internal.StopWithDebug(internal.ErrBadExport, err)
	// 	}
	// }
}
