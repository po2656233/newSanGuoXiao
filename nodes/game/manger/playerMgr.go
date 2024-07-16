package manger

import protoMsg "superman/internal/protocol/gofile"

// Player 玩家属性
type Player struct {
	//基本信息
	*protoMsg.PlayerInfo
	//用户包裹
	Knapsack *protoMsg.KnapsackInfo
	// 用于追踪用户行为
	GameHandle *IGameOperate // 所在房间 时刻更新最新房间 状态
}
