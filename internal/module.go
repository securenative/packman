package internal

import (
	"github.com/securenative/packman/internal/business"
	"github.com/securenative/packman/internal/data"
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
	localStorage, err := data.NewFileLocalStorage("")
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
