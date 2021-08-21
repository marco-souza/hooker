package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var hooksFolder = ".hooks"

func check(err error) {
	if err != nil {
		panic(err)
	}
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

func makeFormatedError(template string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(template, a...))
}
