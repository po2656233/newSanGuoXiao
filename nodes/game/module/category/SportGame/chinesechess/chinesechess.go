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
	p.Init()
	return p
}

// Init 重置
func (self *Chinesechess) Init() {
	self.Game.Init()
	self.winUid = 0
	self.initBord()
	self.redCamp = nil
	self.blackCamp = nil
	self.timeout = YamlObj.Chinesechess.Duration.Play
}

// Scene 场景
func (self *Chinesechess) Scene(args []interface{}) {
	person := args[0].(*mgr.Player)
	if person == nil {
		log.Warnf("[%v:%v][Scene:%v] person is nil.", self.Name, self.T.Id, self.GameInfo.Scene)
		return
	}
	//
	uid := person.UserID
	mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessSceneResp{
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
	case protoMsg.GameScene_Free: //
	case protoMsg.GameScene_Setting: //
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStateSetResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Settime),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_WaitOperate: //
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStateConfirmResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Confirm),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Start: // 开始
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStateStartResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Start),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Playing: // 下棋
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStatePlayingResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Play),
			Uid:   self.curUid,
		})
	case protoMsg.GameScene_Opening: // 开奖
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStateOpenResp{
			Times:  getTimeInfo(YamlObj.Chinesechess.Duration.Open),
			WinUid: self.winUid,
		})
	case protoMsg.GameScene_Over: // 结算
		mgr.GetClientMgr().SendTo(uid, &protoMsg.ChineseChessStateOverResp{
			Times: getTimeInfo(YamlObj.Chinesechess.Duration.Over),
			Result: &protoMsg.ChineseChessResult{
				RedCamp:   self.redCamp,
				BlackCamp: self.blackCamp,
			},
		})
	default:
		log.Warnf("[%v:%v][Scene:%v] person:%v no have scence.", self.Name, self.T.Id, self.GameInfo.Scene, uid)
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
	if person.UserID != self.curUid {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game21])
		return
	}
	if !isValidMove(self.board, m.Origin, m.Target) {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game37])
		return
	}
	otherUid := self.redCamp.Uid
	if self.curUid == self.redCamp.Uid {
		otherUid = self.blackCamp.Uid
	}

	code := CanMove(self.board, m.Origin, m.Target)
	can := false
	switch code {
	case moveOk:
		can = true
	case moveFail:
	case moveBeJiangJu:
	case moveJiangJu:
		mgr.GetClientMgr().NotifyOthers(self.PlayerList, &protoMsg.ChineseChessJiangJuResp{BeJiangUid: otherUid})
	case moveJueSha:
		can = true
		mgr.GetClientMgr().NotifyOthers(self.PlayerList, &protoMsg.ChineseChessJueShaResp{BeJueShaUid: otherUid})
		self.winUid = self.curUid
	default:
	}
	if !can {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game63])
		return
	}
	msg := &protoMsg.ChineseChessMoveResp{
		Uid:    person.UserID,
		Origin: m.Origin,
		Target: m.Target,
	}
	count := 0
	for _, cell := range self.board.Cells {
		if cell.Row == m.Origin.Row && cell.Col == m.Origin.Col {
			cell.Core = protoMsg.XQPiece_NoXQPiece
			count++
		} else if cell.Row == m.Target.Row && cell.Col == m.Target.Col {
			cell.Core = m.Target.Core
			count++
		}
		if count == 2 {
			break
		}
	}
	mgr.GetClientMgr().NotifyOthers(self.PlayerList, msg)
	self.Timer.Stop()
	self.onPlay()
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
	mgr.GetClientMgr().NotifyOthers(self.PlayerList, msg)
	if self.redCamp != nil && self.blackCamp != nil {
		self.onSetTime()
	}
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
	if m.Timeout <= INVALID {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game62])
		return
	}

	person := agent.UserData().(*mgr.Player)
	if person.UserID != self.curUid {
		mgr.GetClientMgr().SendResult(agent, FAILED, StatusText[Game21])
		return
	}

	self.timeout = m.Timeout
	msg := &protoMsg.ChineseChessSetTimeResp{
		Uid:     person.UserID,
		Timeout: m.Timeout,
	}
	mgr.GetClientMgr().NotifyOthers(self.PlayerList, msg)
	self.onConfirmTime()
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
	mgr.GetClientMgr().NotifyOthers(self.PlayerList, msg)
	if !m.IsAgree {
		if self.Timer != nil {
			self.Timer.Stop()
		}
		self.T.RemoveChair(person.UserID)
		self.onReady()
		return
	}
}

// /////////////////////////[定时器事件]//////////////////////////////////////////

func (self *Chinesechess) onReady() {
	list := self.T.GetPlayList()
	if self.redCamp != nil {
		log.Infof("[%v:%d] 红方玩家:%d 已准备好", self.GameInfo.Name, self.T.Id, self.redCamp.Uid)
		list = utils.RemoveValue(list, self.redCamp.Uid)
	}
	if self.blackCamp != nil {
		log.Infof("[%v:%d] 黑方玩家:%d 已准备好", self.GameInfo.Name, self.T.Id, self.blackCamp.Uid)
		list = utils.RemoveValue(list, self.blackCamp.Uid)
	}
	self.SetPlaylist(list)
	resp := &protoMsg.ChineseChessStateReadyResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Ready,
			TotalTime: YamlObj.Chinesechess.Duration.Ready,
		},
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Ready, resp, YamlObj.Chinesechess.Duration.Ready, nil, self.onSetTime)
}

func (self *Chinesechess) onSetTime() {
	resp := &protoMsg.ChineseChessStateSetResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Settime,
			TotalTime: YamlObj.Chinesechess.Duration.Settime,
		},
		Uid: self.redCamp.Uid,
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Setting, resp, YamlObj.Chinesechess.Duration.Settime, func() bool {
		if self.blackCamp == nil || self.redCamp == nil {
			log.Infof("[%v:%d] 等他其他玩家准备 ", self.GameInfo.Name, self.T.Id)
			return false
		}
		self.SetPlaylist(self.T.GetPlayList())
		self.curUid = self.redCamp.Uid
		return true
	}, self.onConfirmTime)
}

func (self *Chinesechess) onConfirmTime() {
	self.curUid = self.blackCamp.Uid
	resp := &protoMsg.ChineseChessStateConfirmResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Confirm,
			TotalTime: YamlObj.Chinesechess.Duration.Confirm,
		},
		Uid: self.curUid,
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Confirm, resp, YamlObj.Chinesechess.Duration.Confirm, nil, self.onStart)

}
func (self *Chinesechess) onStart() {
	resp := &protoMsg.ChineseChessStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Start,
			TotalTime: YamlObj.Chinesechess.Duration.Start,
		},
		Uid: self.curUid,
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Start, resp, YamlObj.Chinesechess.Duration.Start, nil, self.onPlay)
}

func (self *Chinesechess) onPlay() {
	self.curUid = self.T.NextChairUID(self.curUid)
	resp := &protoMsg.ChineseChessStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.timeout,
			TotalTime: self.timeout,
		},
		Uid:      self.curUid,
		NowBoard: self.board,
	}
	// 如果时间到了,玩家未操作,则另一方获胜
	self.ChangeStateAndWork(protoMsg.GameScene_Playing, resp, self.timeout, func() bool {
		if self.curUid == INVALID {
			log.Errorf("出现严重错误,当前玩家ID %v 出错", self.curUid)
			return false
		}
		if self.winUid != INVALID {
			self.onOpen()
			return false
		}
		return true
	}, self.onOpen)
}

func (self *Chinesechess) onOpen() {
	if self.winUid == INVALID {
		// 超时后，玩家胜利
		self.winUid = self.redCamp.Uid
		if self.curUid == self.redCamp.Uid {
			self.winUid = self.blackCamp.Uid
		}
	}
	resp := &protoMsg.ChineseChessStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.Chinesechess.Duration.Open,
			TotalTime: YamlObj.Chinesechess.Duration.Open,
		},
		WinUid:   self.winUid,
		NowBoard: self.board,
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Opening, resp, YamlObj.Chinesechess.Duration.Open, nil, self.onOver)
}
func (self *Chinesechess) onOver() {
	resp := &protoMsg.ChineseChessStateOverResp{
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
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Over, resp, YamlObj.Chinesechess.Duration.Over, func() bool {
		// 结算
		self.Over(nil)
		if self.RunCount == self.T.Amount { // 满足运行次数之后,释放资源
			self.Close(func() *mgr.Table {
				return self.T
			})
			return false
		}
		self.Reset()
		self.T.ClearChairs()
		return true
	}, self.onReady)
}

// ///////////////////////[初始化]///////////////////////////////////
// 初始化棋盘
func (self *Chinesechess) initBord() {
	self.board = &protoMsg.XQBoardInfo{
		Cells: make([]*protoMsg.XQGrid, 0),
	}
	// 第一排
	row := 0
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
	// 红炮
	row = 2
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
	// 红兵
	row = 3
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

	// 黑卒
	row = 6
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
	// 黑炮
	row = 7
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

	// 最后一排
	row = 9
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
