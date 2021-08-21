package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [hook] [command]",
	Short: "Add a new hook command",
	Run: func(cmd *cobra.Command, args []string) {
		hook, commands := args[0], args[1:]
		addHook(hook, strings.Join(commands, " "))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return makeFormatedError("Please specify a hook")
		case 1:
			return makeFormatedError("Please specify a command to be bound to '%s' hook", args[0])
		}

		hook := args[0]
		if !availableHooks.Contains(hook) {
			return makeFormatedError("Oops, `%s` is not a git-hook, try: %s", args[0], availableHooks)
		}

		if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
			return makeFormatedError("Please, initialize your project with `hooker init`.")
		}

		hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
		if _, err := os.Stat(hookFilename); !os.IsNotExist(err) {
			return makeFormatedError("Hmm, looks like `%s` hook already exists.", hook)
		}

		return nil
	},
}

func addHook(hook string, cmd string) {
	data := []byte("#! /bin/bash\n" + cmd)
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	err := ioutil.WriteFile(hookFilename, data, 0755)
	check(err)

	bindHook(hook)
	fmt.Printf("- 🎉 All right, `%s` hook is ready to go!\n", hook)
}
