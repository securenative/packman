package controllers

import (
	"fmt"
	"github.com/securenative/packman/internal"
	"github.com/urfave/cli"
	"strings"
)

var AuthController = cli.Command{
	Name:      "auth",
	Aliases:   []string{"a"},
	Usage:     "packman auth <username> <password>",
	UsageText: "saving the auth information to your git repositories",
	Action: func(c *cli.Context) error {
		return auth(c)
	},
}

var ScriptEngineController = cli.Command{
	Name:      "script",
	Aliases:   []string{"s"},
	Usage:     "packman script <script_command>",
	UsageText: "changes the command which meant to run the packman template script",
	Action: func(c *cli.Context) error {
		return changeScriptEngine(c)
	},
}

func auth(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("auth expects exactly 2 arguments but got %d arguments", c.NArg())
	}
	username := c.Args().Get(0)
	password := c.Args().Get(1)
	return internal.M.ConfigService.SetAuth(username, password)
}

func changeScriptEngine(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("script expects exactly 1 arguments but got %d arguments", c.NArg())
	}
	script := c.Args().Get(0)
	script = strings.ReplaceAll(script, `"`, "")
	return internal.M.ConfigService.SetDefaultEngine(script)
}
