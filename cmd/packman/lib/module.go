package lib

import (
	"github.com/securenative/packman/internal/business"
	"github.com/securenative/packman/internal/data"
	"gopkg.in/urfave/cli.v2"
)

type PackmanModule struct {
	Config         PackmanConfig
	configStore    data.ConfigStore
	backend        data.Backend
	templateEngine data.TemplateEngine
	scriptEngine   data.ScriptEngine

	ProjectInit business.ProjectInit
	Packer      business.Packer
	Unpacker    business.Unpacker

	Commands []*cli.Command
}

func NewPackmanModule(config PackmanConfig, commands []*cli.Command) *PackmanModule {

	configStore := data.NewLocalConfigStore(config.ConfigPath)
	backend := data.NewGithubBackend(configStore)
	template := data.NewGoTemplateEngine()
	script := data.NewGoScriptEngine()

	init := business.NewGitProjectInit()
	packer := business.NewPackmanPacker(backend)
	unpacker := business.NewPackmanUnpacker(backend, template, script)

	return &PackmanModule{
		Config:         config,
		configStore:    configStore,
		backend:        backend,
		templateEngine: template,
		scriptEngine:   script,
		ProjectInit:    init,
		Packer:         packer,
		Unpacker:       unpacker,
		Commands:       commands,
	}
}
