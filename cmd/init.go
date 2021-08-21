package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Init() {
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
