package chinesechess

import (
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	mgr "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
	"time"
)

const (
	RowCount = 10
	ColCount = 9
)

type Chinesechess struct {
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
func New(game *mgr.Game, table *mgr.Table) *Chinesechess {
	p := &Chinesechess{
		Game: game,
		T:    table,
		board: &protoMsg.XQBoardInfo{
			Cells: make([]*protoMsg.XQGrid, 0),
		},
	}
	p.Reset()
	return p
}

// Reset 重置
func (self *Chinesechess) Reset() {
	self.Game.Reset()
	self.winUid = 0
	self.initBord()
	self.timeout = YamlObj.Chinesechess.Duration.Play
}

// Scene 场景
func (self *Chinesechess) Scene(args []interface{}) {
	agent := args[0].(mgr.Agent)
	userData := agent.UserData()
	person := userData.(*mgr.Player)
	if person == nil {
		log.Warnf("[%v:%v][Scene:%v] person is nil.", self.Name, self.T.Id, self.GameInfo.Scene)
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
	switch self.GameInfo.Scene {
	case protoMsg.GameScene_Setting: //
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateSetResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Settime),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_WaitOperate: //
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateConfirmResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Confirm),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Start: // 开始
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateStartResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Start),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Playing: // 下棋
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStatePlayingResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Play),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Opening: // 开奖
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateOpenResp{
			Times:  getTimeInfo(YamlObj.Chinesechess.Duration.Open),
			WinUid: self.winUid,
		})
	case protoMsg.GameScene_Over: // 结算
		mgr.GetClientMgr().SendData(agent, &protoMsg.ChineseChessStateOverResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Over),
			Result: &protoMsg.ChineseChessResult{
				RedCamp:   self.redCamp,
				BlackCamp: self.blackCamp,
			},
		})
	default:
		log.Warnf("[%v:%v][Scene:%v] person:%v no have scence.", self.Name, self.T.Id, self.GameInfo.Scene, person.UserID)
	}
}

// Ready 准备
func (self *Chinesechess) Ready(args []interface{}) {
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
	} else {
		self.onSetTime()
	}

}

// Start 开始
func (self *Chinesechess) Start(args []interface{}) {
	self.Game.Start(args)
	if self.blackCamp == nil || self.redCamp == nil {
		return
	}
	// 开始设置下棋时长
	self.onSetTime()
}

// Playing 结 算
func (self *Chinesechess) Playing(args []interface{}) {
	//
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.ChineseChessMoveReq)
	agent := args[1].(mgr.Agent)
	if self.GameInfo.Scene != protoMsg.GameScene_Playing {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game04])
		return
	}

	person := agent.UserData().(*mgr.Player)
	msg := &protoMsg.ChineseChessMoveResp{
		Uid:    person.UserID,
		Origin: m.Origin,
		Target: m.Target,
	}
	mgr.GetClientMgr().NotifyOthers(self.T.GetPlayList(), msg)
}

// Over 结 算
func (self *Chinesechess) Over(args []interface{}) {

}

// UpdateInfo 更新玩家信息
func (self *Chinesechess) UpdateInfo(args []interface{}) bool {
	flag, ok := args[0].(protoMsg.PlayerState)
	if !ok {
		return false
	}
	switch flag {
	case protoMsg.PlayerState_PlayerSetTime:
		if 2 < len(args) {
			self.SetTimeOP(args[2:])
		}
	case protoMsg.PlayerState_PlayerAgree:
		if 2 < len(args) {
			self.ConfirmOP(args[2:])
		}
	default:
		return self.Game.UpdateInfo(args)
	}
	return true
}

// ReadyOP 准备
func (self *Chinesechess) ReadyOP(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.ChineseChessReadyReq)
	agent := args[1].(mgr.Agent)
	if protoMsg.GameScene_Start <= self.GameInfo.Scene {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game04])
		return
	}

	person := agent.UserData().(*mgr.Player)
	msg := &protoMsg.ChineseChessReadyResp{
		Uid:     person.UserID,
		IsReady: m.IsReady,
	}
	mgr.GetClientMgr().NotifyOthers(self.T.GetPlayList(), msg)
}

// SetTimeOP 设置时长
func (self *Chinesechess) SetTimeOP(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.ChineseChessSetTimeReq)
	agent := args[1].(mgr.Agent)
	if self.GameInfo.Scene != protoMsg.GameScene_Setting {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game04])
		return
	}
	if maxTimeout < m.Timeout {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game61])
		return
	}
	person := agent.UserData().(*mgr.Player)
	self.timeout = m.Timeout
	msg := &protoMsg.ChineseChessSetTimeResp{
		Uid:     person.UserID,
		Timeout: m.Timeout,
	}
	mgr.GetClientMgr().NotifyOthers(self.T.GetPlayList(), msg)
}

// ConfirmOP 确认时长
func (self *Chinesechess) ConfirmOP(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.ChineseChessAgreeTimeReq)
	agent := args[1].(mgr.Agent)
	if self.GameInfo.Scene != protoMsg.GameScene_WaitOperate {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game04])
		return
	}

	person := agent.UserData().(*mgr.Player)
	msg := &protoMsg.ChineseChessAgreeTimeResp{
		Uid:     person.UserID,
		IsAgree: m.IsAgree,
	}
	mgr.GetClientMgr().NotifyOthers(self.T.GetPlayList(), msg)
}

// /////////////////////////[定时器事件]//////////////////////////////////////////

func (self *Chinesechess) onReady() {
	self.ChangeState(protoMsg.GameScene_Ready)
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Ready)*time.Second, self.onSetTime)
	list := self.T.GetPlayList()
	if self.redCamp != nil {
		list = utils.RemoveValue(list, self.redCamp.Uid)
	}
	if self.blackCamp != nil {
		list = utils.RemoveValue(list, self.blackCamp.Uid)
	}
	GlobalSender.NotifyOthers(list, &protoMsg.ChineseChessStateReadyResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Ready,
			TotalTime: YamlObj.Chinesechess.Duration.Ready,
		},
	})
}

func (self *Chinesechess) onSetTime() {
	self.ChangeState(protoMsg.GameScene_Setting)
	if self.blackCamp == nil || self.redCamp == nil {
		self.onReady()
		return
	}
	self.curUid = self.redCamp.Uid
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Settime)*time.Second, self.onConfirmTime)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStateSetResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Settime,
			TotalTime: YamlObj.Chinesechess.Duration.Settime,
		},
		Uid: self.curUid,
	})
}

func (self *Chinesechess) onConfirmTime() {
	self.ChangeState(protoMsg.GameScene_WaitOperate)
	self.curUid = self.blackCamp.Uid
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Confirm)*time.Second, self.onStart)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStateConfirmResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Confirm,
			TotalTime: YamlObj.Chinesechess.Duration.Confirm,
		},
		Uid: self.curUid,
	})
}
func (self *Chinesechess) onStart() {
	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Start)*time.Second, self.onPlay)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Start,
			TotalTime: YamlObj.Chinesechess.Duration.Start,
		},
		Uid: self.curUid,
	})

}

func (self *Chinesechess) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	self.curUid = self.T.NextChairUID(self.curUid)
	if self.curUid == INVALID {
		log.Errorf("出现严重错误,当前玩家ID %v 出错", self.curUid)
		return
	}
	if self.winUid != INVALID {
		self.onOpen()
		return
	}
	self.Timer = time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Play)*time.Second, self.onPlay)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Play,
			TotalTime: YamlObj.Chinesechess.Duration.Play,
		},
		Uid:      self.curUid,
		NowBoard: self.board,
	})
}

func (self *Chinesechess) onOpen() {
	self.ChangeState(protoMsg.GameScene_Opening)
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Open)*time.Second, self.onOver)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Open,
			TotalTime: YamlObj.Chinesechess.Duration.Open,
		},
		WinUid:   self.winUid,
		NowBoard: self.board,
	})
}
func (self *Chinesechess) onOver() {
	self.ChangeState(protoMsg.GameScene_Over)
	time.AfterFunc(time.Duration(YamlObj.Chinesechess.Duration.Over)*time.Second, self.onReady)
	playList := self.T.GetPlayList()
	GlobalSender.NotifyOthers(playList, &protoMsg.ChineseChessStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Over,
			TotalTime: YamlObj.Chinesechess.Duration.Over,
		},
		Result: &protoMsg.ChineseChessResult{
			RedCamp:   self.redCamp,
			BlackCamp: self.blackCamp,
		},
	})
}

// ///////////////////////[初始化]///////////////////////////////////
// 初始化棋盘
func (self *Chinesechess) initBord() {
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
