package controllers

import (
	"fmt"
	"github.com/securenative/packman/internal"
	"github.com/urfave/cli"
)

var UnpackController = cli.Command{
	Name:      "unpack",
	Aliases:   []string{"u"},
	Usage:     "packman unpack <path> <remote_url> [-flagName flagValue]...",
	UsageText: "unpacking a template project with the given flags",
	Action: func(c *cli.Context) error {
		return unpack(c)
	},
}

func unpack(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("unpack expects exactly 2 arguments but got %d arguments", c.NArg())
	}

	path := c.Args().Get(0)
	remote := c.Args().Get(1)
	flagsMap := make(map[string]string)

	for _, flagName := range c.FlagNames() {
		flagsMap[flagName] = c.String(flagName)
	}

	return internal.M.TemplatingService.Unpack(remote, path, flagsMap)
}
