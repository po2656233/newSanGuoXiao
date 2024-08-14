//go:build match

package start

import (
	"github.com/urfave/cli/v2"
	"superman/nodes/match"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "match",
		Usage:     "run match node",
		UsageText: "node match --path=./examples/config/profile-gc.json --node=20001",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			match.Run(path, node)
			return nil
		},
	}
}
