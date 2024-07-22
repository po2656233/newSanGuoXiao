package player

import (
	cst "superman/internal/constant"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/category/CardGame/sanguoxiao"
	"superman/nodes/game/module/category/SportGame/chinesechess"
	"time"
)

////////////////////////////创建游戏////////////////////////////////////////////////

// NewGame 创建游戏
func NewGame(tid, gid int64) manger.IGameOperate {
	info := manger.GetGameInfoMgr().GetGame(gid)
	if info == nil {
		return nil
	}
	game := &manger.Game{
		GameInfo:   info,
		Tid:        tid,
		IsStart:    true,
		IsClear:    false,
		ReadyCount: 0,
		RunCount:   0,
		TimeStamp:  time.Now().Unix(),
	}
	switch info.Id {
	case cst.ChineseChess:
		return chinesechess.New(game)
	case cst.Chess:
	case cst.SanGuoXiao:
		return sanguoxiao.New(game)
	default:

	}
	return nil
}
