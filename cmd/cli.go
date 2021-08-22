package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func makeHookerCli() *cobra.Command {
	cli := &cobra.Command{
		Use:   "hooker",
		Short: "A git hook manager",
		Long:  `Hooker is a CLI tool for managing git hooks.`,
	}

	cli.AddCommand(addCmd)
	cli.AddCommand(initCmd)
	cli.AddCommand(dropCmd)
	cli.AddCommand(listCmd)

	return cli
}

func Execute() {
	cli := makeHookerCli()
	if err := cli.Execute(); err != nil {
		log.Fatal("Error execution your cli app", err)
	}
}
