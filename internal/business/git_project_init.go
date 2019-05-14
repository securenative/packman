package business

import (
	"github.com/mingrammer/cfmt"
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

	cfmt.Info("Creating path ", path, "\n")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		cfmt.Error("Cannot create the following path: ", path, ", ", err.Error(), "\n")
		return err
	}

	cfmt.Info("Writing the main.go script file", "\n")
	scriptPath := filepath.Join(path, "main.go")
	if err := this.write(scriptPath, replyScript); err != nil {
		cfmt.Error("Cannot create ", scriptPath, ", ", err.Error(), "\n")
		return err
	}

	cfmt.Info("Writing the go.mod file", "\n")
	modPath := filepath.Join(path, "go.mod")
	if err := this.write(modPath, modeFile); err != nil {
		cfmt.Error("Cannot create ", scriptPath, ", ", err.Error(), "\n")
		return err
	}

	cfmt.Info("Initialing the git repository", "\n")
	gitInit := exec.Command("git", "init")
	gitInit.Dir = destPath
	if err := gitInit.Run(); err != nil {
		cfmt.Error("Cannot init git repository, ", err.Error(), "\n")
		return err
	}

	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Dir = destPath
	if err := gitAdd.Run(); err != nil {
		cfmt.Error("Failed to add untracked files to the git repository, ", err.Error(), "\n")
		return err
	}

	cfmt.Info("Creating the first commit", "\n")
	gitCommit := exec.Command("git", "commit", "-m", `"First Commit"`)
	gitCommit.Dir = destPath
	if err := gitCommit.Run(); err != nil {
		cfmt.Error("Failed to commit changes, ", err.Error(), "\n")
		return err
	}

	cfmt.Success("Packman package created successfully!")
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
	ProjectPath string
	Flags map[string]string
}

func main() {
	// flags sent by packman's driver will be forwarded to here:
	flags := pm.ParseFlags(os.Args[3:])

	// Build your own model to represent the templating you need
	model := MyData{
		PackageName: flags[pm.PackageNameFlag], 
		ProjectPath: flags[pm.PackagePathFlag], 
		Flags: flags,
	}

	// Reply to packman's driver:
	pm.Reply(model)
}
`

const modeFile = `module packmanScript

require (
	github.com/securenative/packman latest
)
`
