package internal

import (
	"errors"
	"github.com/securenative/packman/internal/business"
	"github.com/securenative/packman/internal/data"
	"os"
	"path/filepath"
	"runtime"
)

type Module struct {
	remoteStorage  data.RemoteStorage
	scriptEngine   data.ScriptEngine
	templateEngine data.TemplateEngine
	localStorage   data.LocalStorage

	TemplatingService business.TemplatingService
	ConfigService     business.ConfigService
}

var M *Module

func init() {
	home, err := homeDir()
	if err != nil {
		panic(err)
	}

	localFilePath := filepath.Join(home, "packman_config.json")
	localStorage, err := data.NewFileLocalStorage(localFilePath)
	if err != nil {
		panic(err)
	}

	scriptCommand, err := localStorage.Get(string(data.DefaultScript))
	if err != nil {
		scriptCommand = "go run"
	}

	remoteStorage := data.NewGitRemoteStorage(localStorage)
	scriptEngine := data.NewGenericScriptEngine(scriptCommand)
	templateEngine := data.NewGolangTemplateEngine()

	M = &Module{
		remoteStorage:     remoteStorage,
		scriptEngine:      scriptEngine,
		templateEngine:    templateEngine,
		localStorage:      localStorage,
		TemplatingService: business.NewTemplateService(remoteStorage, scriptEngine, templateEngine),
		ConfigService:     business.NewConfigService(localStorage),
	}
}

func homeDir() (string, error) {
	env, enverr := "HOME", "$HOME"
	switch runtime.GOOS {
	case "windows":
		env, enverr = "USERPROFILE", "%userprofile%"
	case "plan9":
		env, enverr = "home", "$home"
	}
	if v := os.Getenv(env); v != "" {
		return v, nil
	}
	// On some geese the home directory is not always defined.
	switch runtime.GOOS {
	case "nacl":
		return "/", nil
	case "android":
		return "/sdcard", nil
	case "darwin":
		if runtime.GOARCH == "arm" || runtime.GOARCH == "arm64" {
			return "/", nil
		}
	}
	return "", errors.New(enverr + " is not defined")
}
