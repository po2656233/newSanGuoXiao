package tetris

import (
	protoMsg "superman/internal/protocol/gofile"
	mgr "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
)

// Tetris 俄罗斯方块
type Tetris struct {
	*mgr.Game
	T         *mgr.Table
	board     *protoMsg.XQBoardInfo
	redCamp   *protoMsg.PlayerSimpleInfo
	blackCamp *protoMsg.PlayerSimpleInfo
	curUid    int64
	winUid    int64
	timeout   int32
}

// New 棋盘 Chinesechess
func New(game *mgr.Game, table *mgr.Table) *Tetris {
	p := &Tetris{
		Game: game,
		T:    table,
		board: &protoMsg.XQBoardInfo{
			Cells: make([]*protoMsg.XQGrid, 0),
		},
	}
	p.Init()
	return p
}

// Init 重置
func (self *Tetris) Init() {
	self.Game.Init()
	self.winUid = 0
	self.timeout = YamlObj.Chinesechess.Duration.Play
}
