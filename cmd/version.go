/*
Copyright © 2023 Erez Mizrahi <erez.mizrahi@develeap.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
// TODO: implement
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of dustomize",
	Long:  `All software has versions. This is dustomize's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dustomize -- HEAD")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
