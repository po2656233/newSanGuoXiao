// go:build leaf
package start

//package start
//
//import (
//	"github.com/urfave/cli/v2"
//	"superman/nodes/leaf"
//)
//
//func Command() *cli.Command {
//	return &cli.Command{
//		Name:      "leaf",
//		Usage:     "run leaf node",
//		UsageText: "node leaf --path=./examples/config/profile-gc.json --node=20001",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			leaf.Run(path, node)
//			return nil
//		},
//	}
//}
