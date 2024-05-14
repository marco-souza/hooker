package hooks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type hook struct {
	folder         string
	availableHooks hooks
}

type hooks []string

var instance *hook

// New creates a new instance of hook service
func New() *hook {
	if instance == nil {
		instance = &hook{
			folder: "./.hooks",
			availableHooks: hooks{
				"applypatch-msg",
				"pre-commit",
				"pre-rebase",
				"commit-msg",
				"pre-receive",
				"fsmonitor-watchman",
				"pre-merge-commit",
				"push-to-checkout",
				"post-update",
				"prepare-commit-msg",
				"update",
				"pre-applypatch",
				"pre-push",
			},
		}
	}

	return instance
}

func (h *hook) Init() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("No git repository found! 😢")
	}

	if _, err := os.Stat(h.folder); os.IsNotExist(err) {
		fmt.Println("🪝 Creating hooks folder")

		err := os.Mkdir(h.folder, 0755)
		check(err)
	}

	if hooks := h.ListHooks(); len(hooks) > 0 {
		fmt.Println("🔗 Binding hooks ")
		for _, hook := range hooks {
			h.bindHook(hook)
		}
	}

	fmt.Println("🎉 Your hooker is ready to go!")
}

func (h *hook) AddHook(hook string, cmd string) {
	data := []byte("#! /bin/bash\n" + cmd)
	hookFilename := fmt.Sprintf("%s/%s", h.folder, hook)
	err := os.WriteFile(hookFilename, data, 0755)
	check(err)

	h.bindHook(hook)
	fmt.Printf("- 🎉 All right, `%s` hook is ready to go!\n", hook)
}

func (h *hook) DropHook(hook string) {
	hookFilename := fmt.Sprintf("%s/%s", h.folder, hook)
	err := os.Remove(hookFilename)
	check(err)

	hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
	err = os.Remove(hookTarget)
	check(err)

	fmt.Printf("- 🎉 Ok, `%s` hook is no more!\n", hook)
}

func (h *hook) DropAll() {
	err := os.RemoveAll(h.folder)
	check(err)

	// remove all hooks symlinks
	for _, hook := range h.ListHooks() {
		hookTarget := fmt.Sprintf(".git/hooks/%s", hook)
		err = os.Remove(hookTarget)
		check(err)
	}

	fmt.Println("🎉 All right, no hookers here!")
}

func (h *hook) ListHooks() []string {
	hooksNameList := []string{}
	files, err := os.ReadDir(h.folder)
	// read dir error or no files

	if err != nil || len(files) == 0 {
		return hooksNameList
	}
	for _, file := range files {
		hooksNameList = append(hooksNameList, file.Name())
	}
	return hooksNameList
}

func (h *hook) CheckIsValidHook(hook string) error {
	if !h.isValidHook(hook) {
		return fmt.Errorf("Oops, `%s` is not a git-hook, try: %s", hook, h.availableHooks)
	}
	return nil
}

func (h *hook) CheckHasHookerInitialized() error {
	if _, err := os.Stat(h.folder); os.IsNotExist(err) {
		return fmt.Errorf("Please, initialize your project with `hooker init`.")
	}
	return nil
}

func (h *hook) HasHook(hook string) bool {
	hookFilename := fmt.Sprintf("%s/%s", h.folder, hook)
	if _, err := os.Stat(hookFilename); os.IsNotExist(err) {
		return false
	}
	return true
}

func (h *hook) isValidHook(hook string) bool {
	for _, v := range h.availableHooks {
		if v == hook {
			return true
		}
	}

	return false
}

func (h *hook) bindHook(hook string) {
	hookFilename := fmt.Sprintf("%s/%s", h.folder, hook)
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
