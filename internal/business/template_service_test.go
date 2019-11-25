package business

import (
	"errors"
	copy2 "github.com/otiai10/copy"
	"github.com/securenative/packman/internal/data"
	"github.com/securenative/packman/internal/etc"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var underTest = NewTemplateService(
	data.NewGitRemoteStorage(&envLocalStorage{}),
	data.NewGenericScriptEngine("go run"),
	data.NewGolangTemplateEngine(),
)

func TestTemplateService_Unpack(t *testing.T) {
	skipOnMissingEnv(t)
	gitPath := filepath.Join(os.TempDir(), "packman_service_test")
	defer os.RemoveAll(gitPath)

	err := underTest.Unpack("https://github.com/matang28/packman_git_test.git", gitPath,
		map[string]string{"Key": "My Key", "Value": "My Value"},
	)
	assert.Nil(t, err)

	content, err := etc.ReadFile(filepath.Join(gitPath, "template.txt"))
	assert.Nil(t, err)
	assert.EqualValues(t, "My Key My Value\n", content)
}

func TestTemplateService_Render(t *testing.T) {
	skipOnMissingEnv(t)
	gitPath := filepath.Join(os.TempDir(), "packman_service_test")
	defer os.RemoveAll(gitPath)

	git := data.NewGitRemoteStorage(&envLocalStorage{})
	err := git.Pull("https://github.com/matang28/packman_git_test.git", gitPath)
	assert.Nil(t, err)

	defer os.RemoveAll(gitPath + "-rendered")
	err = underTest.Render(gitPath, gitPath+"-rendered",
		map[string]string{"Key": "My Key", "Value": "My Value"},
	)
	assert.Nil(t, err)

	content, err := etc.ReadFile(filepath.Join(gitPath+"-rendered", "template.txt"))
	assert.Nil(t, err)
	assert.EqualValues(t, "My Key My Value\n", content)
}

func TestTemplateService_Pack(t *testing.T) {
	skipOnMissingEnv(t)
	t.Skip()
	gitPath := filepath.Join(os.TempDir(), "packman_service_test")
	git := data.NewGitRemoteStorage(&envLocalStorage{})
	defer os.RemoveAll(gitPath)

	err := git.Pull("https://github.com/matang28/packman_git_test.git", gitPath)
	assert.Nil(t, err)

	err = copy2.Copy(gitPath, gitPath+"-temp")
	assert.Nil(t, err)

	err = underTest.Pack("https://github.com/matang28/packman_git_test.git", gitPath+"-temp")
	assert.Nil(t, err)
}

type envLocalStorage struct{}

func (this *envLocalStorage) Put(key, value string) error {
	return nil
}

func (this *envLocalStorage) Get(key string) (string, error) {
	switch key {
	case string(data.GitUsername):
		return os.Getenv("USERNAME"), nil
	case string(data.GitPassword):
		return os.Getenv("PASSWORD"), nil
	}
	return "", errors.New("")
}

func skipOnMissingEnv(t *testing.T) {
	if os.Getenv("USERNAME") == "" {
		t.Skip("Skipped because no env variable called 'USERNAME' exists")
	}

	if os.Getenv("PASSWORD") == "" {
		t.Skip("Skipped because no env variable called 'PASSWORD' exists")
	}
}
