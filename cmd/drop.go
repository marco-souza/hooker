package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop [hook]",
	Short: "Drop your hooks",
	Run: func(cmd *cobra.Command, args []string) {
		hook := args[0]
		dropHook(hook)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return makeFormatedError("Please specify a hook")
		}

		hook := args[0]
		if !availableHooks.Contains(hook) {
			return makeFormatedError("Oops, `%s` is not a git-hook, try: %s", args[0], availableHooks)
		}

		if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
			return makeFormatedError("Please, initialize your project with `hooker init`.")
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
