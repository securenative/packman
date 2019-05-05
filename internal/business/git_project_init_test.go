package business

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPackmanProjectInit_Init(t *testing.T) {

	init := NewGitProjectInit()

	tempDir := filepath.Join(os.TempDir(), "packtest")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		t.Fail()
	}

	err := init.Init(filepath.Join(tempDir, "my_pkg"))
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(tempDir, "my_pkg", "packman", "main.go"))
	assert.Nil(t, err)

	assert.Equal(t, replyScript, string(bytes))
	_ = os.RemoveAll(tempDir)

}
