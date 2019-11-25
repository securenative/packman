package controllers

import (
	"fmt"
	"github.com/securenative/packman/internal"
	"github.com/securenative/packman/internal/etc"
	"github.com/urfave/cli"
)

var RenderController = cli.Command{
	Name:      "render",
	Aliases:   []string{"r"},
	Usage:     "packman render <path> [-flagName flagValue]...",
	UsageText: "unpacking a template project with the given flags",
	Action: func(c *cli.Context) error {
		return render(c)
	},
}

func render(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("unpack expects exactly 1 argument but got %d arguments", c.NArg())
	}

	path := c.Args().Get(0)
	flagsMap := etc.ArgsToFlagsMap(c.Args()[1:])

	return internal.M.TemplatingService.Render(path, fmt.Sprintf("%s-rendered", path), flagsMap)
}
