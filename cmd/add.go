package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Hooks []string

var availableHooks = Hooks{"pre-push", "pre-commit"}

func (hooks Hooks) Contains(str string) bool {

	for _, v := range hooks {
		if v == str {
			return true
		}
	}

	return false
}

func AddHook(hook string, cmd string) {
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
	fmt.Printf("- 🎉 All right, `%s` hook is ready to go!\n", hook)
}
