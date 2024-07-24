package chinesechess

import (
	log "github.com/po2656233/superplace/logger"
	protoMsg "superman/internal/protocol/gofile"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/category"
	"time"
)

const (
	RowCount = 10
	ColCount = 9
)

type ChineseChess struct {
	*mgr.Game
	board     *protoMsg.XQBoardInfo
	redCamp   *protoMsg.PlayerSimpleInfo
	blackCamp *protoMsg.PlayerSimpleInfo
	curUid    int64
	winUid    int64
}

// New 棋盘 ChineseChess
func New(game *mgr.Game) *ChineseChess {
	p := &ChineseChess{
		Game: game,
		board: &protoMsg.XQBoardInfo{
			Cells: make([]*protoMsg.XQGrid, 0),
		},
	}
	p.Init()
	return p
}

// Init 初始化信息
func (self *ChineseChess) Init() {
	self.Reset()
	self.initBord()
}

// Reset 重置
func (self *ChineseChess) Reset() {
	self.Game.Reset()
	self.winUid = 0
}

// Scene 场景
func (self *ChineseChess) Scene(args []interface{}) {
	agent := args[0].(mgr.Agent)
	userData := agent.UserData()
	person := userData.(*mgr.Player)
	if person == nil {
		log.Warnf("[%v:%v][Scene:%v] person is nil.", self.Name, self.Tid, self.Game.Scene)
		return
	}
	//
	mgr.GetClientMgr().SendTo(person.UserID, &protoMsg.ChineseChessSceneResp{
		TimeStamp: time.Now().Unix(),
		Inning:    self.Inning,
		Board:     self.board,
		RedCamp:   self.redCamp,
		BlackCamp: self.blackCamp,
	})
	// 获取时间信息
	var getTimeInfo = func(waitTime int32) *protoMsg.TimeInfo {
		t := &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			TotalTime: waitTime,
		}
		t.OutTime = int32(time.Now().Unix() - t.TimeStamp)
		t.WaitTime = t.TotalTime - t.OutTime
		if t.WaitTime < 0 {
			t.WaitTime = 0
		}
		return t
	}
	switch self.Game.Scene {
	case protoMsg.GameScene_Free: // 等待新玩家
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateFreeResp{
			Times: getTimeInfo(category.YamlObj.ChineseChess.FreeTime),
		})
	case protoMsg.GameScene_Start: // 开始
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateStartResp{
			Times: getTimeInfo(category.YamlObj.ChineseChess.StartTime),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Playing: // 下棋
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStatePlayingResp{
			Times: getTimeInfo(category.YamlObj.ChineseChess.PlayTime),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Opening: // 开奖
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateOpenResp{
			Times:  getTimeInfo(category.YamlObj.ChineseChess.OpenTime),
			WinUid: self.winUid,
		})
	case protoMsg.GameScene_Over: // 结算
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateOverResp{
			Times: getTimeInfo(category.YamlObj.ChineseChess.OverTime),
			Result: &protoMsg.ChineseChessResult{
				RedCamp:   self.redCamp,
				BlackCamp: self.blackCamp,
			},
		})
	default:
		log.Warnf("[%v:%v][Scene:%v] person:%v no have scence.", self.Name, self.Tid, self.Game.Scene, person.UserID)
	}
}

// Ready 准备
func (self *ChineseChess) Ready(args []interface{}) {
	person, ok := args[0].(*mgr.Player)
	if !ok || person == nil {
		return
	}
	simpleInfo := &protoMsg.PlayerSimpleInfo{
		Uid:     person.UserID,
		HeadId:  person.FaceID,
		ChairId: person.InChairId,
		Score:   person.Gold,
		RankNo:  person.Ranking,
		Name:    person.Name,
	}
	if self.redCamp == nil {
		self.redCamp = simpleInfo
	} else if self.blackCamp == nil {
		self.blackCamp = simpleInfo
	}
}

// Start 开始
func (self *ChineseChess) Start(args []interface{}) {
	if self.blackCamp == nil || self.redCamp == nil {
		return
	}
}

// Playing 结 算
func (self *ChineseChess) Playing(args []interface{}) {
	//

}

// Over 结 算
func (self *ChineseChess) Over(args []interface{}) {

}

// UpdateInfo 更新
func (self *ChineseChess) UpdateInfo(args []interface{}) bool {
	return true

}

// SuperControl 超级控制
func (self *ChineseChess) SuperControl(args []interface{}) bool {

	return true
}

// /////////////////////////[定时器事件]//////////////////////////////////////////
func (self *ChineseChess) onStart() {
	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(category.YamlObj.ChineseChess.Duration.StartTime)*time.Second, self.onSetTime)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BaccaratStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Baccarat.Duration.StartTime,
			TotalTime: self.Config.Baccarat.Duration.StartTime,
		},
		Inning: self.InningInfo.Number,
	})
}

func (self *ChineseChess) onSetTime() {
	self.ChangeState(protoMsg.GameScene_SetTing)
}

func (self *ChineseChess) onPlay() {
	self.ChangeState(protoMsg.GameScene_Start)
}

// ///////////////////////[初始化]///////////////////////////////////
// 初始化棋盘
func (self *ChineseChess) initBord() {
	self.board = &protoMsg.XQBoardInfo{
		Cells: make([]*protoMsg.XQGrid, 0),
	}
	rowHalf := RowCount / 2
	for row := 0; row < rowHalf; row++ {
		if row == 0 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(0),
				Core: protoMsg.XQPiece_BlackJu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(1),
				Core: protoMsg.XQPiece_BlackMa,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(2),
				Core: protoMsg.XQPiece_BlackXiang,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(3),
				Core: protoMsg.XQPiece_BlackShi,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(4),
				Core: protoMsg.XQPiece_BlackJiang,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(5),
				Core: protoMsg.XQPiece_BlackShi,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(6),
				Core: protoMsg.XQPiece_BlackXiang,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(7),
				Core: protoMsg.XQPiece_BlackMa,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(8),
				Core: protoMsg.XQPiece_BlackJu,
			})
		}
		if row == 2 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(1),
				Core: protoMsg.XQPiece_BlackPao,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(7),
				Core: protoMsg.XQPiece_BlackPao,
			})
		}
		if row == 3 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(0),
				Core: protoMsg.XQPiece_BlackZu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(2),
				Core: protoMsg.XQPiece_BlackZu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(4),
				Core: protoMsg.XQPiece_BlackZu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(6),
				Core: protoMsg.XQPiece_BlackZu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(8),
				Core: protoMsg.XQPiece_BlackZu,
			})
		}
	}
	for row := RowCount - 1; row >= rowHalf; row-- {
		if row == RowCount-1 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(0),
				Core: protoMsg.XQPiece_RedJu,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(1),
				Core: protoMsg.XQPiece_RedMa,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(2),
				Core: protoMsg.XQPiece_RedXiang,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(3),
				Core: protoMsg.XQPiece_RedShi,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(4),
				Core: protoMsg.XQPiece_RedShuai,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(5),
				Core: protoMsg.XQPiece_RedShi,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(6),
				Core: protoMsg.XQPiece_RedXiang,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(7),
				Core: protoMsg.XQPiece_RedMa,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(8),
				Core: protoMsg.XQPiece_RedJu,
			})
		}
		if row == RowCount-3 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(1),
				Core: protoMsg.XQPiece_RedPao,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(7),
				Core: protoMsg.XQPiece_RedPao,
			})
		}
		if row == RowCount-4 {
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(0),
				Core: protoMsg.XQPiece_RedBing,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(2),
				Core: protoMsg.XQPiece_RedBing,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(4),
				Core: protoMsg.XQPiece_RedBing,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(6),
				Core: protoMsg.XQPiece_RedBing,
			})
			self.board.Cells = append(self.board.Cells, &protoMsg.XQGrid{
				Row:  int32(row),
				Col:  int32(8),
				Core: protoMsg.XQPiece_RedBing,
			})
		}
	}

}
