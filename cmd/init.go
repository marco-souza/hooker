package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "install"},
	Short:   "Initialize hooker on local repo",
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func initialize() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("No git repository found! 😢")
	}

	if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
		fmt.Println("🪝 Creating hooks folder")

		err := os.Mkdir(hooksFolder, 0755)
		check(err)

		if hooks := listHooks(); len(hooks) > 0 {
			fmt.Println("🔗 Binding hooks ")
			for _, hook := range hooks {
				bindHook(hook)
			}
		}
	}

	fmt.Println("🎉 Your hooker is ready to go!")
}
