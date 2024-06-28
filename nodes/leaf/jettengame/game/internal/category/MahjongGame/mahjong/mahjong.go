package mahjong

//麻将

import (
	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
	"sort"
	"strconv"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/conf"
	. "superman/nodes/leaf/jettengame/game/internal/category"
	. "superman/nodes/leaf/jettengame/game/internal/category/MahjongGame"
	mselector "superman/nodes/leaf/jettengame/gamedata/gocLib/mahjonglib_cplus/selector"
	mjlib "superman/nodes/leaf/jettengame/gamedata/goclib/mahjonglib_cplus"
	. "superman/nodes/leaf/jettengame/manger"
	_ "superman/nodes/leaf/jettengame/sql/mysql"
	"time"
	"unsafe"
)

// MahjongGame 继承于GameItem
type MahjongGame struct {
	*Game

	bankerID    int64 //庄家ID
	currentID   int64 //当前出了牌的玩家
	currentCard int32 //当前牌值
	currentOp   protoMsg.MJOperate
	quanFeng    protoMsg.Direct //圈风
	startPos    int32

	isFlow    bool //是否流局
	isWaitOP  bool //需要等待操作 当前玩家还有操作 如杠后自摸
	isWaitOut bool //需等待出完牌
	index     int  //当前索引
	haveGang  bool

	wallHandle *mjlib.Wall
	selector   *mselector.MSelector
	mjMans     map[int]*protoMsg.MahjongPlayer      //索引加玩家
	wantDos    map[int]*protoMsg.MahjongOperateResp //

	openInfo *protoMsg.MahjongStateOpenResp
}

// NewMahjong 创建实例
func NewMahjong(game *Game) *MahjongGame {
	p := &MahjongGame{
		Game: game,
	}
	//algorithm.TestMJ()
	p.Config = conf.YamlObj

	//仅初始化一次
	p.bankerID = INVALID
	p.quanFeng = protoMsg.Direct_Centre
	p.wallHandle = mjlib.NewWall()
	p.selector = mselector.NewMSelector()
	p.Init()
	return p
}

// ---------------------------麻将----------------------------------------------//
// 初始化信息

func (self *MahjongGame) Init() {
	//需要每局都重置的数据
	self.IsStart = false // 是否启动
	self.IsClear = false // 不进行清场
	self.isFlow = false
	self.isWaitOP = false
	self.currentCard = INVALID
	self.startPos = INVALID
	self.currentID = INVALID
	self.selector.Clean()
	self.selector.SetTiles(MahjongCards144)
	self.selector.SetAILevel(mselector.AI_PLATINUM)
	self.openInfo = &protoMsg.MahjongStateOpenResp{}
	self.mjMans = make(map[int]*protoMsg.MahjongPlayer, 0)
	self.wantDos = make(map[int]*protoMsg.MahjongOperateResp, 0)
}

//------------------------国标麻将----------------------------------------------//
//初始化信息

func (self *MahjongGame) Scene(args []interface{}) {
	_ = args[1]
	level := args[0].(int32)
	agent := args[1].(gate.Agent)
	userData := agent.UserData()
	if userData == nil {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game02])
		log.Release("[Error][%v:%v场景] [未能查找到相关玩家] ")
		return
	}
	if level != self.G.Level {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game01])
		log.Release("[Error][%v:%v场景] [房间等级不匹配] ")
		return
	}

	//场景信息
	person := userData.(*Player)
	StateInfo := &protoMsg.MahjongSceneResp{}
	StateInfo.Inning = self.InningInfo.Number
	StateInfo.BankerID = self.bankerID
	StateInfo.TimeStamp = self.TimeStamp //////已过时长 应当该为传时间戳
	StateInfo.RunCount = int32(self.RunCount)
	StateInfo.CurrentID = self.currentID
	StateInfo.HuCard = self.currentCard
	StateInfo.QuanFeng = self.quanFeng
	StateInfo.StartPos = self.startPos
	StateInfo.RemainCount = int32(self.wallHandle.RemainLength())
	for _, sitter := range self.GetAllSitter() {
		man := self.toMahjongPlayer(sitter)
		handleMan := protoMsg.MahjongPlayer{
			MyInfo:     man.MyInfo,
			HandCards:  man.HandCards,
			KeZiCards:  man.KeZiCards,
			TableCards: man.TableCards,
			IsTing:     man.IsTing,
			GainScore:  man.GainScore,
			TotalScore: man.TotalScore,
			MenFeng:    man.MenFeng,
			SitDirect:  man.SitDirect,
			OpHints:    man.OpHints,
		}
		if man.MyInfo.UserID != person.UserID && self.G.Scene < protoMsg.GameScene_Opening {
			handleMan.HandCards = make([]int32, len(man.HandCards))
		}
		StateInfo.AllPlayers = append(StateInfo.AllPlayers, &handleMan)
	}

	//反馈场景信息
	GlobalSender.SendData(agent, StateInfo)

	//通知游戏状态
	sit, _ := self.GetSitter(person.UserID)
	man := self.toMahjongPlayer(sit)
	timeInfo := &protoMsg.TimeInfo{}
	timeInfo.TimeStamp = self.TimeStamp
	timeInfo.OutTime = int32(time.Now().Unix() - self.TimeStamp)
	switch self.G.Scene {
	case protoMsg.GameScene_Free: //空闲
		timeInfo.TotalTime = self.Config.Mahjong.Duration.OverTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateFreeResp{
			Times:  timeInfo,
			Inning: self.InningInfo.Number,
		})
	case protoMsg.GameScene_SitDirect: //定方位
		timeInfo.TotalTime = self.Config.Mahjong.Duration.DirectTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateDirectResp{
			Times:    timeInfo,
			MyDirect: man.SitDirect,
		})
	case protoMsg.GameScene_Decide: //定庄 轮庄
		timeInfo.TotalTime = self.Config.Mahjong.Duration.DecideTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateDecideResp{
			Times:    timeInfo,
			QuanFeng: self.quanFeng,
			MenFeng:  man.MenFeng,
			BankerID: self.bankerID,
		})
	case protoMsg.GameScene_RollDice: //掷骰子时长
		timeInfo.TotalTime = self.Config.Mahjong.Duration.RollTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateRollDiceResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Start:
		timeInfo.TotalTime = self.Config.Mahjong.Duration.DealTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateStartResp{
			Times:    timeInfo,
			StartPos: self.startPos,
		})
	case protoMsg.GameScene_Playing: //
		timeInfo.TotalTime = self.Config.Mahjong.Duration.PlayTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStatePlayingResp{
			Times:  timeInfo,
			UserID: self.currentID,
			Card:   int32(Fault),
		})
	case protoMsg.GameScene_WaitOperate:
		timeInfo.TotalTime = self.Config.Mahjong.Duration.OperateTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateWaitOperateResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //结算
		timeInfo.TotalTime = self.Config.Mahjong.Duration.OpenTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, self.openInfo)
	case protoMsg.GameScene_Over:
		timeInfo.TotalTime = self.Config.Mahjong.Duration.OverTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.MahjongStateOverResp{
			Times: timeInfo,
		})
	}

}

// 更新
func (self *MahjongGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]
	//
	_ = args[1]
	flag := args[0].(int32)
	userID := args[1].(int64)
	switch flag {
	case GameUpdateEnter: //更新玩家列表
		if self.UpdatePlayerList(ADD, userID, false) {
			chair, ok := self.GetSitter(userID)
			if !ok {
				GlobalSender.SendResultX(userID, FAILED, StatusText[Room03])
				return
			}
			GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.EnterGameMJResp{
				ChairNum: int32(chair.ChairID),
				Player:   self.toMahjongPlayer(chair),
				GameID:   self.ID,
			})
		}
	case GameUpdateOut: //玩家离开 不再向该玩家广播消息[] 删除
		self.UpdatePlayerList(DEL, userID, true)
		//self.host(args[2:])
	case GameUpdateRollDice:
		self.rollResp(args[2:])
	case GameUpdateOffline:
	case GameUpdateReconnect:
	case GameUpdateSiting:
	case GameUpdateTrustee:
		self.Trustee(args[2:], self.currentID, self.onAutoOut)
	case GameUpdateDisbanded:
		self.IsClear = true

	}
}

// 开始
func (self *MahjongGame) Start(args []interface{}) {
	_ = args
	//直接扣除金币
	log.Release("[%v:%v]游戏創建成功", self.G.Name, self.ID)
	self.Reset()
}

// 操作
func (self *MahjongGame) Playing(args []interface{}) {

	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.MahjongOperateReq)
	//【传输对象】
	agent := args[1].(gate.Agent)

	userData := agent.UserData()
	if nil == userData {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[User02], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	if self.G.Scene == protoMsg.GameScene_Start {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game48], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game48])
		return
	}
	if self.isWaitOut { // 自动出牌中
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game48], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game48])
		return
	}
	person := userData.(*Player)
	sitter, ok := self.GetSitter(person.UserID)
	if !ok {
		log.Release("[%v:%v][%v] Playing:->%v %v", self.G.Name, self.ID, StatusText[Game21], m, person.UserID)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game21])
		return
	}
	if !self.isWaitOP && person.UserID != self.currentID { //不需要等待时,当前操作者必须是该玩家
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game21], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game21])
		return
	}
	if SeatNoCode < sitter.Code {
		log.Release("[%v:%v][%v] Playing:->%v %v %v", self.G.Name, self.ID, StatusText[Game44], m, person.UserID, sitter.Code)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game44])
		return
	}

	//吃碰胡 自己的牌  不进行补花操作
	if self.currentID == person.UserID && (m.Code == protoMsg.MJOperate_Chi ||
		m.Code == protoMsg.MJOperate_Pong || m.Code == protoMsg.MJOperate_Hu ||
		m.Code == protoMsg.MJOperate_Gang) {
		log.Release("[%v:%v][%v] Playing:->%v %v", self.G.Name, self.ID, StatusText[Game37], m, person.UserID)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game37])
		return
	}
	//自摸|杠|暗杠 必须是自己摸得牌
	if self.currentID != person.UserID && (m.Code == protoMsg.MJOperate_ZiMo || m.Code == protoMsg.MJOperate_AnGang ||
		m.Code == protoMsg.MJOperate_MingGang) {
		log.Release("[%v:%v][%v] Playing:->%v %v", self.G.Name, self.ID, StatusText[Game37], m, person.UserID)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game37])
		return
	}

	//获取座位上玩家,并校准索引
	canDo := false
	curIndex := INVALID
	msg := &protoMsg.MahjongOperateResp{
		Code:   m.Code,
		UserID: person.UserID,
		Cards:  m.Cards,
	}
	//校验操作是否合法
	for index, temMan := range self.mjMans {
		if temMan.MyInfo.UserID == person.UserID {
			curIndex = index
			//校验操作是否合法
			for _, hint := range self.mjMans[curIndex].OpHints {
				if hint.Code == m.Code {
					if protoMsg.MJOperate_Ting == m.Code {
						//听牌
						temMan.IsTing = true
						GlobalSender.SendData(agent, msg)
						return
					}
					if protoMsg.MJOperate_Ting < m.Code {
						canDo = true
						break
					}

					//校验牌值
					size := len(m.Cards)
					if 0 < size && size == len(hint.Cards) {
						//是否符合牌值
						//排序
						sort.Slice(m.Cards, func(i, j int) bool {
							return m.Cards[i] < m.Cards[j]
						})
						sort.Slice(hint.Cards, func(i, j int) bool {
							return hint.Cards[i] < hint.Cards[j]
						})
						if m.Cards[0] == hint.Cards[0] {
							m.Cards = hint.Cards //校准牌值
							canDo = true
							break
						}
					}
				}
			}
			break
		}
	}
	//检验索引
	if curIndex == INVALID {
		log.Error("[%v:%v] 玩家:%v 检验索引出错Playing:->%v ", self.G.Name, self.ID, person.UserID, curIndex)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game21])
		return
	}

	//非法操作
	if !canDo && m.Code != protoMsg.MJOperate_Pass {
		log.Error("[%v:%v] 玩家:%v 不能操作Playing:->%v ", self.G.Name, self.ID, person.UserID, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game37])
		return
	}

	//合法操作 则进行检验是否所有(有提示)玩家已经操作了
	//杠牌时,仍需检测是否有玩家胡该牌值
	if m.Code == protoMsg.MJOperate_Gang || m.Code == protoMsg.MJOperate_MingGang || m.Code == protoMsg.MJOperate_AnGang {
		self.currentCard = m.Cards[0]
		for _, temMan := range self.mjMans {
			if temMan.MyInfo.UserID == person.UserID {
				continue
			}
			self.getHints(temMan.MyInfo.UserID, self.currentCard)
			for i, hint := range temMan.OpHints {
				//
				if hint.Code == protoMsg.MJOperate_Hu {
					if sit, ok1 := self.GetSitter(temMan.MyInfo.UserID); ok1 {
						sit.Code = SeatNoCode
					}
					GlobalSender.SendTo(temMan.MyInfo.UserID, &protoMsg.MahjongHintResp{
						UserID: temMan.MyInfo.UserID,
						Hints: []*protoMsg.MahjongHint{hint, {
							Code: protoMsg.MJOperate_Pass,
						}},
					})
				} else {
					temMan.OpHints = append(temMan.OpHints[:i], temMan.OpHints[i+1:]...)
				}
			}
		} //

	}
	//标记用户的操作状态 并统计是否都完成操作
	sitter.Code = byte(m.Code)
	isFinish := false
	nCount := INVALID
	for _, temMan := range self.mjMans {
		//统计还未操作的用户
		if sit, ok1 := self.GetSitter(temMan.MyInfo.UserID); ok1 && INVALID < len(temMan.OpHints) && SeatNoCode == sit.Code {
			nCount++
		}
	}
	if nCount == INVALID {
		isFinish = true
	}

	//过牌
	if m.Code == protoMsg.MJOperate_Pass {
		GlobalSender.SendData(agent, msg)
	} else {
		//记录玩家操作
		self.wantDos[curIndex] = msg
	}
	if isFinish {
		self.onEndWait()
	} else {
		_ = m.Code
	}

}

// 结算
func (self *MahjongGame) Over(args []interface{}) {
	_ = args
	//直接扣除金币
	log.Release("[%v:%v]结算", self.G.Name, self.ID)
	// 统一结算
	self.Calculate(GlobalSqlHandle.DeductMoney, true, true)
}

// 操作
func (self *MahjongGame) SuperControl(args []interface{}) {
	//直接扣除金币
	log.Release("操作")
	//return
	//m := args[0].(*protoMsg.GameMahjongPlaying)
	_ = args[0]
	flag := args[0].(int32)
	switch flag {
	case GameControlClear:
		self.IsClear = true
	}
}

//

func (self *MahjongGame) Ready(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.MahjongReadyReq)
	//【传输对象】
	agent := args[1].(gate.Agent)

	person := agent.UserData().(*Player)
	if self.G.Scene != protoMsg.GameScene_Free {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game04])
		return
	}

	if nil == m || person == nil {
		log.Release(":%v", StatusText[Game02])
		GlobalSender.SendResult(agent, FAILED, StatusText[Game02])
		return
	}

	//if int(self.Config.Mahjong.MaxPerson) < len(self.PlayerList) {
	//	log.Release(":%v", StatusText[Game13])
	//	GlobalSender.SendResult(agent, FAILED, StatusText[Game13])
	//	return
	//}

	if sitter, ok := self.GetSitter(person.UserID); ok {
		if sitter.Gold < int64(self.T.LessScore) {
			log.Release("[%v:%v]:%v", self.G.Name, self.ID, StatusText[User12])
			GlobalSender.SendResult(agent, FAILED, StatusText[User12])
			return
		}
		if m.IsReady {
			if sitter.State == protoMsg.PlayerState_PlayerAgree {
				GlobalSender.SendResult(agent, FAILED, StatusText[Game30])
				return
			}
			person.State = protoMsg.PlayerState_PlayerAgree
			self.ReadyCount++
			log.Release("[%v:%v]\t玩家:%v 准备了。当前准备人数 %v", self.G.Name, self.ID, person.UserID, self.ReadyCount)

		} else {
			if sitter.State == protoMsg.PlayerState_PlayerLookOn {
				GlobalSender.SendResult(agent, FAILED, StatusText[Game31])
				return
			}
			person.State = protoMsg.PlayerState_PlayerLookOn
			self.ReadyCount--
			log.Release("[%v:%v]\t玩家:%v 取消准备了。 当前准备人数:%v", self.G.Name, self.ID, person.UserID, self.ReadyCount)

		}
		sitter.State = person.State
	} else {
		self.UpdatePlayerList(DEL, person.UserID, true)
		return
	}

	log.Release("---------------------------")
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.MahjongReadyResp{
		UserID:  person.UserID,
		IsReady: m.IsReady,
	})
	if self.Config.Mahjong.MaxPerson <= self.ReadyCount {
		self.ChangeState(protoMsg.GameScene_Start)
	}
	if self.Config.Mahjong.MinPerson <= self.ReadyCount && !self.IsStart { //桌面上玩家准备好了,才进行发牌
		//启动游戏
		self.IsStart = true
		time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.DirectTime)*time.Second, self.onDirect)
	}
}

func (self *MahjongGame) DispatchCard() {
	if INVALID == self.startPos || int32(len(self.mjMans)) < self.Config.Mahjong.MinPerson {
		self.closeTimers()
		log.Release("[%v:%v]发牌出错", self.G.Name, self.ID)
		GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Game42],
		})
		time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.DealTime)*time.Second, self.onFree)
		return
	}

	//设置牌墙并洗牌
	self.wallHandle.SetTiles(MahjongCards144)
	self.wallHandle.Shuffle()
	self.wallHandle.SetStart(int(self.startPos))
	//if len(self.mjMans) < self.index || self.index == INVALID {
	//	self.index = IndexStart
	//}

	msg := &protoMsg.MahjongDealResp{}
	pos := self.startPos
	index := self.index
	size := DealCard
	for i := DealCard * len(self.mjMans); i > 0; i-- {
		for k, man := range self.mjMans {
			if index == k {
				if i <= SitCount {
					size = IndexStart
				}
				pos += int32(size)
				msg.UserID = man.MyInfo.UserID
				msg.CurPos = pos
				msg.HandCards = make([]int32, size)
				GlobalSender.NotifyButOne(self.PlayerList, msg.UserID, msg)

				msg.HandCards = ToArray(self.wallHandle.ForwardDrawMulti(size), int32(0)).([]int32)
				man.HandCards = append(man.HandCards, msg.HandCards...)
				sort.Slice(man.HandCards, func(i, j int) bool {
					return man.HandCards[i] < man.HandCards[j]
				})
				GlobalSender.SendTo(msg.UserID, msg)
				break
			}
		}
		index++
		if len(self.mjMans) < index {
			index = IndexStart
		}
	}

}

func (self *MahjongGame) CallScore(args []interface{}) {
	_ = args

}
func (self *MahjongGame) BetScore(args []interface{}) {
	_ = args

}

// 出牌处理
func (self *MahjongGame) OutCard(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.MahjongOutCardReq)
	//【传输对象】
	agent := args[1].(gate.Agent)
	userData := agent.UserData()
	if nil == userData {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[User02], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	if self.G.Scene == protoMsg.GameScene_Start || self.G.Scene == protoMsg.GameScene_Opening {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game48], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game48])
		return
	}
	if self.G.Scene != protoMsg.GameScene_Playing {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game21], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game21])
		return
	}

	person := userData.(*Player)
	sitter, ok := self.GetSitter(person.UserID)
	if !ok || sitter.UserID != self.currentID {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game21], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game21])
		return
	}

	canOut := false
	self.currentID = person.UserID
	var curMan *protoMsg.MahjongPlayer = nil
	for index, temMan := range self.mjMans {
		if temMan.MyInfo.UserID == person.UserID {
			curMan = temMan
			self.index = index
			log.Release("前  [%v:%v]手动出牌 玩家%v  牌:%v info:%v", self.G.Name, self.ID, person.UserID, m.Card, temMan.HandCards)
			//玩家手上是否有该张牌值
			for _, card := range temMan.HandCards {
				if card == m.Card {
					canOut = true
					temMan.HandCards = DeleteValue(temMan.HandCards, m.Card).([]int32)
					sort.Slice(temMan.HandCards, func(i, j int) bool {
						return temMan.HandCards[i] < temMan.HandCards[j]
					})
					log.Release("后  [%v:%v]手动出牌 玩家%v 牌:%v info:%v", self.G.Name, self.ID, person.UserID, m.Card, temMan.HandCards)

					//
					if sitter.Timer != nil {
						sitter.Timer.Stop()
					}

					//补花存至刻字里  继续抓牌
					if MAHJONG_SEASON1 <= card {
						ke := &protoMsg.MahjongKeZi{
							ByUid: self.currentID,
							Code:  protoMsg.MJOperate_BuHua,
							Cards: []int32{card},
						}
						temMan.KeZiCards = append(temMan.KeZiCards, ke)
						self.currentOp = protoMsg.MJOperate_BuHua
						GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.MahjongOperateResp{
							Code:   protoMsg.MJOperate_BuHua,
							UserID: temMan.MyInfo.UserID,
							Cards:  []int32{m.Card},
						})
						self.isWaitOP = true // 要等待出牌
						self.ChangeState(protoMsg.GameScene_WaitOperate)
						self.onPlay()
						return
					}
					temMan.TableCards = append(temMan.TableCards, m.Card)
					break
				}
			} //玩家手上是否有该张牌值
			break
		}
	}
	if !canOut {
		if curMan != nil {
			log.Release("[%v:%v] 玩家:%v 不能出牌:->%v 手牌:%v", self.G.Name, self.ID, person.UserID, m, curMan.HandCards)
		} else {
			log.Release("[%v:%v] 玩家:%v 不能出牌:->%v 手牌:无", self.G.Name, self.ID, person.UserID, m)
		}
		GlobalSender.SendResult(agent, FAILED, StatusText[Game41])
		return
	}

	//通知其他玩家
	self.currentCard = m.Card
	self.currentOp = protoMsg.MJOperate_NULL
	msg := &protoMsg.MahjongOutCardResp{
		UserID: person.UserID,
		Card:   m.Card,
	}
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	for k := range self.wantDos {
		self.wantDos[k] = nil
	}
	//向别的玩家发起提示
	canNext := true
	for _, oneMan := range self.mjMans {
		if self.currentID == oneMan.MyInfo.UserID {
			continue
		}
		if msgHint := self.getHints(oneMan.MyInfo.UserID, m.Card); msgHint != nil {
			msgHint = append(msgHint, &protoMsg.MahjongHint{
				Code: protoMsg.MJOperate_Pass,
			})
			GlobalSender.SendTo(oneMan.MyInfo.UserID, &protoMsg.MahjongHintResp{
				UserID: oneMan.MyInfo.UserID,
				Hints:  msgHint,
			})
			canNext = false
		}

	}
	log.Release("[%v:%v]玩家%v 出牌:->%v ", self.G.Name, self.ID, person.UserID, m)
	//下一位玩家
	if canNext {
		self.onNext()
	} else {
		//等待当前玩家操作
		self.onOperate()
	}

}

// 在此处时操作 即MahjongOperateReq的处理
func (self *MahjongGame) Discard(args []interface{}) {
	_ = args

}

//-----------------------逻辑层---------------------------

func (self *MahjongGame) onFree() {
	self.Reset()
	//重置开奖结果
	self.openInfo.AllPlayers = make([]*protoMsg.MahjongPlayer, 0)
	self.openInfo.Fans = make([]byte, 0)
	//场景信息
	StateInfo := &protoMsg.MahjongSceneResp{}
	StateInfo.Inning = self.InningInfo.Number
	StateInfo.BankerID = self.bankerID
	StateInfo.TimeStamp = self.TimeStamp //////已过时长 应当该为传时间戳
	StateInfo.RunCount = int32(self.RunCount)
	StateInfo.CurrentID = self.currentID
	StateInfo.HuCard = self.currentCard
	StateInfo.QuanFeng = self.quanFeng
	StateInfo.RemainCount = INVALID
	for _, sitter := range self.GetAllSitter() {
		//关闭所有玩家定时器
		if sitter.Timer != nil {
			sitter.Timer.Stop()
			sitter.Timer = nil
		}
		StateInfo.AllPlayers = append(StateInfo.AllPlayers, self.toMahjongPlayer(sitter))
	}

	//反馈场景信息
	GlobalSender.NotifyOthers(self.PlayerList, StateInfo)

	//通知状态
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.MahjongStateFreeResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: time.Now().Unix(),
			WaitTime:  self.Config.Mahjong.Duration.OverTime,
			OutTime:   INVALID,
			TotalTime: self.Config.Mahjong.Duration.OverTime,
		},
		Inning: self.InningInfo.Number,
	})
}

// 方位
func (self *MahjongGame) onDirect() {
	self.IsStart = true
	//过滤掉不在座位上 或未准备的玩家
	for _, sitter := range self.GetAllSitter() {
		if sitter.State == protoMsg.PlayerState_PlayerAgree || sitter.State == protoMsg.PlayerState_PlayerPickUp {
			size := len(self.mjMans)
			if int32(size) < self.Config.Mahjong.MinPerson {
				key := size + IndexStart
				self.mjMans[key] = self.toMahjongPlayer(sitter)
			}
		}
	}

	if int32(len(self.mjMans)) < self.Config.Mahjong.MinPerson || self.ReadyCount < self.Config.Mahjong.MinPerson {
		self.IsStart = false
		//--------------------
		self.ChangeState(protoMsg.GameScene_Free)
		GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.ResultResp{State: FAILED, Hints: StatusText[Game34]})
		log.Release("[%v:%v]人数不足。 当前准备人数:%v", self.G.Name, self.ID, self.ReadyCount)
		return
	}

	//改成每局都下发座位
	if self.RunCount == IndexStart {
		self.bankerID = INVALID
	}
	self.ChangeState(protoMsg.GameScene_SitDirect)
	msg := &protoMsg.MahjongStateDirectResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  self.Config.Mahjong.Duration.DirectTime,
			TotalTime: self.Config.Mahjong.Duration.DirectTime,
		},
	}
	//仅发送给座位玩家
	for _, sitter := range self.GetAllSitter() {
		if sitter.State < protoMsg.PlayerState_PlayerAgree {
			msg.MyDirect = protoMsg.Direct_Centre
			GlobalSender.SendTo(sitter.UserID, msg)
		}
	}

	//重置玩家信息
	for k, v := range self.mjMans {
		v.IsTing = false
		v.SitDirect = protoMsg.Direct(k)
		msg.MyDirect = v.SitDirect
		GlobalSender.SendTo(v.MyInfo.UserID, msg)
	}

	time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.DirectTime)*time.Second, self.onDecide)
	return

	//self.onDecide()
}

// 定庄
func (self *MahjongGame) onDecide() {
	self.ChangeState(protoMsg.GameScene_Decide)
	//圈风
	val := self.RunCount % 16
	//if 0 == val {
	//	self.quanFeng = protoMsg.Direct_Centre
	//}
	if 0 < val && val <= 4 {
		self.quanFeng = protoMsg.Direct_EAST
	} else if 4 < val && val <= 8 {
		self.quanFeng = protoMsg.Direct_South
	} else if 8 < val && val <= 12 {
		self.quanFeng = protoMsg.Direct_West
	} else if 12 < val && val < 16 {
		self.quanFeng = protoMsg.Direct_North
	} else {
		self.quanFeng = protoMsg.Direct_North
	}

	//锁定上一位庄家的索引
	_, self.index = self.getMan(self.bankerID)
	self.index += IndexStart
	if len(self.mjMans) < self.index {
		self.index = IndexStart
	}
	man, ok := self.mjMans[self.index]
	if !ok {
		log.Release("[%v:%v]索引无效:%v 详情:%v ", self.G.Name, self.ID, self.index, self.mjMans)
		self.onFree()

		return
	}
	//庄家
	self.bankerID = man.MyInfo.UserID
	self.currentID = self.bankerID

	//门风
	man.MenFeng = protoMsg.Direct_EAST
	for k, manOther := range self.mjMans {
		if self.index < k {
			manOther.MenFeng = protoMsg.Direct_EAST + protoMsg.Direct(k-self.index)
		} else if k < self.index {
			manOther.MenFeng = protoMsg.Direct_North + protoMsg.Direct_EAST - protoMsg.Direct(self.index-k)
		}

	}

	msg := &protoMsg.MahjongStateDecideResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  self.Config.Mahjong.Duration.DecideTime,
			TotalTime: self.Config.Mahjong.Duration.DecideTime,
		},
		BankerID: self.bankerID,
		QuanFeng: self.quanFeng,
	}

	//发送给座位玩家
	for _, man := range self.mjMans {
		msg.MenFeng = man.MenFeng
		GlobalSender.SendTo(man.MyInfo.UserID, msg)
	}

	//通知其他玩家
	for _, sitter := range self.GetAllSitter() {
		if sitter.State < protoMsg.PlayerState_PlayerAgree {
			msg.MenFeng = protoMsg.Direct_Centre
			GlobalSender.SendTo(sitter.UserID, msg)
		}

	}
	bankSit, ok1 := self.GetSitter(self.bankerID)
	if !ok1 {
		log.Release("[%v:%v]索引无效:%v 庄家:%v ", self.G.Name, self.ID, self.index, self.bankerID)
		return
	}
	if bankSit.Timer != nil {
		bankSit.Timer.Stop()
	}

	bankSit.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.DecideTime)*time.Second, self.onRollDice)
}

// onRollDice 掷骰子
func (self *MahjongGame) onRollDice() {
	self.ChangeState(protoMsg.GameScene_RollDice)
	msg := &protoMsg.MahjongStateRollDiceResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  self.Config.Mahjong.Duration.RollTime,
			TotalTime: self.Config.Mahjong.Duration.RollTime,
		},
	}
	GlobalSender.NotifyOthers(self.PlayerList, msg)

	bankSit, ok := self.GetSitter(self.bankerID)
	if !ok {
		GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.ResultPopResp{
			Flag:  SUCCESS,
			Title: StatusText[Title003],
			Hints: StatusText[Game22],
		})
		self.onFree()
		return
	}
	if bankSit.Timer != nil {
		bankSit.Timer.Stop()
	}
	bankSit.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.RollTime)*time.Second, self.onStart)
}

// 开始发牌
func (self *MahjongGame) onStart() {
	self.ChangeState(protoMsg.GameScene_Start)
	bankSit, ok := self.GetSitter(self.bankerID)
	if !ok {
		log.Release("[%v:%v]索引无效:%v 庄家:%v ", self.G.Name, self.ID, self.index, self.bankerID)
		return
	}
	if bankSit.Timer != nil {
		bankSit.Timer.Stop()
	}
	if self.startPos == INVALID {
		//表示庄家没有掷骰子,代为投掷
		dices := RollDice(DiceNum)
		self.startPos = int32(dices[0] + dices[1])
		msgRoll := &protoMsg.MahjongRollResp{
			UserID: self.bankerID,
			Dice:   dices,
		}
		GlobalSender.SendResultX(self.bankerID, SUCCESS, StatusText[Game40])
		GlobalSender.NotifyOthers(self.PlayerList, msgRoll)
		bankSit.Timer = time.AfterFunc(time.Duration(WaitTime)*time.Second, self.onStart)
		return
	}

	//通知开始游戏
	msg := &protoMsg.MahjongStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  self.Config.Mahjong.Duration.DealTime,
			TotalTime: self.Config.Mahjong.Duration.DealTime,
		},
		StartPos: self.startPos,
	}
	GlobalSender.NotifyOthers(self.PlayerList, msg)

	//发牌
	self.DispatchCard()

	bankSit.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.DealTime)*time.Second, self.onPlay)
}

// 出牌
func (self *MahjongGame) onPlay() {
	if int32(len(self.mjMans)) < self.Config.Mahjong.MinPerson {
		self.index = INVALID
		log.Release("[%v:%v]当前人数%v", self.G.Name, self.ID, len(self.mjMans))
		return
	}
	if self.wallHandle.RemainLength() <= INVALID {
		self.onOpen(true)
		return
	}
	//关闭所有玩家定时器
	self.closeTimers()
	for k := range self.wantDos {
		self.wantDos[k] = nil
	}
	sitter, ok := self.GetSitter(self.mjMans[self.index].MyInfo.UserID)
	if !ok {
		//self.onNext()
		log.Release("[%v:%v]无效的玩家%v信息", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID)
		GlobalSender.SendResultX(self.mjMans[self.index].MyInfo.UserID, FAILED, StatusText[User02])
		return
	}

	//发送牌值
	self.getCard(sitter)
	self.ChangeState(protoMsg.GameScene_Playing)
}

func (self *MahjongGame) onAutoOut() {
	self.isWaitOut = true
	defer func() {
		self.isWaitOut = false
	}()
	//关闭所有玩家定时器
	self.ChangeState(protoMsg.GameScene_WaitOperate)
	self.closeTimers()
	sitter, ok := self.GetSitter(self.mjMans[self.index].MyInfo.UserID)
	if !ok {
		log.Error("[%v:%v]无效的玩家%v信息", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID)
		GlobalSender.SendResultX(self.mjMans[self.index].MyInfo.UserID, FAILED, StatusText[User02])
		return
	}
	self.isWaitOP = false
	man, _ := self.getMan(sitter.UserID)
	size := len(man.HandCards)
	if man != nil && 0 < size {
		//删除手牌
		//card := man.HandCards[size-1]
		// 设置参数
		//self.selector.SetLack(0)
		self.selector.SetHandTilesSlice(*(*[]int)(unsafe.Pointer(&man.HandCards)))
		for _, keZi := range man.KeZiCards {
			self.selector.AddShowTilesSlice(*(*[]int)(unsafe.Pointer(&keZi.Cards)))
		}
		//self.selector.SetShowTilesSlice(showTiles)
		//self.selector.SetDiscardTilesSlice(discardTiles)
		self.selector.SetRemainTilesSlice(self.wallHandle.Remains())

		// 选牌
		card := int32(self.selector.GetSuggest())
		self.currentCard = card
		log.Release("[%v:%v]自动出牌 玩家%v 牌 info:%v", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID, man.HandCards)
		man.HandCards = DeleteValue(man.HandCards, card).([]int32)
		log.Release("[%v:%v]自动出牌 玩家%v 牌信息 %v info:%v", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID, card, man.HandCards)

		//判断是否自摸
		if msgHint := self.getHints(sitter.UserID, card); msgHint != nil {
			for i := len(msgHint) - 1; i >= 0; i-- {
				if msgHint[i].Code == protoMsg.MJOperate_ZiMo {
					//胡牌 则计算番数 除了胡牌可以同时多家
					self.openInfo.WinnerID = sitter.UserID
					man.HandCards = append(man.HandCards, card)
					self.onOpen(true)
					return
				}
				if msgHint[i].Code == protoMsg.MJOperate_MingGang || msgHint[i].Code == protoMsg.MJOperate_AnGang {
					msgN := &protoMsg.MahjongStatePlayingResp{
						UserID: sitter.UserID,
						Times: &protoMsg.TimeInfo{
							TotalTime: self.Config.Mahjong.Duration.PlayTime,
							TimeStamp: self.TimeStamp,
							WaitTime:  self.Config.Mahjong.Duration.PlayTime,
							OutTime:   INVALID,
						},
						Card: Fault,
					}
					GlobalSender.NotifyOthers(self.PlayerList, msgN)
					//删除碰时候的刻子
					for k, keZi := range man.KeZiCards {
						if keZi.Code == protoMsg.MJOperate_Pong && keZi.Cards[0] == msgHint[i].Cards[0] {
							man.KeZiCards = append(man.KeZiCards[:k], man.KeZiCards[k+1:]...)
							break
						}
					}
					// 明确牌值
					man.HandCards = append(man.HandCards, card)
					// 删除手牌
					for _, card1 := range msgHint[i].Cards {
						for k, handCard := range man.HandCards {
							if handCard == card1 {
								man.HandCards = append(man.HandCards[:k], man.HandCards[k+1:]...)
								break
							}
						}
					}
					//增加刻字
					if protoMsg.MJOperate_Pass < msgHint[i].Code && msgHint[i].Code < protoMsg.MJOperate_Ting {
						keZi := &protoMsg.MahjongKeZi{
							Code:  msgHint[i].Code,
							Cards: msgHint[i].Cards,
							ByUid: msgHint[i].ByUid,
						}
						man.KeZiCards = append(man.KeZiCards, keZi)
					}
					self.ChangeState(protoMsg.GameScene_WaitOperate)
					self.onPlay()
					return
				}
			}

		}

		//反馈操作
		msg := &protoMsg.MahjongOutCardResp{
			UserID: sitter.UserID,
			Card:   card,
		}

		self.currentOp = protoMsg.MJOperate_NULL
		GlobalSender.NotifyOthers(self.PlayerList, msg)

		if MAHJONG_SEASON1 <= card {
			ke := &protoMsg.MahjongKeZi{
				ByUid: sitter.UserID,
				Code:  protoMsg.MJOperate_BuHua,
				Cards: []int32{card},
			}
			man.KeZiCards = append(man.KeZiCards, ke)
			self.currentOp = protoMsg.MJOperate_BuHua
			GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.MahjongOperateResp{
				Code:   protoMsg.MJOperate_BuHua,
				UserID: sitter.UserID,
				Cards:  []int32{card},
			})
			self.isWaitOP = true
			self.ChangeState(protoMsg.GameScene_WaitOperate)
			self.onPlay()
			return
		}
		man.TableCards = append(man.TableCards, card)
		man.OpHints = nil //出完牌后不参与该牌的可操作项
		for _, oneMan := range self.mjMans {
			if oneMan.MyInfo.UserID == sitter.UserID {
				continue
			}
			if msgHint := self.getHints(oneMan.MyInfo.UserID, card); msgHint != nil {
				msgHint = append(msgHint, &protoMsg.MahjongHint{
					Code: protoMsg.MJOperate_Pass,
				})
				GlobalSender.SendTo(oneMan.MyInfo.UserID, &protoMsg.MahjongHintResp{
					UserID: oneMan.MyInfo.UserID,
					Hints:  msgHint,
				})
				self.isWaitOP = true
			}

		}
	}

	if self.isWaitOP {
		//出完牌就是等待操作
		self.onOperate()
	} else {
		self.onNext()
	}

}

// 等待操作
func (self *MahjongGame) onOperate() {
	self.ChangeState(protoMsg.GameScene_WaitOperate)
	log.Release("[%v:%v] 玩家:%v 等待中... Card:%v ", self.G.Name, self.ID, self.currentID, self.currentCard)
	for k := range self.wantDos {
		self.wantDos[k] = nil
	}
	msg := &protoMsg.MahjongStateWaitOperateResp{
		Times: &protoMsg.TimeInfo{
			TotalTime: self.Config.Mahjong.Duration.OperateTime,
			TimeStamp: self.TimeStamp,
			WaitTime:  self.Config.Mahjong.Duration.OperateTime,
			OutTime:   INVALID,
		},
	}
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	currentSit, ok := self.GetSitter(self.currentID)
	if !ok {
		log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, self.currentID)
		return
	}
	currentSit.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.OperateTime)*time.Second, self.onEndWait)
}

// 过了等待时间
func (self *MahjongGame) onEndWait() {
	//每次等待 确保只执行一次onEndWait
	if !self.isWaitOP {
		log.Release("[%v:%v][%v] onEndWait:->Card:%v op:%v UID:%v", self.G.Name, self.ID, StatusText[Game21], self.currentCard, self.currentOp, self.currentID)
		//GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.ResultResp{State: FAILED, Hints: StatusText[Game21]})
		return
	}
	self.isWaitOP = true
	//关闭所有玩家定时器
	self.closeTimers()

	// 操作优先级:一、自摸|胡; 二、杠; 三、碰|吃 四、过|听
	// 校验可操作的牌值数据 取操作优先级最高的玩家
	curIndex := INVALID
	currentOp := protoMsg.MJOperate_NULL
	for index, opMsg := range self.wantDos {
		if opMsg != nil && currentOp < opMsg.Code {
			currentOp = opMsg.Code
			curIndex = index
		}
	}
	//没有玩家进行实际操作,则转至下一位
	if curIndex == INVALID {
		self.onNext()
		return
	}
	//删除玩家桌牌
	otherMan := self.mjMans[self.index]
	if otherMan != nil {
		otherMan.TableCards = DeleteValue(otherMan.TableCards, self.currentCard).([]int32)
	}

	// 执行操作项
	for index, opMsg := range self.wantDos {
		if opMsg != nil && currentOp == opMsg.Code && (curIndex == index || currentOp == protoMsg.MJOperate_Hu) { //可以多家同时操作胡牌
			temMan := self.mjMans[index]
			for _, hint := range temMan.OpHints {
				if hint.Code == opMsg.Code && hint.Cards[0] == opMsg.Cards[0] && opMsg.UserID == temMan.MyInfo.UserID {
					// 通知玩家操作
					temMan.OpHints = make([]*protoMsg.MahjongHint, 0)
					GlobalSender.NotifyOthers(self.PlayerList, opMsg)
					log.Release("[%v:%v] 玩家:%v 操作后出牌:->%v ", self.G.Name, self.ID, opMsg.UserID, opMsg)

					//自摸 胡牌优先处理
					if opMsg.UserID == self.currentID {
						self.currentCard = hint.Cards[0]
					}
					switch opMsg.Code {
					case protoMsg.MJOperate_Hu, protoMsg.MJOperate_ZiMo: //胡牌 自摸
						//胡牌 则计算番数 除了胡牌可以同时多家
						self.openInfo.WinnerID = opMsg.UserID
						self.onOpen(true)
						return
					case protoMsg.MJOperate_AnGang: //暗杠
						fallthrough
					case protoMsg.MJOperate_MingGang: //明杠
						//删除碰时候的刻子
						for k, keZi := range temMan.KeZiCards {
							if keZi.Code == protoMsg.MJOperate_Pong && keZi.Cards[0] == opMsg.Cards[0] {
								temMan.KeZiCards = append(temMan.KeZiCards[:k], temMan.KeZiCards[k+1:]...)
								break
							}
						}
						// 明确牌值
						opMsg.Cards = []int32{opMsg.Cards[0], opMsg.Cards[0], opMsg.Cards[0], opMsg.Cards[0]}
					default:
					}
					//别人的牌都统一先添加到手牌
					if opMsg.UserID != self.currentID {
						temMan.HandCards = append(temMan.HandCards, self.currentCard)
					}
					// 删除手牌
					for _, card := range opMsg.Cards {
						for k, handCard := range temMan.HandCards {
							if handCard == card {
								temMan.HandCards = append(temMan.HandCards[:k], temMan.HandCards[k+1:]...)
								break
							}
						}
					}

					//增加刻字
					if protoMsg.MJOperate_Pass < opMsg.Code && opMsg.Code < protoMsg.MJOperate_Ting {
						keZi := &protoMsg.MahjongKeZi{
							Code:  opMsg.Code,
							Cards: opMsg.Cards,
							ByUid: hint.ByUid,
						}
						temMan.KeZiCards = append(temMan.KeZiCards, keZi)
					}
					self.mjMans[index] = temMan

					//转至下一位

					self.index = index
					self.currentID = opMsg.UserID
					if opMsg.Code == protoMsg.MJOperate_Gang || opMsg.Code == protoMsg.MJOperate_MingGang || opMsg.Code == protoMsg.MJOperate_AnGang {
						//进行摸牌
						self.ChangeState(protoMsg.GameScene_WaitOperate)
						self.onPlay()
						return
					}
					//等待玩家出牌
					self.ChangeState(protoMsg.GameScene_Playing)
					msgN := &protoMsg.MahjongStatePlayingResp{
						UserID: opMsg.UserID,
						Times: &protoMsg.TimeInfo{
							TotalTime: self.Config.Mahjong.Duration.PlayTime,
							TimeStamp: self.TimeStamp,
							WaitTime:  self.Config.Mahjong.Duration.PlayTime,
							OutTime:   INVALID,
						},
						Card: Fault,
					}
					for k := range self.wantDos {
						self.wantDos[k] = nil
					}
					GlobalSender.NotifyOthers(self.PlayerList, msgN)
					currentSit, ok := self.GetSitter(self.currentID)
					if !ok {
						log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, self.currentID)
						return
					}
					if currentSit.State == protoMsg.PlayerState_PlayerPickUp {
						timeout := 2 * time.Second
						if self.haveGang {
							timeout = time.Duration(self.Config.Mahjong.Duration.PlayTime+2) * time.Second
						}
						currentSit.Timer = time.AfterFunc(timeout, self.onAutoOut)
					} else {
						currentSit.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.PlayTime)*time.Second, self.onAutoOut)
					}
					return
				}
			}

		}
	}

	// 都开奖后,则进行结算
	if self.G.Scene == protoMsg.GameScene_Opening {
		time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.OpenTime)*time.Second, self.onOver)
	} else {
		self.onNext()
	}
}

// 开奖
func (self *MahjongGame) onOpen(isOver bool) {
	self.ChangeState(protoMsg.GameScene_Opening)
	//关闭所有玩家定时器
	self.closeTimers()

	//结算
	if strTiles, ok1 := self.toTilesString(self.openInfo.WinnerID); ok1 {

		isMy := self.mjMans[self.index].MyInfo.UserID == self.openInfo.WinnerID
		isJueZhang := self.isLast(self.currentCard)
		isHaiDi := self.wallHandle.RemainLength() == 0
		isGang := self.currentOp == protoMsg.MJOperate_MingGang || self.currentOp == protoMsg.MJOperate_AnGang
		if !isMy {
			strTiles += MJTilesText[int(self.currentCard)]
		}

		strHua := self.toHuaString(self.openInfo.WinnerID)
		man, _ := self.getMan(self.openInfo.WinnerID)

		fanTypes, fanCount, bRet := mjlib.CheckHu(strTiles, strHua, int(self.quanFeng), int(man.MenFeng), isMy, isJueZhang, isHaiDi, isGang, false)
		if bRet {
			self.openInfo.HuCard = self.currentCard
			self.openInfo.IsFlow = false
			self.openInfo.Fans = fanTypes
			self.openInfo.FanCount = fanCount

			//胡牌统计
			if isMy {
				//自摸
				for _, otherMan := range self.mjMans {
					if otherMan.MyInfo.UserID != self.openInfo.WinnerID {
						//输家
						otherMan.GainScore = -int64(self.T.LessScore) * int64(fanCount)
						otherMan.TotalScore += otherMan.GainScore
						currentSit, ok := self.GetSitter(otherMan.MyInfo.UserID)
						if !ok {
							log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, otherMan.MyInfo.UserID)
							return
						}
						currentSit.Gain = otherMan.GainScore
						currentSit.Total = otherMan.TotalScore

						//赢家
						man.GainScore -= otherMan.GainScore
						man.TotalScore += man.GainScore
					}
				}
			} else {
				//胡
				for _, otherMan := range self.mjMans {
					if otherMan.MyInfo.UserID == self.mjMans[self.index].MyInfo.UserID {
						//输家
						otherMan.GainScore = -int64(self.T.LessScore) * int64(fanCount)
						otherMan.TotalScore += otherMan.GainScore
						curSit, ok := self.GetSitter(otherMan.MyInfo.UserID)
						if !ok {
							log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, otherMan.MyInfo.UserID)
							return
						}
						curSit.Gain = otherMan.GainScore
						curSit.Total = otherMan.TotalScore

						//赢家
						man.GainScore -= otherMan.GainScore
						man.TotalScore += man.GainScore
					}
				}
			}
			//赢家
			winSit, ok := self.GetSitter(self.openInfo.WinnerID)
			if !ok {
				log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, self.openInfo.WinnerID)
				return
			}
			man.GainScore = man.GainScore * int64(self.T.Commission) * 100 / 10000
			winSit.Gain = man.GainScore
			winSit.Total = man.TotalScore
			log.Release("[%v:%v]\t玩家:%v 胡牌 共计%v番 牌值:%v", self.G.Name, self.ID, self.openInfo.WinnerID, fanCount, strTiles)
		} else {
			log.Error("[%v:%v]\t玩家:%v 胡牌 共计%v番 牌值:%v 当前牌值:%v", self.G.Name, self.ID, self.openInfo.WinnerID, fanCount, strTiles, self.currentCard)
			GlobalSender.SendResultX(self.openInfo.WinnerID, FAILED, StatusText[Game37])
			return
		}
	}

	//检测当前玩家的牌值是否满足胡牌
	// checkHu
	if self.openInfo.WinnerID == INVALID && self.wallHandle.RemainLength() == INVALID {
		self.openInfo.IsFlow = true
	}
	self.openInfo.AllPlayers = make([]*protoMsg.MahjongPlayer, 0)
	for _, man := range self.mjMans {
		self.openInfo.AllPlayers = append(self.openInfo.AllPlayers, man)
	}
	self.openInfo.Times = &protoMsg.TimeInfo{
		TotalTime: self.Config.Mahjong.Duration.OperateTime,
		TimeStamp: self.TimeStamp,
		WaitTime:  self.Config.Mahjong.Duration.OperateTime,
		OutTime:   INVALID,
	}
	GlobalSender.NotifyOthers(self.PlayerList, self.openInfo)
	if isOver {
		time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.OpenTime)*time.Second, self.onOver)
	}

}

// 结算
func (self *MahjongGame) onOver() {
	self.ChangeState(protoMsg.GameScene_Over)

	msg := &protoMsg.MahjongStateOverResp{
		Times: &protoMsg.TimeInfo{
			TotalTime: self.Config.Mahjong.Duration.OverTime,
			TimeStamp: self.TimeStamp,
			WaitTime:  self.Config.Mahjong.Duration.OverTime,
			OutTime:   INVALID,
		},
	}
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	self.Over(nil)

	// 公告炫富
	self.FlauntWealth()
	//自动清场
	if self.IsClear || (SYSTEMID != self.T.HostID && self.T.Amount <= self.RunCount) {
		err := GlobalSqlHandle.DelGame(self.ID, INVALID)
		log.Release("清场结果:%v", err)
		time.AfterFunc(time.Duration(WAITTIME)*time.Second, self.GameOver)
		return
	} else if SYSTEMID != self.T.HostID {
		remainCount := self.T.Amount - self.RunCount
		if self.T.Amount < self.RunCount {
			remainCount = 0
		}
		err := GlobalSqlHandle.UpdateGameAmount(self.ID, remainCount)
		log.Release("第%v局结束:%v err:%v", self.RunCount, err)
	}
	time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.OverTime)*time.Second, self.onFree)
}

// 下一位
func (self *MahjongGame) onNext() {
	if len(self.mjMans) == INVALID {
		log.Error("[%v:%v]无效的玩家信息", self.G.Name, self.ID)
		self.index = INVALID
		return
	}

	if self.wallHandle.RemainLength() <= INVALID {
		self.onOpen(true)
		return
	}
	//关闭所有玩家定时器
	self.closeTimers()
	for k := range self.wantDos {
		self.wantDos[k] = nil
	}

	//移至下一位
	if len(self.mjMans) < self.index+1 {
		self.index = IndexStart
	} else {
		self.index++
	}
	curIndex := INVALID
	for i, player := range self.mjMans {
		if self.currentID == player.MyInfo.UserID {
			curIndex = i
			break
		}
	}
	if curIndex == self.index {
		log.Error("[%v:%v]索引有误 玩家信息%v", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID)
		return
	}
	sitter, ok1 := self.GetSitter(self.mjMans[self.index].MyInfo.UserID)
	if !ok1 {
		log.Error("[%v:%v]无效的玩家信息%v", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID)
		return
	}

	//发送牌值
	self.getCard(sitter)
	self.ChangeState(protoMsg.GameScene_Playing)
}

func (self *MahjongGame) Reset() int64 {
	self.Init()
	return self.Game.Reset()
}

// -----------------------逻辑层---------------------------
func (self *MahjongGame) closeTimers() {
	//关闭所有玩家定时器
	for _, sitter := range self.GetAllSitter() {
		sitter.Code = SeatNoCode
		if sitter != nil && sitter.Timer != nil {
			sitter.Timer.Stop()
			sitter.Timer = nil
		}
	}
}

// -----------------------逻辑层---------------------------
func (self *MahjongGame) getHints(uid int64, card int32) []*protoMsg.MahjongHint {
	self.haveGang = false
	man, _ := self.getMan(uid)
	if man != nil && 0 < len(man.HandCards) {
		man.OpHints = make([]*protoMsg.MahjongHint, 0)
		//是否可以吃
		nSameCount := INVALID
		isMy := self.currentID == uid
		//1.吃碰杠
		for i := 1; i < len(man.HandCards); i++ {
			//是否可以吃
			if !isMy {
				if card < man.HandCards[i-1] && mjlib.IsSequence(card, man.HandCards[i-1], man.HandCards[i]) {
					hint := &protoMsg.MahjongHint{
						Code:  protoMsg.MJOperate_Chi,
						Cards: []int32{card, man.HandCards[i-1], man.HandCards[i]},
						ByUid: self.currentID,
					}
					man.OpHints = append(man.OpHints, hint)
				} else if man.HandCards[i-1] < card && card < man.HandCards[i] && mjlib.IsSequence(man.HandCards[i-1], card, man.HandCards[i]) {
					hint := &protoMsg.MahjongHint{
						Code:  protoMsg.MJOperate_Chi,
						Cards: []int32{man.HandCards[i-1], card, man.HandCards[i]},
						ByUid: self.currentID,
					}
					man.OpHints = append(man.OpHints, hint)
				} else if man.HandCards[i] < card && card-man.HandCards[i] == 1 && mjlib.IsSequence(man.HandCards[i-1], man.HandCards[i], card) {
					hint := &protoMsg.MahjongHint{
						Code:  protoMsg.MJOperate_Chi,
						Cards: []int32{man.HandCards[i-1], man.HandCards[i], card},
						ByUid: self.currentID,
					}
					man.OpHints = append(man.OpHints, hint)
				}
				if card == man.HandCards[i] {
					nSameCount++
				}
			}

		}
		if card == man.HandCards[0] {
			nSameCount++
		}
		if 1 < nSameCount {
			//可以碰
			hint := &protoMsg.MahjongHint{
				Code:  protoMsg.MJOperate_Pong,
				Cards: []int32{card, card, card},
				ByUid: self.currentID,
			}
			man.OpHints = append(man.OpHints, hint)

			if 2 < nSameCount {
				//可以杠
				hint1 := &protoMsg.MahjongHint{
					Code:  protoMsg.MJOperate_Gang,
					Cards: []int32{card, card, card, card},
					ByUid: self.currentID,
				}
				//可以暗杠 自己的牌值
				if isMy && 4 == nSameCount {
					hint1.Code = protoMsg.MJOperate_AnGang
					self.haveGang = true
				}

				man.OpHints = append(man.OpHints, hint1)
			}
		}

		//2.继续找可杠的牌
		if isMy {
			//从刻子里找是否可以明杠
			for _, keZi := range man.KeZiCards {
				nSameCount = INVALID
				for _, kzCard := range keZi.Cards {
					if kzCard == card {
						nSameCount++
					}
				}
				if nSameCount == 3 {
					hint1 := &protoMsg.MahjongHint{
						Code:  protoMsg.MJOperate_MingGang,
						Cards: []int32{card, card, card, card},
						ByUid: keZi.ByUid,
					}
					self.haveGang = true
					man.OpHints = append(man.OpHints, hint1)
					break
				}
			}

			////////////////////
			alreadyGang := false
			for _, keZi := range man.KeZiCards {
				if keZi.Cards[0] == card && (keZi.Code == protoMsg.MJOperate_AnGang ||
					keZi.Code == protoMsg.MJOperate_MingGang || keZi.Code == protoMsg.MJOperate_Gang) {
					alreadyGang = true
					break
				}
			}

			//没有杠过此牌
			if !alreadyGang {
				nSameCount = IndexStart
				for i := 1; i < len(man.HandCards); i++ {
					if man.HandCards[i-1] == man.HandCards[i] {
						nSameCount++
						if nSameCount == 4 || (nSameCount == 3 && man.HandCards[i-1] == card) {
							hint := &protoMsg.MahjongHint{
								Code:  protoMsg.MJOperate_AnGang,
								Cards: []int32{man.HandCards[i], man.HandCards[i], man.HandCards[i], man.HandCards[i]},
								ByUid: man.MyInfo.UserID,
							}
							self.haveGang = true
							man.OpHints = append(man.OpHints, hint)
							nSameCount = IndexStart
						}
					} else {
						nSameCount = IndexStart
					}
				}
			}
		}

		//3.是否可以胡牌
		if tiles, ok := self.toTilesString(uid); ok {
			tiles += MJTilesText[int(card)]
			_, _, bRet := mjlib.CheckHu(tiles, "", int(self.quanFeng), int(man.MenFeng), isMy, false, false, false, false)
			if bRet {
				log.Release("[%v:%v]玩家:%v 可以胡牌:%v", self.G.Name, self.ID, uid, tiles)
				hint := &protoMsg.MahjongHint{
					Code:  protoMsg.MJOperate_Hu,
					Cards: []int32{card},
					ByUid: self.currentID,
				}
				if isMy {
					hint.Code = protoMsg.MJOperate_ZiMo
				}
				man.OpHints = append(man.OpHints, hint)
			}
			if 0 < len(man.OpHints) {
				self.isWaitOP = true
			}
			//是否可以听牌
			if !man.IsTing && isMy {
				if cards, ok1 := mjlib.CheckTing(tiles); ok1 {
					needCards := make([]int32, len(cards))
					for k, card := range cards {
						needCards[k] = mjlib.AdjustCard(card)
					}
					hint := &protoMsg.MahjongHint{
						Code:  protoMsg.MJOperate_Ting,
						Cards: needCards,
						ByUid: self.currentID,
					}
					man.OpHints = append(man.OpHints, hint)
				}
			}
		} else {
			log.Release("[%v:%v]当前牌值:%v CL:%v", self.G.Name, self.ID, tiles, self.currentCard)
		}
		if 0 < len(man.OpHints) {
			return man.OpHints
		}

	}
	return nil
}

//

func (self *MahjongGame) getMan(uid int64) (*protoMsg.MahjongPlayer, int) {
	for index, v := range self.mjMans {
		if v.MyInfo != nil && v.MyInfo.UserID == uid {
			return v, index
		}
	}
	return nil, INVALID
}

func (self *MahjongGame) getCard(sitter *Chair) {
	self.isWaitOP = false
	self.currentID = sitter.UserID
	msg := &protoMsg.MahjongStatePlayingResp{
		UserID: sitter.UserID,
		Times: &protoMsg.TimeInfo{
			TotalTime: self.Config.MahjongGD.Duration.PlayTime,
			TimeStamp: self.TimeStamp,
			WaitTime:  self.Config.MahjongGD.Duration.PlayTime,
			OutTime:   INVALID,
		},
		Card: INVALID,
	}
	GlobalSender.NotifyButOne(self.PlayerList, sitter.UserID, msg)
	//清空玩家提示内容
	for _, tempMan := range self.mjMans {
		tempMan.OpHints = make([]*protoMsg.MahjongHint, 0)
	}
	card := int32(self.wallHandle.ForwardDraw())
	man := self.mjMans[self.index]
	//校验玩家的手牌牌值
	if 2 != (len(man.HandCards)+1)%3 || INVALID == card {
		GlobalSender.SendPopResultX(sitter.UserID, FAILED, StatusText[Title004], StatusText[Game41])
		log.Release("[%v:%v]无效的玩家%v 信息 :%v keZi:%v table:%v", self.G.Name, self.ID, self.mjMans[self.index].MyInfo.UserID, man.HandCards, man.KeZiCards, man.TableCards)
		return
	}

	msg.Card = card
	GlobalSender.SendTo(sitter.UserID, msg)
	msgHint := self.getHints(man.MyInfo.UserID, card)
	sort.Slice(man.HandCards, func(i, j int) bool {
		return man.HandCards[i] < man.HandCards[j]
	})
	man.HandCards = append(man.HandCards, card) // 应当新增手牌后,再发提示信息
	if msgHint != nil {
		GlobalSender.SendTo(man.MyInfo.UserID, &protoMsg.MahjongHintResp{
			UserID: man.MyInfo.UserID,
			Hints:  msgHint,
		})
	}

	log.Release("[%v:%v]玩家:%v 抓牌:%v 当前手牌:%v", self.G.Name, self.ID, man.MyInfo.UserID, card, man.HandCards)
	if sitter.State == protoMsg.PlayerState_PlayerPickUp {
		timeout := 2 * time.Second
		if self.haveGang {
			timeout = time.Duration(self.Config.Mahjong.Duration.PlayTime+2) * time.Second
		}
		sitter.Timer = time.AfterFunc(timeout, self.onAutoOut)
	} else {
		sitter.Timer = time.AfterFunc(time.Duration(self.Config.Mahjong.Duration.PlayTime)*time.Second, self.onAutoOut)
	}
}

//
//func (self *MahjongGame) getSize(man *protoMsg.MahjongPlayer) int32 {
//	cardSize := len(man.HandCards)
//	for _, ke := range man.KeZiCards {
//		cardSize += len(ke.Cards)
//	}
//	return cardSize
//}

func (self *MahjongGame) isLast(card int32) bool {
	if self.wallHandle.IsLast(int(card)) {
		//牌堆里已经没有
		nCount := 0
		for _, man := range self.mjMans {
			for _, tCard := range man.TableCards {
				if tCard == card {
					nCount++
				}
			}
			for _, keZi := range man.KeZiCards {
				for _, kCard := range keZi.Cards {
					if kCard == card {
						nCount++
					}
				}
			}

		}
		return 3 <= nCount

	}
	return false
}

func (self *MahjongGame) toTilesString(uid int64) (string, bool) {
	man, index := self.getMan(uid)
	strTiles := ""
	bRet := false
	if man != nil {

		for _, keZi := range man.KeZiCards {
			if keZi.Code == protoMsg.MJOperate_BuHua {
				continue
			}
			strTiles += "["
			size := len(keZi.Cards)
			for i := 0; i < size-1; i++ {
				if keZi.Cards[i] < MAHJONG_DOT_PLACEHOLDER {
					strTiles += strconv.Itoa(int(keZi.Cards[i] % 10))
				} else {
					strTiles += MJTilesText[int(keZi.Cards[i])]
				}

			}
			_, otherIndex := self.getMan(keZi.ByUid)
			strTiles += MJTilesText[int(keZi.Cards[size-1])]
			switch keZi.Code {
			case protoMsg.MJOperate_Chi:
				strTiles += "," + mjlib.OtherSit(index, otherIndex, false)
			case protoMsg.MJOperate_Pong:
				strTiles += "," + mjlib.OtherSit(index, otherIndex, false)
			case protoMsg.MJOperate_Gang:
				strTiles += "," + mjlib.OtherSit(index, otherIndex, false)
			case protoMsg.MJOperate_MingGang:
				strTiles += "," + mjlib.OtherSit(index, otherIndex, true)
			case protoMsg.MJOperate_AnGang:
			}
			strTiles = strings.Trim(strTiles, ",")
			strTiles += "]"
		}

		for _, handCard := range man.HandCards {
			strTiles += MJTilesText[int(handCard)]
		}
		bRet = true
	}
	return strTiles, bRet
}

func (self *MahjongGame) toHuaString(uid int64) string {
	man, _ := self.getMan(uid)
	strTiles := ""
	if man != nil {
		for _, keZi := range man.KeZiCards {
			if keZi.Code == protoMsg.MJOperate_BuHua {
				strTiles += MJTilesText[int(keZi.Cards[0])]
			}
		}
	}
	return strTiles
}

func (self *MahjongGame) toMahjongPlayer(sitter *Chair) *protoMsg.MahjongPlayer {
	for k, v := range self.mjMans {
		if v.MyInfo != nil && v.MyInfo.UserID == sitter.UserID {
			v.SitDirect = protoMsg.Direct(k)
			return v
		}
	}
	person := &protoMsg.MahjongPlayer{
		MyInfo:     sitter.ToMsg(),
		GainScore:  sitter.Gain,
		TotalScore: sitter.Total,
		SitDirect:  protoMsg.Direct_Centre,
		MenFeng:    protoMsg.Direct_Centre,
		IsTing:     false,
		HandCards:  make([]int32, 0),
		TableCards: make([]int32, 0),
		KeZiCards:  make([]*protoMsg.MahjongKeZi, 0),
		OpHints:    make([]*protoMsg.MahjongHint, 0),
	}

	return person
}

func (self *MahjongGame) rollResp(args []interface{}) {

	a := args[0].([]interface{})
	agent := a[1].(gate.Agent)
	person := agent.UserData().(*Player)
	if person == nil {
		GlobalSender.SendError(agent)
		return
	}
	if _, ok := a[0].(*protoMsg.MahjongRollReq); ok && self.G.Scene == protoMsg.GameScene_RollDice && person.UserID == self.bankerID {
		dices := RollDice(DiceNum)
		self.startPos = int32(dices[0] + dices[1])
		msgRoll := &protoMsg.MahjongRollResp{
			UserID: self.bankerID,
			Dice:   dices,
		}
		//关闭之前的定时器
		curSit, ok := self.GetSitter(self.currentID)
		if !ok {
			log.Release("[%v:%v]索引无效:%v 当前玩家ID:%v ", self.G.Name, self.ID, self.index, self.currentID)
			return
		}
		if curSit.Timer != nil {
			curSit.Timer.Stop()
		}
		GlobalSender.NotifyOthers(self.PlayerList, msgRoll)
		self.ChangeState(protoMsg.GameScene_Playing)
		curSit.Timer = time.AfterFunc(time.Duration(WaitTime)*time.Second, self.onStart)
		return
	}
	GlobalSender.SendResult(agent, FAILED, StatusText[User21])
}
