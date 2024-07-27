package player

import (
	cst "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/category/CardGame/sanguoxiao"
	"superman/nodes/game/module/category/SportGame/chinesechess"
	"time"
)

////////////////////////////创建游戏////////////////////////////////////////////////

// NewGame 创建游戏
func NewGame(gid, tid int64) manger.IGameOperate {
	info := manger.GetGameInfoMgr().GetGame(gid)
	if info == nil || info.State == protoMsg.GameState_InitTB || info.State == protoMsg.GameState_CloseTB {
		return nil
	}
	info.Id = tid
	game := &manger.Game{
		GameInfo:   info,
		IsStart:    true,
		IsClear:    false,
		ReadyCount: 0,
		RunCount:   0,
		TimeStamp:  time.Now().Unix(),
		PlayIDList: make([]int64, 0),
	}
	switch gid {
	case cst.ChineseChess:
		return chinesechess.New(game)
	case cst.Chess:
	case cst.SanGuoXiao:
		return sanguoxiao.New(game)
	default:
	}
	return nil
}
