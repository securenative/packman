package controllers

import (
	"errors"
	"github.com/securenative/packman/pkg"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

var UnpackCommand = cli.Command{
	Name:      "unpack",
	Aliases:   []string{"u"},
	Usage:     "unpack <package-name> <dest-path>",
	UsageText: "Will unpack the given package to the given destination path",
	Action:    unpack,
}

var DryUnpackCommand = cli.Command{
	Name:      "render",
	Aliases:   []string{"r"},
	Usage:     "render <package-name> <source-path> <dest-path>",
	UsageText: "will render the source project into the dest path WARNING: THIS WILL REMOVE THE DEST PATH",
	Action:    render,
}

func render(context *cli.Context) error {

	packageName := context.Args().Get(0)
	sourcePath := context.Args().Get(1)
	destPath := context.Args().Get(2)

	if packageName == "" {
		return errors.New("you must provide a package name")
	}

	if sourcePath == "" {
		return errors.New("you must provide a source path")
	}

	if destPath == "" {
		return errors.New("you must provide a destination path")
	}

	flags := extractFlags(context, packageName, destPath)

	return PackmanModule.Unpacker.DryUnpack(sourcePath, destPath, flags)
}

func unpack(context *cli.Context) error {

	packageName := context.Args().Get(0)
	path := context.Args().Get(1)

	if packageName == "" {
		return errors.New("you must provide a package name")
	}

	if path == "" {
		return errors.New("you must provide a path to the project")
	}

	flags := extractFlags(context, packageName, path)

	return PackmanModule.Unpacker.Unpack(packageName, path, flags)
}

func extractFlags(context *cli.Context, packageName string, path string) []string {
	flags := flagsArray(context)
	flags = append(flags, pkg.PackageNameFlag, packageName)
	flags = append(flags, pkg.PackagePathFlag, path)
	return flags
}

func flagsArray(ctx *cli.Context) []string {
	out := make([]string, 0)

	for idx, arg := range ctx.Args().Slice() {
		if strings.HasPrefix(arg, "-") {
			out = append(out, arg[1:], ctx.Args().Get(idx+1))
		} else if strings.HasPrefix(arg, "--") {
			out = append(out, arg[2:], ctx.Args().Get(idx+1))
		}
	}

	return out
}
