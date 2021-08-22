package cmd

import (
	"bufio"
	"fmt"
	"os"

	hooks "github.com/marco-souza/hooker/services"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:     "drop <hook>",
	Aliases: []string{"d", "uninstall"},
	Short:   "Drop hook",
	Long:    "Drop the informed hook, or drop hooker if none was passed",
	Args:    validateDropArgs,
	Run:     dropHandler,
}

func dropHandler(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		hooks.DropAll()
		return
	}
	hook := args[0]
	hooks.DropHook(hook)
}

func validateDropArgs(cmd *cobra.Command, args []string) error {
	if err := hooks.CheckHasHookerInitialized(); err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Println("This will drop ALL git hooks, are you sure? (y/N)")

		reader := bufio.NewReader(os.Stdin)
		char, _, _ := reader.ReadLine()

		switch string(char) {
		case "Y", "y", "YES", "yes":
			return nil
		}

		os.Exit(0)
	}

	hook := args[0]
	if err := hooks.CheckIsValidHook(hook); err != nil {
		return err
	}
	if !hooks.HasHook(hook) {
		return fmt.Errorf("Hmm, looks like `%s` hook doesn't exists.", hook)
	}

	return nil
}
