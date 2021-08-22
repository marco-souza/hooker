package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "Initialize hooker on local repo",
	Run:     listHandler,
	Args:    argsValidation,
}

func listHandler(cmd *cobra.Command, args []string) {
	hooks := listHooks()
	if len(hooks) == 0 {
		fmt.Println("😖 Sorry, no hook found")
	}
	for _, hook := range hooks {
		hookPath := fmt.Sprintf(".hooks/%s", hook)
		command, err := ioutil.ReadFile(hookPath)
		check(err)
		fmt.Printf("🪝 %s\n===\n%s\n\n", hook, string(command))
	}
}

func argsValidation(cmd *cobra.Command, args []string) error {
	if err := checkHasHooker(); err != nil {
		return err
	}

	return nil
}

func listHooks() []string {
	hooksNameList := []string{}
	files, err := ioutil.ReadDir(hooksFolder)
	if err != nil || len(files) == 0 {
		return hooksNameList
	}
	for _, file := range files {
		hooksNameList = append(hooksNameList, file.Name())
	}
	return hooksNameList
}
