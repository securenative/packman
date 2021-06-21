package business

import (
	"fmt"
	copy2 "github.com/otiai10/copy"
	"github.com/securenative/packman/internal/data"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const PackageName = "package_name"
const PackagePath = "package_path"

type templateService struct {
	remoteStorage  data.RemoteStorage
	scriptEngine   data.ScriptEngine
	templateEngine data.TemplateEngine
}

func NewTemplateService(remoteStorage data.RemoteStorage, scriptEngine data.ScriptEngine, templateEngine data.TemplateEngine) *templateService {
	return &templateService{remoteStorage: remoteStorage, scriptEngine: scriptEngine, templateEngine: templateEngine}
}

func (this *templateService) Render(templatePath string, packagePath string, flags map[string]string) error {
	if templatePath != packagePath {
		_ = os.RemoveAll(packagePath)
		if err := copy2.Copy(templatePath, packagePath); err != nil {
			return err
		}
	}

	flags[PackagePath] = packagePath
	flags[PackageName] = filepath.Base(packagePath)

	scriptPath, err := toScriptPath(packagePath)
	if err != nil {
		return err
	}

	scriptData, err := this.scriptEngine.Run(scriptPath, flags)
	if err != nil {
		return err
	}

	err = filepath.Walk(packagePath, func(path string, info os.FileInfo, err error) error {
		if shouldSkip(info, path) {
			return nil
		}

		ierr := this.templateEngine.Run(path, scriptData)
		if ierr != nil {
			return ierr
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = os.RemoveAll(filepath.Join(packagePath, "packman"))
	if err != nil {
		return err
	}

	return nil
}

func (this *templateService) Pack(remoteUrl string, packagePath string) error {
	return this.remoteStorage.Push(packagePath, remoteUrl)
}

func (this *templateService) Unpack(remoteUtl string, packagePath string, flags map[string]string) error {
	err := this.remoteStorage.Pull(remoteUtl, packagePath)
	if err != nil {
		return err
	}

	err = os.RemoveAll(filepath.Join(packagePath, ".git"))
	if err != nil {
		return err
	}

	return this.Render(packagePath, packagePath, flags)
}

func toScriptPath(prefix string) (string, error) {
	packmanPath := filepath.Join(prefix, "packman")
	packmanDir, err := ioutil.ReadDir(packmanPath)
	if err != nil {
		return "", err
	}

	for _, file := range packmanDir {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "main") {
			return filepath.Join(packmanPath, file.Name()), nil
		}
	}
	return "", fmt.Errorf("in order for packman to work you must have a 'main.*' file within your packman folder")
}

func shouldSkip(info os.FileInfo, path string) bool {
	if info.IsDir() {
		return true
	}
	if strings.Contains(path, ".git/") {
		return true
	}
	return false
}
