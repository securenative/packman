package business

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type GitProjectInit struct{}

func NewGitProjectInit() *GitProjectInit {
	return &GitProjectInit{}
}

func (this *GitProjectInit) Init(destPath string) error {
	path := packmanPath(destPath)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	scriptPath := filepath.Join(path, "main.go")
	if err := this.write(scriptPath, replyScript); err != nil {
		return err
	}

	modPath := filepath.Join(path, "go.mod")
	if err := this.write(modPath, modeFile); err != nil {
		return err
	}

	gitInit := exec.Command("git", "init")
	gitInit.Dir = destPath
	if err := gitInit.Run(); err != nil {
		return err
	}

	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Dir = destPath
	if err := gitAdd.Run(); err != nil {
		return err
	}

	gitCommit := exec.Command("git", "commit", "-m", `"First Commit"`)
	gitCommit.Dir = destPath
	if err := gitCommit.Run(); err != nil {
		return err
	}

	return nil
}

func (this *GitProjectInit) write(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
}

const replyScript = `package main

import (
	"os"
	pm "github.com/securenative/packman/pkg"
)

type MyData struct {
	PackageName string
	Args []string
}

func main() {
	// Args sent by packman's driver will be forwarded to here:
	args := os.Args[2:]

	// Build your own model to represent the templating you need
	model := MyData{PackageName: "my_pkg", Args: args}

	// Reply to packman's driver:
	pm.Reply(model)
}
`

const modeFile = `module packmanScript

require (
	github.com/securenative/packman latest
)
`
