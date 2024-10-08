//go:build chat

package start

import (
	"github.com/urfave/cli/v2"
	"superman/nodes/chat"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "chat",
		Usage:     "run chat node",
		UsageText: "node chat --path=./examples/config/profile-gc.json --node=gc-chat-1",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			chat.Run(path, node)
			return nil
		},
	}
}
