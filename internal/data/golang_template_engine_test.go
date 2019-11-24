package data

import (
	"github.com/securenative/packman/internal/etc"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGolangTemplateEngine_Run(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "packman_template_test.any")
	te := NewGolangTemplateEngine()

	defer os.Remove(filePath)
	err := etc.WriteFile(filePath, "{{{ .Key }}} {{{ .Value }}}", etc.StringEncoder)
	assert.Nil(t, err)

	err = te.Run(filePath, map[string]interface{}{"Key": "key", "Value": "value"})
	assert.Nil(t, err)

	content, err := etc.ReadFile(filePath)
	assert.Nil(t, err)
	assert.EqualValues(t, "key value", content)
}
