//go:build game

package start

import (
	"github.com/urfave/cli/v2"
	"superman/nodes/game"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "game",
		Usage:     "run game node",
		UsageText: "node game --path=./examples/config/profile-gc.json --node=20001",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			game.Run(path, node)
			return nil
		},
	}
}
