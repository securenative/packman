package controllers

import "gopkg.in/urfave/cli.v2"

var PackCommand = cli.Command{
	Name:      "pack",
	Aliases:   []string{"p"},
	Usage:     "pack <package-name> <path>",
	UsageText: "Will pack the given folder and push it to the backend so it can be later located using the package-name",
	Action: func(context *cli.Context) error {
		return nil
	},
}
