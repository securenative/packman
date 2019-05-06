package controllers

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

var InitCommand = cli.Command{
	Name:      "init",
	Aliases:   []string{"i"},
	Usage:     "init <path>",
	UsageText: "Will create the minimal folder structure which is required by packman",
	Action: func(context *cli.Context) error {

		if context.NArg() != 1 {
			fmt.Println("init expects exactly 1 argument")
		}

		path := context.Args().Get(0)

		if path == "" {
			return errors.New("you must provide a path to the project")
		}

		return PackmanModule.ProjectInit.Init(path)
	},
}
