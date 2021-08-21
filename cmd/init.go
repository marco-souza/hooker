package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hooker on local repo",
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
	}

	files, err := ioutil.ReadDir(hooksFolder)
	if err != nil || len(files) == 0 {
		return
	}

	fmt.Println("🔗 Binding hooks ")
	for _, file := range files {
		bindHook(file.Name())
	}
}
