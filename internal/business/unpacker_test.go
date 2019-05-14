package business

import (
	"github.com/securenative/packman/internal/data"
	"github.com/securenative/packman/pkg"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const expected = `
package my-pkg
func main() {
fmt.Println("Hello")
fmt.Println("World")
fmt.Println("my-pkg")
}`

func TestPackmanUnpacker_Unpack(t *testing.T) {
	replacer := strings.NewReplacer("\n", "", "\t", "")
	unpacker := NewPackmanUnpacker(&mockBackend{}, data.NewGoTemplateEngine(), data.NewGoScriptEngine())

	path := filepath.Join(os.TempDir(), "packtest", "unpack")
	err := unpacker.Unpack("my-pkg", path, []string{"--", pkg.PackageNameFlag, "my-pkg", "-a", "Hello", "-b", "World"})
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(path, "testfile.go"))
	assert.Nil(t, err)
	assert.Equal(t, replacer.Replace(expected), replacer.Replace(string(bytes)))

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
package {{{ .PackageName }}}

func main() {
	{{{ range .Flags }}}
	fmt.Println("{{{ . }}}")
	{{{ end }}}
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
