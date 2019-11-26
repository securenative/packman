package main

import (
	"github.com/securenative/packman/cmd/packman/controllers"
	"github.com/securenative/packman/internal/etc"
	"github.com/urfave/cli"
	"os"
)

func main() {
	commands := []cli.Command{
		controllers.PackController,
		controllers.UnpackController,
		controllers.RenderController,

		controllers.AuthController,
		controllers.ScriptEngineController,
	}

	app := cli.App{
		Name:     "packman",
		Version:  "0.3",
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		etc.PrintError(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
