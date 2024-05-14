package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "install"},
	Short:   "Initialize hooker on local repo",
	Run: func(cmd *cobra.Command, args []string) {
		h.Init()
	},
}
