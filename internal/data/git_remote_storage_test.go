package data

import (
	"errors"
	"fmt"
	"github.com/securenative/packman/internal/etc"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestGitRemoteStorage_PullWithoutAuth(t *testing.T) {
	gitPath := filepath.Join(os.TempDir(), "packman_git_test")
	git := NewGitRemoteStorage(&emptyLocalStorage{})

	err := git.Pull("https://github.com/matang28/packman_git_test.git", gitPath)
	assert.NotNil(t, err)
}

func TestGitRemoteStorage_PullWithAuth(t *testing.T) {
	skipOnMissingEnv(t)
	gitPath := filepath.Join(os.TempDir(), "packman_git_test")
	git := NewGitRemoteStorage(&envLocalStorage{})

	defer os.RemoveAll(gitPath)
	err := git.Pull("https://github.com/matang28/packman_git_test.git", gitPath)
	assert.Nil(t, err)

	content, err := etc.ReadFile(filepath.Join(gitPath, "README.md"))
	assert.Nil(t, err)
	assert.EqualValues(t, "This is a test\n", content)
}

func TestGitRemoteStorage_Push(t *testing.T) {
	skipOnMissingEnv(t)
	gitPath := filepath.Join(os.TempDir(), "packman_git_test")
	customFilePath := filepath.Join(gitPath, fmt.Sprintf("file-%d", rand.Int()))
	git := NewGitRemoteStorage(&envLocalStorage{})

	defer os.RemoveAll(gitPath)

	clone(git, gitPath, t)
	err := etc.WriteFile(customFilePath, "dummy data", etc.StringEncoder)
	assert.Nil(t, err)

	err = git.Push(gitPath, "https://github.com/matang28/packman_git_test.git")
	assert.Nil(t, err)

	os.RemoveAll(gitPath)
	clone(git, gitPath, t)
	assert.FileExists(t, customFilePath)

	err = os.Remove(customFilePath)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	err = git.Push(gitPath, "https://github.com/matang28/packman_git_test.git")
	assert.Nil(t, err)
}

func clone(git *gitRemoteStorage, gitPath string, t *testing.T) {
	err := git.Pull("https://github.com/matang28/packman_git_test.git", gitPath)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
}

func skipOnMissingEnv(t *testing.T) {
	if os.Getenv("USERNAME") == "" {
		t.Skip("Skipped because no env variable called 'USERNAME' exists")
	}

	if os.Getenv("PASSWORD") == "" {
		t.Skip("Skipped because no env variable called 'PASSWORD' exists")
	}
}

type envLocalStorage struct{}

func (this *envLocalStorage) Put(key, value string) error {
	return nil
}

func (this *envLocalStorage) Get(key string) (string, error) {
	switch key {
	case string(GitUsername):
		return os.Getenv("USERNAME"), nil
	case string(GitPassword):
		return os.Getenv("PASSWORD"), nil
	}
	return "", errors.New("")
}

type emptyLocalStorage struct{}

func (this *emptyLocalStorage) Put(key, value string) error {
	return nil
}

func (this *emptyLocalStorage) Get(key string) (string, error) {
	return "", errors.New("")
}
