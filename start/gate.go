//go:build gate

package start

import (
	"github.com/urfave/cli/v2"
	"sanguoxiao/nodes/gate"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "gate",
		Usage:     "run gate node",
		UsageText: "node gate --path=./examples/config/profile-gc.json --node=gc-gate-1",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			gate.Run(path, node)
			return nil
		},
	}
}
