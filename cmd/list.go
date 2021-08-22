package cmd

import (
	"fmt"
	"io/ioutil"

	hooks "github.com/marco-souza/hooker/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "Initialize hooker on local repo",
	Args:    validateListArgs,
	Run:     listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	hooks := hooks.ListHooks()
	if len(hooks) == 0 {
		fmt.Println("😖 Sorry, no hook found")
	}
	for _, hook := range hooks {
		hookPath := fmt.Sprintf(".hooks/%s", hook)
		command, _ := ioutil.ReadFile(hookPath)
		fmt.Printf("🪝 %s\n===\n%s\n\n", hook, string(command))
	}
}

func validateListArgs(cmd *cobra.Command, args []string) error {
	if err := hooks.CheckHasHookerInitialized(); err != nil {
		return err
	}

	return nil
}
