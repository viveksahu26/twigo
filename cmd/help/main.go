package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/viveksahu26/twigo/cmd/twigo/cli"
	"github.com/viveksahu26/twigo/cmd/twigo/cli/templates"
	errors "github.com/viveksahu26/twigo/cmd/twigo/errors"
)

func main() {
	var dir string
	root := &cobra.Command{
		Use:          "gendoc",
		Short:        "Generate twigo's help docs",
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		RunE: func(*cobra.Command, []string) error {
			err := errors.GenerateExitCodeDocs(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return doc.GenMarkdownTree(cli.New(), dir)
		},
	}
	root.Flags().StringVarP(&dir, "dir", "d", "doc", "Path to directory in which to generate docs")

	templates.SetCustomUsageFunc(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
