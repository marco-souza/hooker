package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func makeHookerCli() *cobra.Command {
	cli := &cobra.Command{
		Use:   "hooker",
		Short: "A git hook manager",
		Long:  `Hooker is a CLI tool for managing git hooks.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hooker - git hook manaker to go", args)
		},
	}
	return cli
}

func Execute() {
	cli := makeHookerCli()
	if err := cli.Execute(); err != nil {
		log.Fatal("Error execution your cli app", err)
	}
}
