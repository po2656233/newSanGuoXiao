package player

import (
	cst "superman/internal/constant"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/category/CardGame/sanguoxiao"
	"time"
)

////////////////////////////创建游戏////////////////////////////////////////////////

// NewGame 创建游戏
func NewGame(gid int64) manger.IGameOperate {
	info := manger.GetGameInfoMgr().GetGame(gid)
	if info == nil {
		return nil
	}
	game := &manger.Game{
		GameInfo:   info,
		IsStart:    true,
		IsClear:    false,
		ReadyCount: 0,
		RunCount:   0,
		TimeStamp:  time.Now().Unix(),
	}
	switch info.Id {
	case cst.ChineseChess:
	case cst.Chess:
	case cst.SanGuoXiao:
		return sanguoxiao.New(game)
	default:

	}
	return nil
}
