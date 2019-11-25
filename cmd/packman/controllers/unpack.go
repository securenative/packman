package controllers

import (
	"fmt"
	"github.com/securenative/packman/internal"
	"github.com/securenative/packman/internal/etc"
	"github.com/urfave/cli"
)

var UnpackController = cli.Command{
	Name:      "unpack",
	Aliases:   []string{"u"},
	Usage:     "packman unpack <remote_url> <path> [-flagName flagValue]...",
	UsageText: "unpacking a template project with the given flags",
	Action: func(c *cli.Context) error {
		return unpack(c)
	},
}

func unpack(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("unpack expects exactly 2 arguments but got %d arguments", c.NArg())
	}

	remote := c.Args().Get(0)
	path := c.Args().Get(1)
	flagsMap := etc.ArgsToFlagsMap(c.Args()[2:])

	return internal.M.TemplatingService.Unpack(remote, path, flagsMap)
}
