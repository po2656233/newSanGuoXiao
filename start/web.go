//go:build web

package start

import (
	"github.com/urfave/cli/v2"
	"sanguoxiao/nodes/web"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "web",
		Usage:     "run web node",
		UsageText: "node web --path=./examples/config/profile-gc.json --node=gc-web-1",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			web.Run(path, node)
			return nil
		},
	}
}
