package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [hook] [command]",
	Aliases: []string{"a"},
	Short:   "Add a new hook command",
	Run: func(cmd *cobra.Command, args []string) {
		hook, commands := args[0], args[1:]
		addHook(hook, strings.Join(commands, " "))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkHasHooker(); err != nil {
			return err
		}

		switch len(args) {
		case 0:
			return makeFormatedError("Please specify a hook")
		case 1:
			return makeFormatedError("Please specify a command to be bound to '%s' hook", args[0])
		}

		hook := args[0]
		if err := checkIsValidHook(hook); err != nil {
			return err
		}

		if hasHook(hook) {
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
