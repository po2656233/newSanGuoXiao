package brtuitongzi

import (
	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
	. "sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/conf"
	. "sanguoxiao/internal/component/jettengame/game/internal/category"
	. "sanguoxiao/internal/component/jettengame/manger"
	_ "sanguoxiao/internal/component/jettengame/sql/mysql" //仅仅希望导入 包内的init函数
	protoMsg "sanguoxiao/internal/protocol/gofile"
	"sync"
	"time"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// BrTuitongziGame 继承于GameItem
type BrTuitongziGame struct {
	*Game

	personBetInfo *sync.Map                     //map[int64][]*protoMsg.BrTuitongziBetResp //下注信息
	openAreas     [][]byte                      //开奖纪录
	openInfo      *protoMsg.BrTuitongziOpenResp //开奖结果

	hostList    []int64 //申请列表(默认11个)
	bankerID    int64   //庄家ID
	superHostID int64   //超级抢庄ID
	bankerScore int64   //庄家积分
	keepTwice   uint8   //连续坐庄次数
	cards       []byte  //牌值
	startArea   int32   //起始区域

	inventory     int64 //库存
	confInventory int64
}

// 创建%v:%v实例
func NewBrTuitongzi(game *Game) *BrTuitongziGame {
	p := &BrTuitongziGame{
		Game: game,
	}

	p.bankerScore = 0
	p.startArea = INVALID
	p.Config = conf.YamlObj
	p.Init()
	return p
}

// ---------------------------%v:%v----------------------------------------------//
// 初始化信息
func (self *BrTuitongziGame) Init() {
	self.IsStart = true  // 是否启动
	self.IsClear = false // 不进行清场

	self.bankerID = 0         // 庄家ID
	self.superHostID = 0      // 超级抢庄ID
	self.keepTwice = 0        // 连续抢庄次数
	self.hostList = []int64{} // 申请列表(默认11个)

	self.openInfo = &protoMsg.BrTuitongziOpenResp{} // 结算结果
	self.personBetInfo = new(sync.Map)              // make(map[int64][]*protoMsg.BrTuitongziBetResp) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 20)         // 开奖纪录
	self.cards = make([]byte, 0)
	// 天-手牌
	self.openInfo.BankerCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}
	// 天-手牌
	self.openInfo.TianCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}
	// 地-手牌
	self.openInfo.DiCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}
	// 顺-手牌
	self.openInfo.ShunCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}

}

//获取玩家列表
//func (self *BrTuitongziGame) getPlayersName() []string {
//	var name []string
//	for _, pid := range self.PlayerList {
//		log.Debug("----玩家：%v", pid)
//		name = append(name, sqlhandle.CheckName(pid))
//	}
//	return name
//}

// //////////////////////////////////////////////////////////////
// 场景信息
func (self *BrTuitongziGame) Scene(args []interface{}) {
	_ = args[1]
	level := args[0].(int32)
	agent := args[1].(gate.Agent)
	userData := agent.UserData()
	if userData == nil {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game02])
		log.Debug("[Error][%v:%v场景] [未能查找到相关玩家] ")
		return
	}
	if level != self.G.Level {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game01])
		log.Debug("[Error][%v:%v场景] [房间等级不匹配] ")
		return
	}

	person := userData.(*Player)
	//场景信息
	StateInfo := &protoMsg.BrTuitongziSceneResp{}
	StateInfo.Inning = self.InningInfo.Number
	StateInfo.AwardAreas = self.openAreas // 录单
	StateInfo.RemainCount = int32(len(self.cards))
	//定时器
	StateInfo.TimeStamp = self.TimeStamp //////已过时长 应当该为传时间戳
	//筹码
	switch level {
	case RoomGeneral:
		StateInfo.Chips = self.Config.Chips.General //筹码
	case RoomMiddle:
		StateInfo.Chips = self.Config.Chips.Middle //筹码
	case RoomHigh:
		StateInfo.Chips = self.Config.Chips.High //筹码
	default:
		StateInfo.Chips = self.Config.Chips.Other //筹码
	}
	//玩家列表
	StateInfo.AllPlayers = self.GetSitterListMsg()

	//下注情况
	StateInfo.AreaBets = make([]int64, AREA_MAX)
	StateInfo.MyBets = make([]int64, AREA_MAX)
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		bets, ok1 := value.([]*protoMsg.BrTuitongziBetResp)
		if !ok1 {
			return true
		}
		for _, betInfo := range bets {
			if uid == person.UserID {
				StateInfo.MyBets[betInfo.BetArea] += betInfo.BetScore
			}
			StateInfo.AreaBets[betInfo.BetArea] += betInfo.BetScore
		}
		return true
	})

	//反馈场景信息
	GlobalSender.SendData(agent, StateInfo)

	//通知游戏状态
	timeInfo := &protoMsg.TimeInfo{}
	timeInfo.TimeStamp = self.TimeStamp
	timeInfo.OutTime = int32(time.Now().Unix() - self.TimeStamp)
	switch self.G.Scene {
	case protoMsg.GameScene_Start: //准备
		timeInfo.TotalTime = self.Config.Brtuitongzi.Duration.StartTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrTuitongziStateStartResp{
			Times:     timeInfo,
			StartArea: self.startArea,
			Inning:    self.InningInfo.Number,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = self.Config.Brtuitongzi.Duration.BetTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrTuitongziStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = self.Config.Brtuitongzi.Duration.OpenTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrTuitongziStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = self.Config.Brtuitongzi.Duration.OverTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrTuitongziStateOverResp{
			Times: timeInfo,
		})
	}

}

// 开始下注前定庄
func (self *BrTuitongziGame) Start(args []interface{}) {
	_ = args
	log.Release("[%v:%v]游戏創建成功", self.G.Name, self.ID)
	if self.IsStart {
		self.onStart()
		self.IsStart = false
	}
}

// 下注 直接扣除金币
func (self *BrTuitongziGame) Playing(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.BrTuitongziBetReq)
	//【传输对象】
	agent := args[1].(gate.Agent)

	userData := agent.UserData()
	if nil == userData {
		log.Debug("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[User02], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}
	if self.G.Scene == protoMsg.GameScene_Start {
		log.Debug("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game24], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game24])
		return
	}
	if self.G.Scene != protoMsg.GameScene_Playing {
		log.Debug("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game03], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game03])
		return
	}
	if m.BetScore <= 0 {
		log.Debug("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game07], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return
	}

	person := userData.(*Player)
	sitter, ok := self.GetSitter(person.UserID)
	if !ok {
		log.Debug("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game33], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game33])
		return
	}
	if self.bankerID == sitter.UserID {
		log.Debug("[%v:%v][%v:%v] Playing:->%v ", self.G.Name, self.ID, sitter.UserID, StatusText[Game27], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game27])
		return
	}
	if sitter.Gold < m.BetScore+sitter.Total {
		log.Debug("[%v:%v][%v:%v] %v Playing:->%v ", self.G.Name, self.ID, sitter.UserID, StatusText[Game05], sitter.Gold, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
		return
	}
	if AREA_MAX <= m.BetArea {
		log.Debug("[%v:%v][%v:%v] Playing:->%v ", self.G.Name, self.ID, sitter.UserID, StatusText[Game06], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game06])
		return
	}
	// log.Debug("[%v:%v] 玩家:%v Playing:->%v ", self.G.Name, self.ID, sitter.UserID, m)

	//下注成功
	msg := &protoMsg.BrTuitongziBetResp{}
	msg.UserID = sitter.UserID
	msg.BetArea = m.BetArea
	msg.BetScore = m.BetScore

	sitter.Score = m.BetScore
	sitter.Total += m.BetScore
	if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
		ok = false
		areaBetInfos, ok1 := value.([]*protoMsg.BrTuitongziBetResp)
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				areaBetInfos[index].BetScore = betItem.BetScore + m.BetScore
				log.Debug("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.G.Name, self.ID, sitter.UserID, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrTuitongziBetResp)
		}
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	} else {
		log.Debug("[%v:%v]\t玩家:%v 下注:%v", self.G.Name, self.ID, sitter.UserID, m)
		areaBetInfos := make([]*protoMsg.BrTuitongziBetResp, 0)
		areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrTuitongziBetResp)
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying

	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)
}

// 结算
func (self *BrTuitongziGame) Over(args []interface{}) {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	for _, sitter := range self.GetAllSitter() {
		uid := sitter.UserID
		if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
			betInfos, ok1 := value.([]*protoMsg.BrTuitongziBetResp)
			if !ok1 {
				continue
			}
			//每一次的下注信息
			for _, betInfo := range betInfos {
				log.Debug("[%v:%v]玩家:%v,下注区域:%v 下注金额:%v", self.G.Name, self.ID, uid, betInfo.BetArea, betInfo.BetScore)
				//玩家奖金
				bonusScore := self.BonusArea(betInfo.BetArea, betInfo.BetScore)
				sitter.Gain += bonusScore
				if bonusScore < 0 {
					allAreaInfo[betInfo.BetArea] -= betInfo.BetScore
				} else {
					allAreaInfo[betInfo.BetArea] += betInfo.BetScore
				}

			}
		}
	}

	//派奖
	checkout := &protoMsg.BrTuitongziCheckoutResp{}
	checkout.Acquires = allAreaInfo
	// 统一结算
	self.Calculate(GlobalSqlHandle.DeductMoney, false, false)
	for _, chair := range self.GetAllSitter() {
		if chair.Code == CodeSettle {
			checkout.MyAcquire = chair.Gain
		}
		GlobalSender.SendTo(chair.UserID, checkout)
	}

	log.Debug("[%v:%v]   \t结算注单... 各区域情况:%v", self.G.Name, self.ID, allAreaInfo)
}

// 更新
func (self *BrTuitongziGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]
	_ = args[1]

	flag := args[0].(int32)
	userID := args[1].(int64)
	switch flag {
	case GameUpdateOut: //玩家离开 不再向该玩家广播消息[] 删除
		self.UpdatePlayerList(DEL, userID, true)
	case GameUpdateEnter: //更新玩家列表
		self.UpdatePlayerList(ADD, userID, true)
	case GameUpdateHost: //更新玩家抢庄信息
		_ = args[2]
		self.host(args[2:])
	case GameUpdateSuperHost: //更新玩家超级抢庄信息
		_ = args[2]
		self.superHost(args[2:])
	case GameUpdateOffline: //更新玩家超级抢庄信息
	case GameUpdateReconnect: //更新玩家超级抢庄信息
	case GameUpdateGold:
		_ = args[2]
		GlobalSender.NotifyOthers(self.PlayerList, args[2].(*protoMsg.NotifyBalanceChange))
	case GameUpdateDisbanded:
		self.IsClear = true
	default:

	}
}

// 超级控制
func (self *BrTuitongziGame) SuperControl(args []interface{}) {
	_ = args[0]
	flag := args[0].(int32)
	switch flag {
	case GameControlClear:
		self.IsClear = true
	}
}

// 发牌
func (self *BrTuitongziGame) DispatchCard() {
	timeoutCount := 0
Dispatch:
	//2张牌
	if INVALID == len(self.cards) || len(self.cards) < 4*CardAmount {
		self.cards = Shuffle(CardListData[:])
	}
	self.openInfo.ShunCard.Cards = Deal(self.cards, CardAmount, IndexStart+Area_Shun)
	self.openInfo.TianCard.Cards = Deal(self.cards, CardAmount, IndexStart+Area_Tian)
	self.openInfo.DiCard.Cards = Deal(self.cards, CardAmount, IndexStart+Area_Di)
	self.openInfo.BankerCard.Cards = Deal(self.cards, CardAmount, IndexStart+Area_Zhuang)
	self.cards = self.cards[IndexStart+Area_Zhuang:]
	// 中奖区域
	self.openInfo.AwardArea = self.DeduceWin()

	//是否满足库存
	personAwardScore, ok := self.simulatedResult()
	if !ok && timeoutCount < 1000 {
		timeoutCount++
		//log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v 尝试次数:%v[库存警告]", self.G.Name, self.ID, self.inventory, personAwardScore, self.bankerScore, timeoutCount)
		goto Dispatch
	}
	log.Debug("[%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v", self.G.Name, self.ID, self.inventory, personAwardScore, self.bankerScore)
	self.inventory = personAwardScore
	self.bankerScore = personAwardScore - self.inventory
	if self.bankerScore < INVALID {
		self.bankerScore = INVALID
	}
	if self.inventory < INVALID {
		log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v [库存警告]", self.G.Name, self.ID, self.inventory, self.bankerScore)
	}
	// ===>>>
	//self.bankerScore = INVALID
	//self.inventory = INVALID
	log.Debug("[%v:%v]\t 库存信息 当前:%v 配置:%v 盈利:[%v]", self.G.Name, self.ID, self.inventory, self.confInventory, self.inventory-self.confInventory)

	//录单记录
	self.openAreas = append(self.openAreas, self.openInfo.AwardArea)
	lenOpen := len(self.openAreas)
	if MaxLoadNum < lenOpen {
		self.openAreas = self.openAreas[lenOpen-MaxLoadNum : lenOpen]
	}
	log.Debug("[%v:%v]\t开奖结果:%v Z:%v S:%v T:%v D:%v", self.G.Name, self.ID, self.openInfo.AwardArea,
		GetCardsText(self.openInfo.BankerCard.Cards), GetCardsText(self.openInfo.ShunCard.Cards),
		GetCardsText(self.openInfo.TianCard.Cards), GetCardsText(self.openInfo.DiCard.Cards))

	// 再次计算点数
}

// 开奖区域
func (self *BrTuitongziGame) DeduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	self.openInfo.ShunCard = JudgeCarType(self.openInfo.ShunCard.Cards)
	self.openInfo.TianCard = JudgeCarType(self.openInfo.TianCard.Cards)
	self.openInfo.DiCard = JudgeCarType(self.openInfo.DiCard.Cards)
	self.openInfo.BankerCard = JudgeCarType(self.openInfo.BankerCard.Cards)

	if Equal(self.openInfo.ShunCard, self.openInfo.BankerCard) {
		pWinArea[Area_Shun] = Draw
	} else if CompareCard(self.openInfo.BankerCard, self.openInfo.ShunCard) {
		pWinArea[Area_Shun] = Win
	} else {
		pWinArea[Area_Shun] = Lose
	}

	if Equal(self.openInfo.TianCard, self.openInfo.BankerCard) {
		pWinArea[Area_Tian] = Draw
	} else if CompareCard(self.openInfo.BankerCard, self.openInfo.TianCard) {
		pWinArea[Area_Tian] = Win
	} else {
		pWinArea[Area_Tian] = Lose
	}

	if Equal(self.openInfo.DiCard, self.openInfo.BankerCard) {
		pWinArea[Area_Di] = Draw
	} else if CompareCard(self.openInfo.BankerCard, self.openInfo.DiCard) {
		pWinArea[Area_Di] = Win
	} else {
		pWinArea[Area_Di] = Lose
	}

	return pWinArea
}

// 区域赔额
func (self *BrTuitongziGame) BonusArea(area int32, betScore int64) int64 {
	if Win == self.openInfo.AwardArea[area] {
		bonus := int64(Odds[int(area)]*100) * betScore / 100
		return bonus
	} else if Draw == self.openInfo.AwardArea[area] {
		return INVALID
	}
	return -betScore
}

// 开始|下注|结算|空闲
func (self *BrTuitongziGame) onStart() {
	self.Reset()
	// 校准库存
	if self.confInventory != self.Config.Brtuitongzi.Inventory {
		self.confInventory = self.Config.Brtuitongzi.Inventory
		//此处会重置以往的库存值
		self.inventory = self.confInventory
	}

	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(self.Config.Brtuitongzi.Duration.StartTime)*time.Second, self.onPlay)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	msg := &protoMsg.BrTuitongziStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brtuitongzi.Duration.StartTime,
			TotalTime: self.Config.Brtuitongzi.Duration.StartTime,
		},
		Dice:   RollDice(DiceCount),
		Inning: self.InningInfo.Number,
	}
	self.startArea = int32((msg.Dice[0] + msg.Dice[1]) % AREA_MAX)
	msg.StartArea = self.startArea
	GlobalSender.NotifyOthers(self.PlayerList, msg)

}

// [下注减法]
func (self *BrTuitongziGame) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	time.AfterFunc(time.Duration(self.Config.Brtuitongzi.Duration.BetTime)*time.Second, self.onOpen)

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrTuitongziStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brtuitongzi.Duration.BetTime,
			TotalTime: self.Config.Brtuitongzi.Duration.BetTime,
		},
	})
}

func (self *BrTuitongziGame) onOpen() {
	self.ChangeState(protoMsg.GameScene_Opening)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)

	//【发牌】
	self.DispatchCard()

	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrTuitongziStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brtuitongzi.Duration.OpenTime,
			TotalTime: self.Config.Brtuitongzi.Duration.OpenTime,
		},
		OpenInfo: self.openInfo,
	})

	time.AfterFunc(time.Duration(self.Config.Brtuitongzi.Duration.OpenTime)*time.Second, self.onOver)

}

// [结算加法]
func (self *BrTuitongziGame) onOver() {
	self.ChangeState(protoMsg.GameScene_Over)

	// 玩家结算(框架消息)
	self.Over(nil)

	//自动清场
	if self.IsClear || (SYSTEMID != self.T.HostID && self.T.Amount <= self.RunCount) {
		err := GlobalSqlHandle.DelGame(self.ID, INVALID)
		log.Release("[%v:%v]清场结果:%v", self.G.Name, self.ID, err)
		time.AfterFunc(time.Duration(WAITTIME)*time.Second, self.GameOver)
		return
	} else if SYSTEMID != self.T.HostID {
		remainCount := self.T.Amount - self.RunCount
		if self.T.Amount < self.RunCount {
			remainCount = 0
		}
		err := GlobalSqlHandle.UpdateGameAmount(self.ID, remainCount)
		log.Release("[%v:%v]第%v局结束 err:%v", self.G.Name, self.ID, self.RunCount, err)
	}

	// 开奖状态
	time.AfterFunc(time.Duration(self.Config.Brtuitongzi.Duration.OverTime)*time.Second, self.onStart)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrTuitongziStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brtuitongzi.Duration.OverTime,
			TotalTime: self.Config.Brtuitongzi.Duration.OverTime,
		},
	})
}

// 抢庄
func (self *BrTuitongziGame) host(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(gate.Agent)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	host := args[2].(*protoMsg.BrTuitongziHostReq)
	userID := userData.(*Player).UserID

	size := len(self.hostList)
	if self.bankerID == userID && host.IsWant {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
		return
	}

	if host.IsWant { //申请(已经是庄家)
		//列表上是否已经申请了
		for _, pid := range self.hostList {
			if pid == userID {
				GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
				return
			}
		}

		if BrTuitongziHostMax < size {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
		self.hostList = CopyInsert(self.hostList, size, userID).([]int64)

	} else { //取消
		for index, pid := range self.hostList {
			if userID == pid {
				self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
				break
			}
		}
		if size == len(self.hostList) {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
	}

	log.Debug("[%v:%v]有人来抢庄啦:%d 列表人数%d", self.G.Name, self.ID, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrTuitongziHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 超级抢庄
func (self *BrTuitongziGame) superHost(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(gate.Agent)
	host := args[2].(*protoMsg.BrTuitongziSuperHostReq)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}
	user := userData.(*Player)
	if user.Level < LessLevel {
		GlobalSender.SendResult(agent, FAILED, StatusText[User11])
		return
	}
	userID := user.UserID
	log.Debug("[%v:%v]有人要超级抢庄--->:%d", self.G.Name, self.ID, userID)

	if host.IsWant {
		if self.superHostID == 0 {
			self.superHostID = userID
			//超级抢庄放申请列表首位
			self.hostList = CopyInsert(self.hostList, 0, userID).([]int64)
		} else {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
	}

	// 通知
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrTuitongziSuperHostResp{
		UserID: self.bankerID,
		IsWant: host.IsWant,
	})

}

// 定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *BrTuitongziGame) permitHost() *protoMsg.BrTuitongziHostResp {
	//校验是否满足庄家条件 [5000 < 金额] 不可连续坐庄15次
	tempList := self.hostList
	log.Debug("[%v:%v]定庄.... 列表数据:%v", self.G.Name, self.ID, self.hostList)

	//befBankerID := self.bankerID 避免重复
	for index, pid := range tempList {
		person := GlobalPlayerManger.Get(pid)
		if person == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == person.UserID {
			log.Debug("[%v:%v]不再连续坐庄：%v", self.G.Name, self.ID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if person.Gold < 5000 {
			log.Debug("[%v:%v]玩家%d 金币%lf 少于5000不能申请坐庄", self.G.Name, self.ID, person.UserID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Debug("[%v:%v]连续坐庄次数:%d", self.G.Name, self.ID, self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]

		if banker := GlobalPlayerManger.Get(self.bankerID); banker != nil {
			self.bankerScore = banker.Gold
		}

		log.Debug("[%v:%v]确定庄家:%d", self.G.Name, self.ID, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		//self.bankerScore = 1000000
		log.Debug("[%v:%v]系统坐庄", self.G.Name, self.ID)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &protoMsg.BrTuitongziHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Debug("[%v:%v]广播上庄", self.G.Name, self.ID)
	return msg
}

// -----------------------逻辑层---------------------------
// 重新初始化[返回结果提供给外部清场用]
func (self *BrTuitongziGame) Reset() int64 {
	self.openInfo.BankerCard.Cards = make([]byte, CardAmount)
	self.openInfo.TianCard.Cards = make([]byte, CardAmount)
	self.openInfo.DiCard.Cards = make([]byte, CardAmount)
	self.openInfo.ShunCard.Cards = make([]byte, CardAmount)
	self.personBetInfo.Range(func(key, value any) bool {
		self.personBetInfo.Delete(key)
		return true
	})
	return self.Game.Reset()
}

func (self *BrTuitongziGame) winScore() int64 {
	//统计龙的区域
	bankerScroe := int64(INVALID)
	personScroe := int64(INVALID)
	self.personBetInfo.Range(func(key, value any) bool {
		betInfos, ok1 := value.([]*protoMsg.BrTuitongziBetResp)
		if !ok1 {
			return true
		}
		for _, betInfo := range betInfos {
			//玩家奖金
			if Win == self.openInfo.AwardArea[betInfo.BetArea] {
				personScroe += int64(Odds[int(betInfo.BetArea)]*100) * betInfo.BetScore / 100
			} else {
				bankerScroe += betInfo.BetScore
			}
		}
		return true
	})
	return bankerScroe - personScroe
}

func (self *BrTuitongziGame) simulatedResult() (int64, bool) {
	//var personAwardScore int64 = self.bankerScore + self.inventory
	if self.T.HostID != SYSTEMID {
		return self.bankerScore, true
	}
	var personAwardScore int64 = self.inventory
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		betInfos, ok1 := value.([]*protoMsg.BrTuitongziBetResp)
		if !ok1 {
			return true
		}
		if GetPlayerManger().CheckRobot(uid) {
			return true
		}
		for _, betInfo := range betInfos {
			//玩家奖金
			if Win == self.openInfo.AwardArea[betInfo.BetArea] {
				personAwardScore -= int64(Odds[int(betInfo.BetArea)]*100)*betInfo.BetScore/100 + betInfo.BetScore
			}
		}
		return true
	})

	log.Debug("[%v:%v]\t模拟结算...前:%v 后:%v   当前库存:%v", self.G.Name, self.ID, self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
