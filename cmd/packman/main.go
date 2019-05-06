package main

import (
	"github.com/securenative/packman/cmd/packman/controllers"
	"github.com/securenative/packman/cmd/packman/lib"
	"gopkg.in/urfave/cli.v2"
	"os"
	"path/filepath"
)

func main() {

	cfg := parseConfig()
	module := lib.NewPackmanModule(cfg, []*cli.Command{
		&controllers.InitCommand,
		&controllers.PackCommand,
		&controllers.UnpackCommand,
	})
	controllers.PackmanModule = module

	app := cli.App{
		Name:     "packman",
		Version:  "0.1",
		Commands: module.Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}

func parseConfig() lib.PackmanConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	configPath := filepath.Join(home, ".packman")
	cfg := lib.PackmanConfig{
		ConfigPath: configPath,
	}
	return cfg
}
