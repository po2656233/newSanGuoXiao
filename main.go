//go:build master || leaf || game || gate || web || center || match

package main

import (
	"os"
	"superman/start"

	"github.com/urfave/cli/v2"
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
