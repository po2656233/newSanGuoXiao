package master

import (
	"github.com/po2656233/superplace"
)

func Run(profileFilePath, nodeId string) {
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)
	app.Startup()
}
