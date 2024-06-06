//go:build center

package start

import (
	"github.com/urfave/cli/v2"
	"sanguoxiao/nodes/center"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "center",
		Usage:     "run center node",
		UsageText: "node center --path=./examples/config/profile-gc.json --node=gc-center",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			center.Run(path, node)
			return nil
		},
	}
}
