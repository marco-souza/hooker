package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var hooksFolder = ".hooks"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initializeHooker() {
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

type Hooks []string

var availableHooks = Hooks{"pre-push", "pre-commit"}

func (s Hooks) Contains(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func addHook(hook string, cmd string) {
	if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
		log.Fatal("Please, initialize your project with `hookers init`.")
	}

	if !availableHooks.Contains(hook) {
		log.Fatalf("Oops, `%s` is not a git-hook.", hook)
	}

	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)

	if _, err := os.Stat(hookFilename); !os.IsNotExist(err) {
		log.Fatalf("Seems `%s` hook already exists.", hook)
	}

	data := []byte("#! /bin/bash\n" + cmd)
	err := ioutil.WriteFile(hookFilename, data, 0755)
	check(err)

	bindHook(hook)
	fmt.Printf("🎉 All right, `%s` hook is ready to go\n", hook)
}

func bindHook(hook string) {
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	hookFile, err := filepath.Abs(hookFilename)
	check(err)

	target := fmt.Sprintf(".git/hooks/%s", hook)
	targetFile, err := filepath.Abs(target)
	check(err)

	if err := os.Remove(target); !os.IsNotExist(err) {
		check(err)
	}

	if err := os.Symlink(hookFile, targetFile); err != nil {
		check(err)
	}
}

func main() {
	initializeHooker()
	addHook("pre-commit", "echo Hello from Hooker")
	addHook("pre-push", "echo Hello from Hooker")
}
