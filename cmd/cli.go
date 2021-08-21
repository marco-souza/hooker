package cmd

import (
	"github.com/spf13/cobra"
	"log"
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

	return cli
}

func Execute() {
	cli := makeHookerCli()
	if err := cli.Execute(); err != nil {
		log.Fatal("Error execution your cli app", err)
	}
}
