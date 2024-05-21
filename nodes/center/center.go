package center

import (
	"github.com/po2656233/superplace"
	"github.com/po2656233/superplace/components/cron"
	"sanguoxiao/internal/data"
	"sanguoxiao/nodes/center/db"
	"sanguoxiao/nodes/center/module/account"
	"sanguoxiao/nodes/center/module/ops"
)

func Run(profileFilePath, nodeId string) {
	app := superplace.Configure(
		profileFilePath,
		nodeId,
		false,
		superplace.Cluster,
	)

	app.Register(superCron.New())
	app.Register(data.New())
	app.Register(db.New())

	app.AddActors(
		&account.ActorAccount{},
		&ops.ActorOps{},
	)

	app.Startup()
}
