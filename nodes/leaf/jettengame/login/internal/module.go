package internal

import (
	"github.com/po2656233/goleaf/module"
	"superman/nodes/leaf/jettengame/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	////奖励时goods数据
	//list := &protoMsg.GoodsList{}
	//list.AllGoods = make([]*protoMsg.GoodsInfo, 0)
	//info := &protoMsg.GoodsInfo{
	//	ID:     1001,
	//	Kind:   1,
	//	Name:   "通用房卡",
	//	Amount: 30,
	//}
	//list.AllGoods = append(list.AllGoods, info)
	//if data, err := base.PBToBytes(list); err == nil {
	//	log.Release("%v", data)
	//}
	//list := &protoMsg.RechargeReq{
	//	UserID:  0,
	//	ByUID:   0,
	//	Payment: 0,
	//	Method:  0,
	//	Reason:  "",
	//}
	//if data, err := base.PBToBytes(list); err == nil {
	//	log.Release("%v", data)
	//}

}

func (m *Module) OnDestroy() {

}
