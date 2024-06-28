//go:build master || leaf || game || gate || web || center

package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"superman/start"
)

func main() {
	args := os.Args
	app := &cli.App{
		Name:        "game cluster node",
		Description: "game cluster node examples",
		Commands: []*cli.Command{
			start.VersionCommand(),
			start.Command(),
		},
	}
	_ = app.Run(args)
}
