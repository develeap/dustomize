package main

import (
	"log"

	"github.com/develeap/dustomize/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := cmd.RootCmd
	err := doc.GenMarkdownTree(cmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
