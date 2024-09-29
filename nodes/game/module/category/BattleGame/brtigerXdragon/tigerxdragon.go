package brtigerXdragon

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"
	. "superman/internal/utils"
	. "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
	"sync"
	"time"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// TigerXdragonGame 继承于GameItem
type TigerXdragonGame struct {
	*Game
	T *Table

	personBetInfo *sync.Map                     //map[int64][]*protoMsg.TigerXdragonBetResp //下注信息
	openAreas     [][]byte                      //开奖纪录
	openInfo      *gameMsg.TigerXdragonOpenResp //开奖结果

	hostList    []int64 //申请列表(默认11个)
	bankerID    int64   //庄家ID
	superHostID int64   //超级抢庄ID
	bankerScore int64   //庄家积分
	keepTwice   uint8   //连续坐庄次数

	inventory     int64 //库存
	confInventory int64
}

// New 龙虎
func New(game *Game, table *Table) *TigerXdragonGame {
	p := &TigerXdragonGame{
		Game: game,
		T:    table,
	}

	p.bankerScore = 0
	p.Init()
	return p
}

// ---------------------------%v:%v----------------------------------------------//

// Init 初始化信息
func (self *TigerXdragonGame) Init() {
	self.bankerID = 0         // 庄家ID
	self.superHostID = 0      // 超级抢庄ID
	self.keepTwice = 0        // 连续抢庄次数
	self.hostList = []int64{} // 申请列表(默认11个)

	self.openInfo = &gameMsg.TigerXdragonOpenResp{} // 结算结果

	self.personBetInfo = new(sync.Map)      //make(map[int64][]*protoMsg.TigerXdragonBetResp) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 20) // 开奖纪录

}

// //////////////////////////////////////////////////////////////

// Scene 场景信息
func (self *TigerXdragonGame) Scene(args []interface{}) bool {
	if !self.Game.Scene(args) {
		return false
	}
	person := args[0].(*Player)
	//场景信息
	StateInfo := &gameMsg.TigerXdragonSceneResp{
		AllPlayers: &protoMsg.PlayerList{
			Items: make([]*protoMsg.PlayerInfo, 0),
		},
		Inning:     self.Inning,
		TimeStamp:  self.TimeStamp,
		AwardAreas: self.openAreas,
		AreaBets:   make([]int64, AREA_MAX),
		MyBets:     make([]int64, AREA_MAX),
	}

	//下注情况
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		bets, ok1 := value.([]*gameMsg.TigerXdragonBetResp)
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
	log.Infof("[%v:%v] [Scene:%v] [当前人数:%v] ", self.GameInfo.Name, self.T.Id, self.GameInfo.Scene, self.T.SitCount())
	self.T.ChairWork(func(chair *Chair) {
		StateInfo.AllPlayers.Items = append(StateInfo.AllPlayers.Items, chair.PlayerInfo)
	})

	//反馈场景信息
	uid := person.UserID
	GlobalSender.SendTo(uid, StateInfo)

	//通知游戏状态
	timeInfo := &protoMsg.TimeInfo{
		TimeStamp: self.TimeStamp,
		OutTime:   int32(time.Now().Unix() - self.TimeStamp),
	}
	switch self.GameInfo.Scene {
	case protoMsg.GameScene_Start: //准备
		timeInfo.TotalTime = YamlObj.TigerXdragon.Duration.Start
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &gameMsg.TigerXdragonStateStartResp{
			Times:  timeInfo,
			Inning: self.Inning,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = YamlObj.TigerXdragon.Duration.Play
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &gameMsg.TigerXdragonStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = YamlObj.TigerXdragon.Duration.Open
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &gameMsg.TigerXdragonStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = YamlObj.TigerXdragon.Duration.Over
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &gameMsg.TigerXdragonStateOverResp{
			Times: timeInfo,
		})
	}
	return true
}

// 开始下注前定庄
func (self *TigerXdragonGame) Start(args []interface{}) bool {
	_ = args
	log.Infof("[%v:%v]游戏創建成功", self.Name, self.T.Id)
	if self.IsStart {
		self.onStart()
		self.IsStart = false
	}
	return true
}

// Playing 下注 直接扣除金币
func (self *TigerXdragonGame) Playing(args []interface{}) bool {
	_ = args[1]
	if !self.Game.Playing(args) {
		return false
	}
	//【消息】
	m := args[0].(*gameMsg.TigerXdragonBetReq)
	//【传输对象】
	agent := args[1].(Agent)
	if m.BetScore <= 0 {
		log.Infof("[%v:%v][%v] Playing:->%v ", self.Name, self.T.Id, StatusText[Game07], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return false
	}
	userData := agent.UserData()
	person := userData.(*Player)
	sitter := self.T.GetChair(person.UserID)
	if sitter == nil {
		log.Infof("[%v:%v][%v] Playing:->%v ", self.Name, self.T.Id, StatusText[Game33], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game33])
		return false
	}
	if self.bankerID == sitter.UserID {
		log.Infof("[%v:%v][%v:%v] Playing:->%v ", self.Name, self.T.Id, sitter.UserID, StatusText[Game27], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game27])
		return false
	}
	if sitter.Gold < m.BetScore+sitter.Total {
		log.Infof("[%v:%v][%v:%v] %v Playing:->%v ", self.Name, self.T.Id, sitter.UserID, StatusText[Game05], sitter.Gold, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
		return false
	}
	if AREA_MAX <= m.BetArea {
		log.Infof("[%v:%v][%v:%v] Playing:->%v ", self.Name, self.T.Id, sitter.UserID, StatusText[Game06], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game06])
		return false
	}
	//下注成功
	msg := &gameMsg.TigerXdragonBetResp{}
	msg.UserID = sitter.UserID
	msg.BetArea = m.BetArea
	msg.BetScore = m.BetScore

	sitter.Score = m.BetScore
	sitter.Total += m.BetScore
	if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
		ok = false
		areaBetInfos, ok1 := value.([]*gameMsg.TigerXdragonBetResp)
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				areaBetInfos[index].BetScore = betItem.BetScore + m.BetScore
				log.Infof("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.Name, self.T.Id, sitter.UserID, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*gameMsg.TigerXdragonBetResp)
		}
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	} else {
		log.Infof("[%v:%v]\t玩家:%v 下注:%+v", self.Name, self.T.Id, sitter.UserID, m)
		areaBetInfos := make([]*gameMsg.TigerXdragonBetResp, 0)
		areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*gameMsg.TigerXdragonBetResp)
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying

	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	return true
}

// 结算
func (self *TigerXdragonGame) Over(args []interface{}) bool {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		betInfos, ok1 := value.([]*gameMsg.TigerXdragonBetResp)
		if !ok || !ok1 {
			return true
		}
		//每一次的下注信息
		sitter := self.T.GetChair(uid)
		if sitter == nil {
			return true
		}
		for _, betInfo := range betInfos {
			log.Infof("[%v:%v]玩家:%v,下注区域:%v 下注金额:%v", self.Name, self.T.Id, uid, betInfo.BetArea, betInfo.BetScore)
			//玩家奖金
			bonusScore := self.BonusArea(betInfo.BetArea, betInfo.BetScore)
			sitter.Gain += bonusScore
			if bonusScore < 0 {
				allAreaInfo[betInfo.BetArea] -= betInfo.BetScore
			} else {
				allAreaInfo[betInfo.BetArea] += betInfo.BetScore
			}

		}
		return true
	})

	// 结算
	result := fmt.Sprintf("龙:%v  虎:%v 中奖区域:%v", GetCardText(self.openInfo.Cards[0]), GetCardText(self.openInfo.Cards[1]), self.openInfo.AwardArea)
	self.T.ChairSettle(self.Name, self.Inning, result)

	// 移除已离线玩家
	noLiveIds := self.T.CheckChairsNoLive()
	for _, uid := range noLiveIds {
		self.RemovePlayer(uid)
	}

	// 派奖信息
	checkout := &gameMsg.TigerXdragonCheckoutResp{}
	checkout.Acquires = allAreaInfo
	self.T.ChairWork(func(chair *Chair) {
		if chair.Total != INVALID {
			checkout.MyAcquire = chair.Gain
		}
		GlobalSender.SendTo(chair.UserID, checkout)
	})

	// 统一结算
	//self.Calculate(GlobalSqlHandle.DeductMoney, false, false)
	//for _, chair := range self.GetAllSitter() {
	//	if chair.Code == CodeSettle {
	//		checkout.MyAcquire = chair.Gain
	//	}
	//	GlobalSender.SendTo(chair.UserID, checkout)
	//}

	log.Infof("[%v:%v]   \t结算注单... 各区域情况:%v", self.Name, self.T.Id, allAreaInfo)
	return true
}

// 发牌
func (self *TigerXdragonGame) DispatchCard() {
	timeoutCount := 0
Dispatch:
	//2张牌
	cards := Shuffle(CardListData[:])
	self.openInfo.Cards = Deal(cards, CardAmount, IndexStart)
	log.Infof("[%v:%v]\t发牌 龙:%v  虎:%v", self.Name, self.T.Id, GetCardText(self.openInfo.Cards[0]), GetCardText(self.openInfo.Cards[1]))
	// 中奖区域
	self.DeduceWin()

	//是否满足库存
	personAwardScore, ok := self.simulatedResult()
	if !ok && timeoutCount < 1000 {
		timeoutCount++
		//log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v 尝试次数:%v[库存警告]", self.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore, timeoutCount)
		goto Dispatch
	}
	log.Infof("[%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v", self.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore)
	self.inventory = personAwardScore
	self.bankerScore = personAwardScore - self.inventory
	if self.bankerScore < INVALID {
		self.bankerScore = INVALID
	}
	if self.inventory < INVALID {
		log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v [库存警告]", self.Name, self.T.Id, self.inventory, self.bankerScore)
	}
	// ===>>>
	//self.bankerScore = INVALID
	//self.inventory = INVALID
	log.Infof("[%v:%v]\t 库存信息 当前:%v 配置:%v 盈利:[%v]", self.Name, self.T.Id, self.inventory, self.confInventory, self.inventory-self.confInventory)

	//录单记录
	self.openAreas = append(self.openAreas, self.openInfo.AwardArea)
	lenOpen := len(self.openAreas)
	if MaxLoadNum < lenOpen {
		self.openAreas = self.openAreas[lenOpen-MaxLoadNum : lenOpen]
	}

	GlobalSender.NotifyOthers(self.PlayerList, self.openInfo)
	// 再次计算点数
}

// 开奖区域
func (self *TigerXdragonGame) DeduceWin() []byte {
	self.openInfo.AwardArea = make([]byte, AREA_MAX)
	//胜利区域--------------------------

	//大小
	Dsum := GetCardValue(self.openInfo.Cards[0])
	Tsum := GetCardValue(self.openInfo.Cards[1])

	if Tsum < Dsum {
		self.openInfo.AwardArea[Area_Dragon] = Win
	} else if Dsum < Tsum {
		self.openInfo.AwardArea[Area_Tiger] = Win
	}
	//龙的花色
	DColor := GetCardColor(self.openInfo.Cards[0])
	switch DColor {
	case COLOR_Diamond:
		self.openInfo.AwardArea[Area_DDiamond] = Win
	case COLOR_Club:
		self.openInfo.AwardArea[Area_DClub] = Win
	case COLOR_Heart:
		self.openInfo.AwardArea[Area_DHeart] = Win
	case COLOR_Spade:
		self.openInfo.AwardArea[Area_DSpade] = Win
	}

	//虎的花色
	TColor := GetCardColor(self.openInfo.Cards[1])
	switch TColor {
	case COLOR_Diamond:
		self.openInfo.AwardArea[Area_TDiamond] = Win
	case COLOR_Club:
		self.openInfo.AwardArea[Area_TClub] = Win
	case COLOR_Heart:
		self.openInfo.AwardArea[Area_THeart] = Win
	case COLOR_Spade:
		self.openInfo.AwardArea[Area_TSpade] = Win
	}

	if Dsum == Tsum {
		if TColor < DColor {
			self.openInfo.AwardArea[Area_Dragon] = Win
		} else if DColor < TColor {
			self.openInfo.AwardArea[Area_Tiger] = Win
		} else {
			self.openInfo.AwardArea[Area_Drawn] = Win
		}
	}

	//庄输赢
	winScroe := self.winScroe()
	if winScroe < INVALID {
		self.openInfo.AwardArea[Area_Lose] = Win
	} else if INVALID < winScroe {
		self.openInfo.AwardArea[Area_Win] = Win
	}

	return self.openInfo.AwardArea
}

// 区域赔额
func (self *TigerXdragonGame) BonusArea(area int32, betScore int64) int64 {
	if Win == self.openInfo.AwardArea[area] {
		bonus := int64(Odds[int(area)]*100)*betScore/100 - betScore
		return bonus
	}
	return -betScore
}

// 开始|下注|结算|空闲
func (self *TigerXdragonGame) onStart() {
	self.reset()
	// 校准库存
	if self.confInventory != YamlObj.TigerXdragon.Inventory {
		self.confInventory = YamlObj.TigerXdragon.Inventory
		//此处会重置以往的库存值
		self.inventory = self.confInventory
	}

	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(YamlObj.TigerXdragon.Duration.Start)*time.Second, self.onPlay)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.TigerXdragon.Duration.Start,
			TotalTime: YamlObj.TigerXdragon.Duration.Start,
		},
		Inning: self.Inning,
	})

}

// [下注减法]
func (self *TigerXdragonGame) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	time.AfterFunc(time.Duration(YamlObj.TigerXdragon.Duration.Play)*time.Second, self.onOpen)

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.TigerXdragon.Duration.Play,
			TotalTime: YamlObj.TigerXdragon.Duration.Play,
		},
	})
}

func (self *TigerXdragonGame) onOpen() {
	self.ChangeState(protoMsg.GameScene_Opening)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)

	//【发牌】
	self.DispatchCard()

	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.TigerXdragon.Duration.Open,
			TotalTime: YamlObj.TigerXdragon.Duration.Open,
		},
		OpenInfo: self.openInfo,
	})

	time.AfterFunc(time.Duration(YamlObj.TigerXdragon.Duration.Open)*time.Second, self.onOver)
	log.Infof("[%v:%v]\t开奖中===>牌值:%v......区域:%v", self.Name, self.T.Id, self.openInfo.Cards, self.openInfo.AwardArea)
}

// [结算加法]
func (self *TigerXdragonGame) onOver() {

	self.ChangeState(protoMsg.GameScene_Over)

	// 玩家结算(框架消息)
	self.Over(nil)

	// 有人玩就减去一场
	self.personBetInfo.Range(func(key, value any) bool {
		self.T.CalibratingRemain(Default)
		return false
	})

	//自动清场
	if self.IsClear || self.Close(func() *Table {
		return self.T
	}) {
		return
	}

	// 开奖状态
	time.AfterFunc(time.Duration(YamlObj.TigerXdragon.Duration.Over)*time.Second, self.onStart)
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.TigerXdragon.Duration.Over,
			TotalTime: YamlObj.TigerXdragon.Duration.Over,
		},
	})
}

// 抢庄
func (self *TigerXdragonGame) host(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	host := args[2].(*gameMsg.TigerXdragonHostReq)
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

		if TigerXdragonHostMax < size {
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

	log.Infof("[%v:%v]\t有人来抢庄啦:%d 列表人数%d", self.Name, self.T.Id, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 超级抢庄
func (self *TigerXdragonGame) superHost(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)
	host := args[2].(*gameMsg.TigerXdragonSuperHostReq)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}
	user := userData.(*Player)
	userID := user.UserID
	log.Infof("[%v:%v]\t有人要超级抢庄--->:%d", self.Name, self.T.Id, userID)

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
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.TigerXdragonSuperHostResp{
		UserID: self.bankerID,
		IsWant: host.IsWant,
	})

}

// 定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *TigerXdragonGame) permitHost() *gameMsg.TigerXdragonHostResp {
	//校验是否满足庄家条件 [5000 < 金额] 不可连续坐庄15次
	tempList := self.hostList
	log.Infof("[%v:%v]\t定庄.... 列表数据:%v", self.Name, self.T.Id, self.hostList)

	//befBankerID := self.bankerID 避免重复
	for index, pid := range tempList {
		person := GlobalPlayerMgr.Get(pid)
		if person == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == person.UserID {
			log.Infof("[%v:%v]\t不再连续坐庄：%v", self.Name, self.T.Id, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if person.Gold < 5000 {
			log.Infof("[%v:%v]\t玩家%d 金币%lf 少于5000不能申请坐庄", self.Name, self.T.Id, person.UserID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Infof("[%v:%v]\t连续坐庄次数:%d", self.Name, self.T.Id, self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]

		if banker := GlobalPlayerMgr.Get(self.bankerID); banker != nil {
			self.bankerScore = banker.Gold
		}

		log.Infof("[%v:%v]\t确定庄家:%d", self.Name, self.T.Id, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		//self.bankerScore = 1000000
		log.Infof("[%v:%v]\t系统坐庄", self.Name, self.T.Id)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &gameMsg.TigerXdragonHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Infof("[%v:%v]\t广播上庄")
	return msg
}

// -----------------------逻辑层---------------------------
// 重新初始化[返回结果提供给外部清场用]
func (self *TigerXdragonGame) reset() {
	self.personBetInfo.Range(func(key, value any) bool {
		self.personBetInfo.Delete(key)
		return true
	})
	self.Game.Init()
	return
}

func (self *TigerXdragonGame) winScroe() int64 {
	//统计龙的区域
	bankerScroe := int64(INVALID)
	personScroe := int64(INVALID)

	self.personBetInfo.Range(func(key, value any) bool {
		betInfos, ok1 := value.([]*gameMsg.BrcowcowBetResp)
		if !ok1 {
			return true
		}
		for _, betInfo := range betInfos {
			if betInfo.BetArea != Area_Lose && betInfo.BetArea != Area_Win {
				//玩家奖金
				if Win == self.openInfo.AwardArea[betInfo.BetArea] {
					personScroe += int64(Odds[int(betInfo.BetArea)]*100) * betInfo.BetScore / 100
				} else {
					bankerScroe += betInfo.BetScore
				}
			}
		}
		return true
	})
	return bankerScroe - personScroe
}

func (self *TigerXdragonGame) simulatedResult() (int64, bool) {
	//var personAwardScore int64 = self.bankerScore + self.inventory
	if self.T.Rid != SYSTEMID {
		return self.bankerScore, true
	}
	personAwardScore := self.inventory
	self.personBetInfo.Range(func(key, value any) bool {
		betInfos, ok1 := value.([]*gameMsg.TigerXdragonBetResp)
		if !ok1 {
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

	log.Infof("[%v:%v]\t模拟结算...前:%v 后:%v   当前库存:%v", self.Name, self.T.Id, self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
