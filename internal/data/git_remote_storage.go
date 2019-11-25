package data

import (
	"github.com/securenative/packman/internal/etc"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"time"
)

type gitRemoteStorage struct {
	localStorage LocalStorage
}

func NewGitRemoteStorage(localStorage LocalStorage) *gitRemoteStorage {
	return &gitRemoteStorage{localStorage: localStorage}
}

func (this *gitRemoteStorage) getAuth() transport.AuthMethod {
	username, _ := this.localStorage.Get(string(GitUsername))
	password, _ := this.localStorage.Get(string(GitPassword))

	if username == "" || password == "" {
		return nil
	}

	return &http.BasicAuth{
		Username: username,
		Password: password,
	}
}

func (this *gitRemoteStorage) Pull(remotePath string, localPath string) error {
	etc.PrintInfo("Pulling %s into %s...\n", remotePath, localPath)
	_, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL:      remotePath,
		Auth:     this.getAuth(),
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}

func (this *gitRemoteStorage) Push(localPath string, remotePath string) error {
	repo, err := git.PlainOpen(localPath)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	if err = w.AddGlob("."); err != nil {
		return err
	}

	_, err = w.Commit("Pushed by packman", &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  "packman",
			Email: "",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       this.getAuth(),
		Progress:   os.Stdout,
	})
	if err != nil {
		return err
	}
	etc.PrintSuccess("Project was pushed to %s successfully.\n", remotePath)
	return nil
}
