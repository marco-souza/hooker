package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop <hook>",
	Short: "Drop hook",
	Long:  "Drop the informed hook, or drop hooker if none was passed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			dropAll()
			return
		}
		hook := args[0]
		dropHook(hook)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
			return makeFormatedError("Please, initialize your project with `hooker init`.")
		}

		switch len(args) {
		case 0:
			fmt.Println("This will drop ALL git hooks, are you sure? (y/N)")

			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadLine()
			check(err)

			switch string(char) {
			case "Y":
			case "y":
			case "YES":
			case "yes":
				return nil
			}
			return makeFormatedError("Please specify a hook")
		}

		hook := args[0]
		if !availableHooks.Contains(hook) {
			return makeFormatedError("Oops, `%s` is not a git-hook, try: %s", args[0], availableHooks)
		}

		hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
		if _, err := os.Stat(hookFilename); os.IsNotExist(err) {
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
	fmt.Println("🎉 All right, no hookers here!")
}
