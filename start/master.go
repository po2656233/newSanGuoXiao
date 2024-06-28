//go:build master
// +build master

package start

import (
	"github.com/urfave/cli/v2"
	"superman/nodes/master"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "master",
		Usage:     "run master node",
		UsageText: "node master --path=./examples/config/profile-gc.json --node=gc-master",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			master.Run(path, node)
			return nil
		},
	}
}
