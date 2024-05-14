package cmd

import (
	"bufio"
	"fmt"
	"os"

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
		h.DropAll()
		return
	}
	hook := args[0]
	h.DropHook(hook)
}

func validateDropArgs(cmd *cobra.Command, args []string) error {
	if err := h.CheckHasHookerInitialized(); err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Println("This will drop ALL git h, are you sure? (y/N)")

		reader := bufio.NewReader(os.Stdin)
		char, _, _ := reader.ReadLine()

		switch string(char) {
		case "Y", "y", "YES", "yes":
			return nil
		}

		os.Exit(0)
	}

	hook := args[0]
	if err := h.CheckIsValidHook(hook); err != nil {
		return err
	}
	if !h.HasHook(hook) {
		return fmt.Errorf("Hmm, looks like `%s` hook doesn't exists.", hook)
	}

	return nil
}
