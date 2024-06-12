package center

import (
	"github.com/po2656233/superplace"
	"github.com/po2656233/superplace/components/cron"
	superGORM "github.com/po2656233/superplace/components/gorm"
	"sanguoxiao/internal/actors"
	"sanguoxiao/internal/conf"
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
	app.Register(conf.New())
	app.Register(db.New())
	app.Register(superGORM.NewComponent())

	app.AddActors(
		&actors.ActorDB{},
		&account.ActorAccount{},
		&ops.ActorOps{},
	)
	app.Startup()
}
