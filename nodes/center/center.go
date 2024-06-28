package center

import (
	"github.com/po2656233/superplace"
	"github.com/po2656233/superplace/components/cron"
	superGORM "github.com/po2656233/superplace/components/gorm"
	"superman/internal/actors"
	"superman/internal/conf"
	"superman/nodes/center/db"
	"superman/nodes/center/module/account"
	"superman/nodes/center/module/ops"
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
