package player

import (
	log "github.com/po2656233/superplace/logger"
	cst "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/category/BattleGame/brbaccarat"
	"superman/nodes/game/module/category/BattleGame/brcowcow"
	"superman/nodes/game/module/category/CardGame/sanguoxiao"
	"superman/nodes/game/module/category/SportGame/chinesechess"
	"time"
)

////////////////////////////创建游戏////////////////////////////////////////////////

// NewGame 创建游戏
func NewGame(gid int64, tb *manger.Table) manger.IGameOperate {
	info := manger.GetGameInfoMgr().GetGame(gid)
	if info == nil || info.State == protoMsg.GameState_InitTB || info.State == protoMsg.GameState_CloseTB {
		return nil
	}
	// 剩余次数没有了,则不再创建该牌桌
	if tb.Remain == cst.FINISH {
		log.Infof("牌桌[%d] 已经没有剩余次数了", tb.Id)
		return nil
	}
	// 使用牌桌ID
	info.Id = tb.Id
	game := &manger.Game{
		GameInfo:   info,
		IsStart:    false,
		IsClear:    false,
		ReadyCount: 0,
		RunCount:   0,
		TimeStamp:  time.Now().Unix(),
	}
	switch gid {
	case cst.ChineseChess:
		return chinesechess.New(game, tb)
	case cst.Chess:
	case cst.SanGuoXiao:
		return sanguoxiao.New(game)
	case cst.Baccarat:
		return brbaccarat.New(game, tb)
	case cst.BrCowcow:
		return brcowcow.New(game, tb)
	default:
	}
	return nil
}
