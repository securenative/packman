package controllers

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

var PackCommand = cli.Command{
	Name:      "pack",
	Aliases:   []string{"p"},
	Usage:     "pack <package-name> <path>",
	UsageText: "Will pack the given folder and push it to the backend so it can be later located using the package-name",
	Action: func(context *cli.Context) error {

		if context.NArg() != 2 {
			fmt.Println("pack expects exactly 2 arguments")
		}

		packageName := context.Args().Get(0)
		path := context.Args().Get(1)

		if packageName == "" {
			return errors.New("you must provide a package name")
		}

		if path == "" {
			return errors.New("you must provide a path to the project")
		}

		return PackmanModule.Packer.Pack(packageName, path)
	},
}
