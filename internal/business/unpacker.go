package business

import (
	. "github.com/otiai10/copy"
	"github.com/securenative/packman/internal/data"
	"github.com/securenative/packman/pkg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PackmanUnpacker struct {
	backend        data.Backend
	templateEngine data.TemplateEngine
	scriptEngine   data.ScriptEngine
}

func NewPackmanUnpacker(backend data.Backend, templateEngine data.TemplateEngine, scriptEngine data.ScriptEngine) *PackmanUnpacker {
	return &PackmanUnpacker{backend: backend, templateEngine: templateEngine, scriptEngine: scriptEngine}
}

func (this *PackmanUnpacker) DryUnpack(sourcePath string, destPath string, args []string) error {
	if err := os.RemoveAll(destPath); err != nil {
		return err
	}

	if err := Copy(sourcePath, destPath); err != nil {
		return err
	}

	if err := this.render(destPath, args); err != nil {
		return err
	}

	if err := this.clean(destPath); err != nil {
		return err
	}

	return nil
}

func (this *PackmanUnpacker) Unpack(name string, destPath string, args []string) error {
	if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
		return err
	}

	if err := this.backend.Pull(name, destPath); err != nil {
		return err
	}

	if err := this.render(destPath, args); err != nil {
		return err
	}

	if err := this.clean(destPath); err != nil {
		return err
	}

	return nil
}

func (this *PackmanUnpacker) clean(destPath string) error {
	if err := os.RemoveAll(filepath.Join(destPath, ".git")); err != nil {
		return err
	}

	if err := os.RemoveAll(filepath.Join(destPath, "packman")); err != nil {
		return err
	}

	return nil
}

func (this *PackmanUnpacker) render(destPath string, args []string) error {
	_ = os.Setenv("PACKMAN_PROJECT", packmanPath(destPath))

	scriptFile := filepath.Join(packmanPath(destPath), "main.go")
	if err := this.scriptEngine.Run(scriptFile, args); err != nil {
		return err
	}

	dataModel, err := pkg.ReadReply(packmanPath(destPath))
	if err != nil {
		return err
	}

	err = filepath.Walk(destPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, "packman") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			rendered, err := this.templateEngine.Render(string(content), dataModel)
			if err != nil {
				return err
			}

			if err = ioutil.WriteFile(path, []byte(rendered), os.ModePerm); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
