package cmd

import (
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
	return fmt.Errorf(template, a...)
}

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

func checkHasHooker() error {
	if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
		return makeFormatedError("Please, initialize your project with `hooker init`.")
	}
	return nil
}

func checkIsValidHook(hook string) error {
	if !availableHooks.Contains(hook) {
		return makeFormatedError("Oops, `%s` is not a git-hook, try: %s", hook, availableHooks)
	}
	return nil
}

func hasHook(hook string) bool {
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	if _, err := os.Stat(hookFilename); os.IsNotExist(err) {
		return false
	}
	return true
}
