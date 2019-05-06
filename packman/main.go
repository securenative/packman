package main

import (
	"os"
	pm "github.com/securenative/packman/pkg"
)

type MyData struct {
	PackageName string
	Args []string
}

func main() {
	// Args sent by packman's driver will be forwarded to here:
	args := os.Args[2:]

	// Build your own model to represent the templating you need
	model := MyData{PackageName: "my_pkg", Args: args}

	// Reply to packman's driver:
	pm.Reply(model)
}
