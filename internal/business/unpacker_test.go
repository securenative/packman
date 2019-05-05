package business

import (
	"github.com/securenative/packman/internal/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPackmanUnpacker_Unpack(t *testing.T) {

	unpacker := NewPackmanUnpacker(&mockBackend{}, data.NewGoTemplateEngine(), data.NewGoScriptEngine())

	path := filepath.Join(os.TempDir(), "packtest", "unpack")
	err := unpacker.Unpack("my-pkg", path, []string{"Hello", "World"})

	assert.Nil(t, err)

	_ = os.RemoveAll(path)
}

type mockBackend struct {
}

func (mockBackend) Push(name string, source string) error {
	return nil
}

func (mockBackend) Pull(name string, destination string) error {
	init := NewGitProjectInit()

	if err := init.Init(destination); err != nil {
		return err
	}

	content := `
package {{ .PackageName }}

func main() {
	{{ range .Args }}
	fmt.Println("{{ . }}")
	{{ end }}
}
`
	if err := ioutil.WriteFile(filepath.Join(destination, "testfile.go"), []byte(content), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (mockBackend) ConfigKey() string {
	return ""
}
