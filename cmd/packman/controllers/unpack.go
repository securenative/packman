package controllers

import "gopkg.in/urfave/cli.v2"

var UnpackCommand = cli.Command{
	Name:      "unpack",
	Aliases:   []string{"u"},
	Usage:     "unpack <package-name> <dest-path>",
	UsageText: "Will unpack the given package to the given destination path",
	Action: func(context *cli.Context) error {
		return nil
	},
}
