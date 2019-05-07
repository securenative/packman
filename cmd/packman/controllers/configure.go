package controllers

import (
	"github.com/securenative/packman/internal/data"
	"gopkg.in/urfave/cli.v2"
)

var ConfigureCommand = cli.Command{
	Name:      "config",
	Aliases:   []string{"c"},
	Usage:     "configure <scope> [--key value]",
	UsageText: "Will set the configuration of the given scope using key-value flag pairs",
	Subcommands: []*cli.Command{
		{
			Name: "github",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:   "username",
					Hidden: false,
				},
				&cli.StringFlag{
					Name:   "token",
					Hidden: false,
				},
				&cli.BoolFlag{
					Name:   "private",
					Hidden: false,
				},
			},
			Action: func(context *cli.Context) error {
				return configureGithub(context)
			}},
	},
}

func configureGithub(context *cli.Context) error {
	cfg := data.GithubConfig{
		Username:    context.String("username"),
		Token:       context.String("token"),
		PrivatePush: context.Bool("private"),
	}
	return PackmanModule.ConfigStore.Put(PackmanModule.Backend.ConfigKey(), cfg)
}
