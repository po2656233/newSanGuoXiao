package start

import (
	"fmt"
	"github.com/po2656233/superplace/const"
	"github.com/po2656233/superplace/net/parser/simple"
	"github.com/urfave/cli/v2"
	"strings"
	"superman/internal/rpc"
)

func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:      "version",
		Aliases:   []string{"ver", "v"},
		Usage:     "view version",
		UsageText: "game cluster node version",
		Action: func(c *cli.Context) error {
			fmt.Println(superConst.Version())
			return nil
		},
	}
}

func getParameters(c *cli.Context) (path, node string) {
	path = c.String("path")
	node = c.String("node")
	// 网络协议区分
	if !strings.Contains(node, "pomelo") {
		// 加载协议文件
		rpc.LoadMsgInfos()
		simple.SetParseProtoFunc(rpc.ParseProto)
	}
	return path, node
}

func getFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Usage:    "profile config path",
			Required: false,
			Value:    "./examples/config/profile-gc.json",
		},
		&cli.StringFlag{
			Name:     "node",
			Usage:    "node id name",
			Required: true,
			Value:    "",
		},
	}
}
