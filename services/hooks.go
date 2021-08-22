package hooks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Hooks []string

var hooksFolder = ".hooks"
var availableHooks = Hooks{
	"applypatch-msg",
	"pre-commit",
	"pre-rebase",
	"commit-msg",
	"pre-commit",
	"pre-receive",
	"fsmonitor-watchman",
	"pre-merge-commit",
	"push-to-checkout",
	"post-update",
	"prepare-commit-msg",
	"update",
	"pre-applypatch",
	"pre-push",
}

func Init() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("No git repository found! 😢")
	}

	if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
		fmt.Println("🪝 Creating hooks folder")

		err := os.Mkdir(hooksFolder, 0755)
		check(err)

		if hooks := ListHooks(); len(hooks) > 0 {
			fmt.Println("🔗 Binding hooks ")
			for _, hook := range hooks {
				bindHook(hook)
			}
		}
	}

	fmt.Println("🎉 Your hooker is ready to go!")
}

func AddHook(hook string, cmd string) {
	data := []byte("#! /bin/bash\n" + cmd)
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	err := ioutil.WriteFile(hookFilename, data, 0755)
	check(err)

	bindHook(hook)
	fmt.Printf("- 🎉 All right, `%s` hook is ready to go!\n", hook)
}

func DropHook(hook string) {
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	err := os.Remove(hookFilename)
	check(err)

	hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
	err = os.Remove(hookTarget)
	check(err)

	fmt.Printf("- 🎉 Ok, `%s` hook is no more!\n", hook)
}

func DropAll() {
	err := os.RemoveAll(hooksFolder)
	check(err)

	// remove all hooks symlinks
	for _, hook := range ListHooks() {
		hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
		err = os.Remove(hookTarget)
		check(err)
	}

	fmt.Println("🎉 All right, no hookers here!")
}

func ListHooks() []string {
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

func CheckIsValidHook(hook string) error {
	if !isValidHook(hook) {
		return fmt.Errorf("Oops, `%s` is not a git-hook, try: %s", hook, availableHooks)
	}
	return nil
}

func CheckHasHookerInitialized() error {
	if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
		return fmt.Errorf("Please, initialize your project with `hooker init`.")
	}
	return nil
}

func HasHook(hook string) bool {
	hookFilename := fmt.Sprintf("%s/%s", hooksFolder, hook)
	if _, err := os.Stat(hookFilename); os.IsNotExist(err) {
		return false
	}
	return true
}

func isValidHook(hook string) bool {
	for _, v := range availableHooks {
		if v == hook {
			return true
		}
	}

	return false
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
