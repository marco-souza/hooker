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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			dropAll()
			return
		}
		hook := args[0]
		dropHook(hook)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkHasHooker(); err != nil {
			return err
		}

		if len(args) == 0 {
			fmt.Println("This will drop ALL git hooks, are you sure? (y/N)")

			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadLine()
			check(err)

			switch string(char) {
			case "Y","y","YES","yes":
				return nil
			}
			return makeFormatedError("Please specify a hook")
		}

		hook := args[0]
		if err := checkIsValidHook(hook); err != nil {
			return err
		}
		if !hasHook(hook) {
			return makeFormatedError("Hmm, looks like `%s` hook doesn't exists.", hook)
		}

		return nil
	},
}

func dropHook(hook string) {
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	err := os.Remove(hookFilename)
	check(err)

	hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
	err = os.Remove(hookTarget)
	check(err)

	fmt.Printf("- 🎉 Ok, `%s` hook is no more!\n", hook)
}

func dropAll() {
	err := os.RemoveAll(hooksFolder)
	check(err)

	// remove all hooks symlinks
	for _, hook := range listHooks() {
		hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
		err = os.Remove(hookTarget)
		check(err)
	}

	fmt.Println("🎉 All right, no hookers here!")
}
