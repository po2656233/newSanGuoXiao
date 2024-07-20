package sanguoxiao

import (
	"github.com/po2656233/goleaf/log"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/category"
	"time"
)

const (
	rowNum = 9
	colNum = 9
	wait   = 15
)

// PieceDamage 伤害列表
var PieceDamage = map[protoMsg.SgxPiece]int64{
	protoMsg.SgxPiece_Jin:  10,
	protoMsg.SgxPiece_Mu:   10,
	protoMsg.SgxPiece_Shui: 10,
	protoMsg.SgxPiece_Huo:  10,
	protoMsg.SgxPiece_Tu:   10,
	protoMsg.SgxPiece_Yao:  100,
}

// SanguoxiaoGame 继承于GameItem
type SanguoxiaoGame struct {
	*mgr.Game
	row        int32
	col        int32
	curUID     int64
	status     protoMsg.GameScene
	curTimer   *time.Timer
	board      *protoMsg.BoardInfo   // 其实质是二维数组
	arrayBoard [][]protoMsg.SgxPiece // 和board等同,游戏主要操作该棋盘
	result     *protoMsg.SanguoxiaoResult

	redPlay  *protoMsg.SanguoxiaoPlayer // 红方玩家
	bluePlay *protoMsg.SanguoxiaoPlayer // 蓝方玩家
	playList []int64
}

// New 创建实例
func New(game *mgr.Game) *SanguoxiaoGame {
	p := &SanguoxiaoGame{
		Game: game,
	}
	category.InitConfig()
	p.row = category.YamlObj.SanGuoXiao.Row
	p.col = category.YamlObj.SanGuoXiao.Col
	p.Init()
	return p
}

// Init 初始化信息
func (self *SanguoxiaoGame) Init() {
	self.initBord(self.row, self.col)
}

// Scene 场 景
func (self *SanguoxiaoGame) Scene(args []interface{}) {
	//level := args[0].(int32)
	self.initBord(self.row, self.col)
	agent := args[0].(mgr.Agent)
	userData := agent.UserData()
	person := userData.(*mgr.Player)
	now := time.Now().Unix()
	mgr.GetClientMgr().SendTo(person.UserID,
		&protoMsg.SanguoxiaoSceneResp{
			TimeStamp: time.Now().Unix(),
			//Inning:    self.InningInfo.Number,
			Board:    self.board,
			RedCamp:  self.redPlay,
			BlueCamp: self.bluePlay,
		})
	switch self.status {
	case protoMsg.GameScene_Start:
	case protoMsg.GameScene_Playing: // 移位
		{
			msg := &protoMsg.SanguoxiaoStatePlayingResp{
				UserID: self.curUID,
				Times: &protoMsg.TimeInfo{
					TimeStamp: now,
					//WaitTime:  self.Duration.PlayTime,
					OutTime: int32(now - self.TimeStamp),
					//TotalTime: self.Duration.PlayTime,
				},
			}
			mgr.GetClientMgr().SendTo(person.UserID, msg)
		}
	case protoMsg.GameScene_Opening: // 消除
		{
			msg := &protoMsg.SanguoxiaoStateEraseResp{
				NowBoard: self.board,
				Times: &protoMsg.TimeInfo{
					TimeStamp: now,
					//WaitTime:  self.Config.SanGuoXiao.Duration.OpenTime,
					OutTime: int32(now - self.TimeStamp),
					//TotalTime: self.Config.SanGuoXiao.Duration.OpenTime,
				},
			}
			mgr.GetClientMgr().SendTo(person.UserID, msg)
		}
	case protoMsg.GameScene_Over: // 结算
		{
			msg := &protoMsg.SanguoxiaoStateOverResp{
				Result: self.result,
				Times: &protoMsg.TimeInfo{
					TimeStamp: now,
					//WaitTime:  self.Config.SanGuoXiao.Duration.OverTime,
					OutTime: int32(now - self.TimeStamp),
					//TotalTime: self.Config.SanGuoXiao.Duration.OverTime,
				},
			}
			mgr.GetClientMgr().SendTo(person.UserID, msg)
		}

	}

}

// Start 开 始 检测玩家是否
func (self *SanguoxiaoGame) Start(args []interface{}) {

}

// Playing 游 戏
func (self *SanguoxiaoGame) Playing(args []interface{}) {
	//
	m := args[0].(*protoMsg.SanguoxiaoSwapReq)
	agent := args[1].(mgr.Agent)
	if protoMsg.GameScene_Playing != self.status {
		log.Error("[Sanguoxiao][Playing] 场景信息不匹配 %v ", self.status)
		return
	}
	userData := agent.UserData()
	if userData == nil { //[0
		//mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game37])
		log.Error("[Sanguoxiao][Playing] 用户数据 %v ", self.status)
		return
	}
	person := userData.(*mgr.Player)
	cells := make([]*protoMsg.Grid, len(self.board.Cells))
	for _, item := range self.board.Cells {
		cells = append(cells, &protoMsg.Grid{
			Row:  item.Row,
			Col:  item.Col,
			Core: item.Core,
		})
	}
	// 交换位置上的数值
	if !self.swapGrid(m.Origin, m.Target) {
		//mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game37])
		log.Error("[Sanguoxiao][Playing] 用户操作错误 %v ", self.status)
		return
	}

	// 通知玩家 当前操作
	mgr.GetClientMgr().NotifyOthers(self.playList, &protoMsg.SanguoxiaoSwapResp{
		UserID: person.UserID,
		Origin: m.Origin,
		Target: m.Target,
	})

	// 轮值至下个玩家
	self.status = protoMsg.GameScene_Opening
	self.OnOpen()
}

// Over 结 算
func (self *SanguoxiaoGame) Over(args []interface{}) {
	if protoMsg.GameScene_Over != self.status {
		log.Error("[Sanguoxiao][Over] 场景信息不匹配 %v ", self.status)
		return
	}

	// 玩家1 胜
	winId := self.redPlay.Info.UserID
	loseId := self.bluePlay.Info.UserID
	if self.redPlay.Health < self.bluePlay.Health {
		// 玩家2 胜
		winId = self.bluePlay.Info.UserID
		loseId = self.redPlay.Info.UserID
	}

	// 奖励  todo 待商榷
	award := make([]*protoMsg.GoodsInfo, 0)
	award = append(award, &protoMsg.GoodsInfo{
		Id:   1,
		Sold: 1,
	})
	self.result = &protoMsg.SanguoxiaoResult{
		WinID:  winId,
		LoseID: loseId,
		Awards: &protoMsg.GoodsList{
			Items: award,
		},
	}
	// 取消选中状态
	heroIds := make([]int64, 0)
	for _, info := range self.redPlay.ArenaTeam {
		heroIds = append(heroIds, info.Id)
	}
	//mysql.SqlHandle().ChooseHeros(self.redPlay.Info.UserID, heroIds, false)
	// 取消选中状态
	heroIds = make([]int64, 0)
	for _, info := range self.bluePlay.ArenaTeam {
		heroIds = append(heroIds, info.Id)
	}
	//mysql.SqlHandle().ChooseHeros(self.bluePlay.Info.UserID, heroIds, false)
}

// UpdateInfo 更新信息
func (self *SanguoxiaoGame) UpdateInfo(args []interface{}) bool {

	return true
}

// SuperControl 超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭
func (self *SanguoxiaoGame) SuperControl(args []interface{}) bool {
	return true
}

// //////////////////////////////////////////////////////////////

// OnNext 轮换下一位
func (self *SanguoxiaoGame) OnNext() {
	if self.curTimer != nil {
		self.curTimer.Stop()
	}
	if self.status != protoMsg.GameScene_Playing {
		log.Error("[Sanguoxiao][OnNext] 场景信息不匹配 %v ", self.status)
		return
	}

	self.TimeStamp = time.Now().Unix()
	if self.redPlay.Info.UserID == self.curUID {
		self.curUID = self.bluePlay.Info.UserID
	} else if self.bluePlay.Info.UserID == self.curUID {
		self.curUID = self.redPlay.Info.UserID
	}

	mgr.GetClientMgr().NotifyOthers(self.playList,
		&protoMsg.SanguoxiaoStatePlayingResp{
			Times: &protoMsg.TimeInfo{
				TimeStamp: self.TimeStamp,
				WaitTime:  wait,
				OutTime:   0,
				TotalTime: wait,
			},
			UserID: self.curUID,
		})
	self.curTimer = time.AfterFunc(time.Duration(wait)*time.Second, self.OnOpen)
}

// OnOpen 消除阶段
func (self *SanguoxiaoGame) OnOpen() {
	if self.curTimer != nil {
		self.curTimer.Stop()
	}
	if self.status != protoMsg.GameScene_Opening {
		log.Error("[Sanguoxiao][OnNext] 场景信息不匹配 %v ", self.status)
		return
	}

	// 消除
	if !self.eraseGrid() {
		// 无法消除,则轮置下位
		self.toState(protoMsg.GameScene_Playing, self.OnNext)
		return
	}

	// 检测双方血量
	if self.redPlay.Health <= 0 || self.bluePlay.Health <= 0 {
		self.toState(protoMsg.GameScene_Over, self.OnOver)
		return
	}
	self.TimeStamp = time.Now().Unix()
	mgr.GetClientMgr().NotifyOthers(self.playList,
		&protoMsg.SanguoxiaoStateEraseResp{
			Times: &protoMsg.TimeInfo{
				TimeStamp: self.TimeStamp,
				//WaitTime:  self.Config.SanGuoXiao.Duration.OpenTime,
				OutTime: 0,
				//TotalTime: self.Config.SanGuoXiao.Duration.OpenTime,
			},
			NowBoard: self.board,
		})
	//self.curTimer = time.AfterFunc(time.Duration(self.Config.SanGuoXiao.Duration.OpenTime)*time.Second, self.OnOpen)
}

// OnOver 结算阶段
func (self *SanguoxiaoGame) OnOver() {
	if self.curTimer != nil {
		self.curTimer.Stop()
	}
	if self.status != protoMsg.GameScene_Over {
		log.Error("[Sanguoxiao][OnNext] 场景信息不匹配 %v ", self.status)
		return
	}

	// 结算
	self.Over(nil)

	self.TimeStamp = time.Now().Unix()
	msg := &protoMsg.SanguoxiaoStateOverResp{
		Result: self.result,
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			//WaitTime:  self.Config.SanGuoXiao.Duration.OverTime,
			OutTime: 0,
			//TotalTime: self.Config.SanGuoXiao.Duration.OverTime,
		},
	}
	mgr.GetClientMgr().NotifyOthers(self.playList, msg)
	//self.curTimer = time.AfterFunc(time.Duration(self.Config.SanGuoXiao.Duration.OverTime)*time.Second, func() {
	//	self = nil
	//})
}

// /////////////////////////////////////////////////////////////////////
func (self *SanguoxiaoGame) toState(scene protoMsg.GameScene, f func()) {
	self.status = scene
	f()
}

////////////////////////////////////////////////////////////////////////

// initBord 初始化方格
func (self *SanguoxiaoGame) initBord(row, col int32) {
	self.board = &protoMsg.BoardInfo{}
	self.arrayBoard = make([][]protoMsg.SgxPiece, row) // [rowNum][colNum]protoMsg.SgxPiece{}
	for i := 0; i < int(row); i++ {
		self.arrayBoard[i] = make([]protoMsg.SgxPiece, col)
	}
	// 避免连续三个相同
	for i := int32(0); i < row; i++ {
		for j := int32(0); j < col; j++ {
			item := &protoMsg.Grid{
				Row:  i,
				Col:  j,
				Core: protoMsg.SgxPiece(utils.GenRandNum32(int32(protoMsg.SgxPiece_Jin), int32(protoMsg.SgxPiece_Yao))),
			}
			if self.canErase(item) {
				for index := protoMsg.SgxPiece_Jin; index <= protoMsg.SgxPiece_Yao; index++ {
					if item.Core == index {
						continue
					}
					item.Core = index
					if !self.canErase(item) {
						break
					}
				}
			}
			self.arrayBoard[i][j] = item.Core
			self.board.Cells = append(self.board.Cells, item)
		}
	}

}

// 获取格子
func (self *SanguoxiaoGame) getGrid(row, col int32) *protoMsg.Grid {
	for _, cell := range self.board.Cells {
		if cell.Row == row && cell.Col == col {
			return cell
		}
	}
	return nil
}

// 设置格子
func (self *SanguoxiaoGame) setGrid(row, col int32, core protoMsg.SgxPiece) {
	if cell := self.getGrid(row, col); cell != nil {
		cell.Core = core
	}
}

// swapGrid 方格属性交换
func (self *SanguoxiaoGame) swapGrid(origin, target *protoMsg.Grid) bool {
	divR := origin.Row - target.Row
	divC := origin.Col - target.Col
	// 必须是相连的方格
	if (1 < divC || divC < -1) && divR != 0 {
		return false
	}
	if (1 < divR || divR < -1) && divC != 0 {
		return false
	}
	tempCell := &protoMsg.Grid{
		Row:  origin.Row,
		Col:  origin.Col,
		Core: target.Core,
	}
	// 交换位置数值
	self.arrayBoard[origin.Row][origin.Col] = target.Core
	self.arrayBoard[target.Row][target.Col] = origin.Core
	if !self.canErase(tempCell) {
		tempCell.Row = target.Row
		tempCell.Col = target.Col
		tempCell.Core = origin.Core
		if !self.canErase(tempCell) {
			// 还原数值
			self.arrayBoard[origin.Row][origin.Col] = origin.Core
			self.arrayBoard[target.Row][target.Col] = target.Core
			return false
		}
	}

	// 交换坐标上的属性
	for _, cell := range self.board.Cells {
		if cell.Row == origin.Row && cell.Col == origin.Col {
			cell.Core = target.Core
			continue
		}
		if cell.Row == target.Row && cell.Col == target.Col {
			cell.Core = origin.Core
		}
	}
	return true

}

// /////////////////////////////////////////////////////////////////////////////////////////
// eraseGrid 消除方格 获得的战力 主要逻辑
func (self *SanguoxiaoGame) eraseGrid() bool {
	msg := &protoMsg.SanguoxiaoTriggerResp{
		UserID: self.curUID,
		PerBoard: &protoMsg.BoardInfo{
			Cells: make([]*protoMsg.Grid, 0),
		},
		Erase:   make([]*protoMsg.Grid, 0),
		Attacks: make([]*protoMsg.SanguoxiaoAttack, 0),
	}

	/////////////////////////游戏核心玩法////////////////////////////////
	// 消除 逐个消除
	for _, cell := range self.board.Cells {
		temp := &protoMsg.Grid{
			Row:  cell.Row,
			Col:  cell.Col,
			Core: cell.Core,
		}
		msg.PerBoard.Cells = append(msg.PerBoard.Cells, temp)
		if self.canErase(cell) {
			msg.Erase = append(msg.Erase, temp)
		}
	}

	// 根据消除的元素数量统计攻击伤害
	for _, grid := range msg.Erase {
		if cell := self.getGrid(grid.Row, grid.Col); cell != nil {
			cell.Core = protoMsg.SgxPiece_NoPiece
		}
		self.arrayBoard[grid.Row][grid.Col] = protoMsg.SgxPiece_NoPiece
	}

	// 掉落
	self.fallDown()

	// 填充空缺
	self.fillEmpty()
	/////////////////////////////////////////////////////////////////////

	// 获取玩家伤害或治疗 统计玩家当前生命值
	myHealth := self.redPlay.Health
	BlueCampHealth := self.bluePlay.Health
	heros := self.redPlay.ArenaTeam
	if self.curUID == self.bluePlay.Info.UserID {
		myHealth = self.bluePlay.Health
		BlueCampHealth = self.redPlay.Health
		heros = self.bluePlay.ArenaTeam
	}
	allDamage := int64(0)
	allTherapy := int64(0)
	msg.Attacks, allDamage, allTherapy = self.getAttacks(msg.Erase, heros)
	msg.RedCampHealth = myHealth + allTherapy
	msg.BlueCampHealth = BlueCampHealth - allDamage
	if msg.BlueCampHealth < 0 {
		msg.BlueCampHealth = 0
	}

	msg.NowBoard = self.board
	mgr.GetClientMgr().NotifyOthers(self.playList, msg)

	// 当前各玩家生命值
	self.redPlay.Health = msg.RedCampHealth
	self.bluePlay.Health = msg.BlueCampHealth
	if self.curUID == self.bluePlay.Info.UserID {
		self.redPlay.Health = msg.BlueCampHealth
		self.bluePlay.Health = msg.RedCampHealth
	}
	return 0 < len(msg.Erase)
}

// 掉落
func (self *SanguoxiaoGame) fallDown() {
	//// 按列掉落
	//for c := int32(0); c < self.col; c++ {
	//	for r := int32(0); r < self.row; r++ {
	//		// 找到空位
	//		if self.arrayBoard[r][c] == protoMsg.SgxPiece_NoPiece {
	//			for nr := r + 1; nr < self.row; nr++ {
	//				// 找到可以用的方块
	//				if self.arrayBoard[nr][c] != protoMsg.SgxPiece_NoPiece {
	//					// 转移数据
	//					self.setGrid(r, c, self.arrayBoard[nr][c])
	//					self.arrayBoard[r][c] = self.arrayBoard[nr][c]
	//					self.arrayBoard[nr][c] = protoMsg.SgxPiece_NoPiece
	//					// [c,r]下落 nr-r 格
	//					break
	//				}
	//			}
	//		}
	//	}
	//}
}

// 填充
func (self *SanguoxiaoGame) fillEmpty() {
	//for r := int32(0); r < self.row; r++ {
	//	for c := int32(0); c < self.col; c++ {
	//		if self.arrayBoard[r][c] == protoMsg.SgxPiece_NoPiece {
	//			self.arrayBoard[r][c] = protoMsg.SgxPiece(GenRandNum32(int32(protoMsg.SgxPiece_Jin), int32(protoMsg.SgxPiece_Yao)))
	//			self.setGrid(r, c, self.arrayBoard[r][c])
	//		}
	//	}
	//}
}

// 获取伤害或治疗
func (self *SanguoxiaoGame) getAttacks(cells []*protoMsg.Grid, heros []*protoMsg.HeroInfo) ([]*protoMsg.SanguoxiaoAttack, int64, int64) {
	attackList := make([]*protoMsg.SanguoxiaoAttack, 0)
	therapyTotal := int64(0)
	damageTotal := int64(0)
	for _, hero := range heros {
		for _, cell := range cells {
			if int32(hero.Faction) == int32(cell.Core) {
				attack := &protoMsg.SanguoxiaoAttack{
					HeroId: hero.Id,
				}
				if cell.Core == protoMsg.SgxPiece_Yao {
					therapyTotal += PieceDamage[cell.Core]
					attack.Therapy += PieceDamage[cell.Core]
				} else {
					damageTotal += PieceDamage[cell.Core]
					attack.Damage += PieceDamage[cell.Core]
				}
				attackList = append(attackList, attack)
			}
		}
	}

	return attackList, damageTotal, therapyTotal

}

// //////////////////////////////////////////////////////////////////////////
// 判断格子前后的是否可消除
func (self *SanguoxiaoGame) canErase(grid *protoMsg.Grid) bool {
	if grid.Core == protoMsg.SgxPiece_NoPiece || 0 == len(self.arrayBoard) {
		return false
	}
	gridUp1 := &protoMsg.Grid{
		Row: grid.Row - 1,
		Col: grid.Col,
	}
	gridUp2 := &protoMsg.Grid{
		Row: grid.Row - 2,
		Col: grid.Col,
	}
	gridDw1 := &protoMsg.Grid{
		Row: grid.Row + 1,
		Col: grid.Col,
	}
	gridDw2 := &protoMsg.Grid{
		Row: grid.Row + 2,
		Col: grid.Col,
	}
	gridLf1 := &protoMsg.Grid{
		Row: grid.Row,
		Col: grid.Col - 1,
	}
	gridLf2 := &protoMsg.Grid{
		Row: grid.Row,
		Col: grid.Col - 2,
	}
	gridRt1 := &protoMsg.Grid{
		Row: grid.Row,
		Col: grid.Col + 1,
	}
	gridRt2 := &protoMsg.Grid{
		Row: grid.Row,
		Col: grid.Col + 2,
	}
	// 获取相连的两格
	// 上方 前两格
	if 0 <= gridUp1.Row && gridUp1.Row < self.row &&
		0 <= gridUp1.Col && gridUp1.Col < self.col {
		gridUp1.Core = self.arrayBoard[gridUp1.Row][gridUp1.Col]
	}
	if 0 <= gridUp2.Row && gridUp2.Row < self.row &&
		0 <= gridUp2.Col && gridUp2.Col < self.col {
		gridUp2.Core = self.arrayBoard[gridUp2.Row][gridUp2.Col]
	}
	// 下方 前两格
	if 0 <= gridDw1.Row && gridDw1.Row < self.row &&
		0 <= gridDw1.Col && gridDw1.Col < self.col {
		gridDw1.Core = self.arrayBoard[gridDw1.Row][gridDw1.Col]
	}
	if 0 <= gridDw2.Row && gridDw2.Row < self.row &&
		0 <= gridDw2.Col && gridDw2.Col < self.col {
		gridDw2.Core = self.arrayBoard[gridDw2.Row][gridDw2.Col]
	}
	// 左方 前两格
	if 0 <= gridLf1.Row && gridLf1.Row < self.row &&
		0 <= gridLf1.Col && gridLf1.Col < self.col {
		gridLf1.Core = self.arrayBoard[gridLf1.Row][gridLf1.Col]
	}
	if 0 <= gridLf2.Row && gridLf2.Row < self.row &&
		0 <= gridLf2.Col && gridLf2.Col < self.col {
		gridLf2.Core = self.arrayBoard[gridLf2.Row][gridLf2.Col]
	}
	// 右方 前两格
	if 0 <= gridRt1.Row && gridRt1.Row < self.row &&
		0 <= gridRt1.Col && gridRt1.Col < self.col {
		gridRt1.Core = self.arrayBoard[gridRt1.Row][gridRt1.Col]
	}
	if 0 <= gridRt2.Row && gridRt2.Row < self.row &&
		0 <= gridRt2.Col && gridRt2.Col < self.col {
		gridRt2.Core = self.arrayBoard[gridRt2.Row][gridRt2.Col]
	}
	return (gridUp1.Core == grid.Core && gridUp2.Core == grid.Core) || //上消除
		(gridDw1.Core == grid.Core && gridDw2.Core == grid.Core) || //下消除
		(gridUp1.Core == grid.Core && gridDw1.Core == grid.Core) || //上下消除
		(gridLf1.Core == grid.Core && gridLf2.Core == grid.Core) || //左消除
		(gridRt1.Core == grid.Core && gridRt2.Core == grid.Core) || //右消除
		(gridLf1.Core == grid.Core && gridRt1.Core == grid.Core) //左右消除
}
