package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [hook] [command]",
	Aliases: []string{"a"},
	Short:   "Add a new hook command",
	Args:    valdiateAddArgs,
	Run:     addHandler,
}

func addHandler(cmd *cobra.Command, args []string) {
	hook, commands := args[0], args[1:]
	h.AddHook(hook, strings.Join(commands, " "))
}

func valdiateAddArgs(cmd *cobra.Command, args []string) error {
	if err := h.CheckHasHookerInitialized(); err != nil {
		return err
	}

	switch len(args) {
	case 0:
		return fmt.Errorf("Please specify a hook")
	case 1:
		return fmt.Errorf("Please specify a command to be bound to '%s' hook", args[0])
	}

	hook := args[0]
	if err := h.CheckIsValidHook(hook); err != nil {
		return err
	}

	if h.HasHook(hook) {
		fmt.Printf("Hmm, looks like `%s` hook already exists. Do you wanna replace it? (Y/n) ", hook)

		reader := bufio.NewReader(os.Stdin)
		char, _, _ := reader.ReadLine()

		switch string(char) {
		case "", "Y", "y", "YES", "yes":
			return nil
		}

		os.Exit(0)
	}

	return nil
}
