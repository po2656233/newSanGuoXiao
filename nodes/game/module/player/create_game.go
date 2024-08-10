package player

import (
	log "github.com/po2656233/superplace/logger"
	cst "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/game/manger"
	"superman/nodes/game/module/category/BattleGame/brbaccarat"
	"superman/nodes/game/module/category/BattleGame/brcowcow"
	"superman/nodes/game/module/category/BattleGame/brtigerXdragon"
	"superman/nodes/game/module/category/BattleGame/brtoubao"
	"superman/nodes/game/module/category/BattleGame/brtuitongzi"
	"superman/nodes/game/module/category/CardGame/sanguoxiao"
	"superman/nodes/game/module/category/SportGame/chinesechess"
	"time"
)

////////////////////////////创建游戏////////////////////////////////////////////////

// NewGame 创建游戏
func NewGame(gid int64, tb *Table) (handel IGameOperate, name string) {
	info := GetGameInfoMgr().GetGame(gid)
	if info == nil || info.State == protoMsg.GameState_InitTB || info.State == protoMsg.GameState_CloseTB {
		return
	}
	// 剩余次数没有了,则不再创建该牌桌
	if tb.Remain == cst.FINISH {
		log.Infof("[%v:%v]牌桌 已经没有剩余次数了", tb.Id, tb.Name)
		return
	}
	// 使用牌桌ID
	info.Id = tb.Id
	game := &Game{
		// 切记 此处GameInfo不为指针
		GameInfo: protoMsg.GameInfo{
			Id:        info.Id,
			Kid:       info.Kid,
			Name:      info.Name,
			Lessscore: info.Lessscore,
			Scene:     info.Scene,
			State:     info.State,
			MaxPlayer: info.MaxPlayer,
			HowToPlay: info.HowToPlay,
		},
		IsStart:    false,
		IsClear:    false,
		ReadyCount: 0,
		RunCount:   0,
		TimeStamp:  time.Now().Unix(),
	}
	name = game.Name
	switch gid {
	case cst.ChineseChess:
		handel = chinesechess.New(game, tb)
	case cst.Chess:
	case cst.SanGuoXiao:
		handel = sanguoxiao.New(game)
	case cst.Baccarat:
		handel = brbaccarat.New(game, tb)
	case cst.BrCowcow:
		handel = brcowcow.New(game, tb)
	case cst.TigerXdragon:
		handel = brtigerXdragon.New(game, tb)
	case cst.BrToubao:
		handel = brtoubao.New(game, tb)
	case cst.BrTuitongzi:
		handel = brtuitongzi.New(game, tb)
	default:
	}
	return
}
