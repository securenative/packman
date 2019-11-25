package controllers

import (
	"fmt"
	"github.com/securenative/packman/internal"
	"github.com/urfave/cli"
)

var PackController = cli.Command{
	Name:      "pack",
	Aliases:   []string{"p"},
	Usage:     "packman pack <path> <remote_url>",
	UsageText: "packing a folder by pushing it to the configured git's remote",
	Action: func(c *cli.Context) error {
		return pack(c)
	},
}

func pack(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("pack expects exactly 2 arguments but got %d arguments", c.NArg())
	}

	path := c.Args().Get(0)
	remote := c.Args().Get(1)
	return internal.M.TemplatingService.Pack(remote, path)
}
