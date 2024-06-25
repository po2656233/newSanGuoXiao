package brcowcow

import (
	"fmt"
	. "sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/conf"
	. "sanguoxiao/internal/component/jettengame/game/internal/category"
	. "sanguoxiao/internal/component/jettengame/manger"
	_ "sanguoxiao/internal/component/jettengame/sql/mysql" //仅仅希望导入 包内的init函数
	protoMsg "sanguoxiao/internal/protocol/gofile"
	"strconv"
	"sync"
	"time"

	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// 继承于GameItem
type BrcowcowGame struct {
	*Game

	logic       *StuCowCow
	roundNumber string

	personBetInfo *sync.Map                  //map[int64][]*protoMsg.BrcowcowBetResp //下注信息
	openAreas     [][]byte                   //开奖纪录
	openInfo      *protoMsg.BrcowcowOpenResp //开奖结果
	odds          []float64

	hostList    []int64 //申请列表(默认11个)
	bankerID    int64   //庄家ID
	superHostID int64   //超级抢庄ID
	bankerScore int64   //庄家积分
	keepTwice   uint8   //连续坐庄次数

	inventory     int64 //库存
	confInventory int64 // 记录配置上的库存,用以判断库存信息是否被更新
}

// 创建百人牛牛实例
func NewBrcowcow(game *Game) *BrcowcowGame {
	p := &BrcowcowGame{
		Game: game,
	}

	p.bankerScore = 0
	p.Config = conf.YamlObj
	p.logic = &StuCowCow{}
	p.openInfo = &protoMsg.BrcowcowOpenResp{} // 结算结果
	p.Init()
	return p
}

// ---------------------------百人牛牛----------------------------------------------//
// 初始化信息
func (self *BrcowcowGame) Init() {
	self.IsStart = true  // 是否启动
	self.IsClear = false // 不进行清场
	self.logic.Init()
	self.bankerID = 0         // 庄家ID
	self.superHostID = 0      // 超级抢庄ID
	self.keepTwice = 0        // 连续抢庄次数
	self.hostList = []int64{} // 申请列表(默认11个)

	self.personBetInfo = new(sync.Map)      // make(map[int64][]*protoMsg.BrcowcowBetResp) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 20) // 开奖纪录
	//
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
	// 玄-手牌
	self.openInfo.XuanCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}
	// 黄-手牌
	self.openInfo.HuangCard = &protoMsg.CardInfo{
		Cards: make([]byte, CardAmount),
	}
}

//获取玩家列表
//func (self *BrcowcowGame) getPlayersName() []string {
//	var name []string
//	for _, pid := range self.PlayerList {
//		log.Debug("----玩家：%v", pid)
//		name = append(name, sqlhandle.CheckName(pid))
//	}
//	return name
//}

// //////////////////////////////////////////////////////////////
// 场景信息
func (self *BrcowcowGame) Scene(args []interface{}) {
	_ = args[1]
	level := args[0].(int32)
	agent := args[1].(gate.Agent)
	userData := agent.UserData()
	if userData == nil {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game02])
		log.Debug("[Error][%v:%v场景] [未能查找到相关玩家] ", self.G.Name, self.ID)
		return
	}
	person := userData.(*Player)
	if level != self.G.Level {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game01])
		log.Debug("[Error][%v:%v场景] [玩家ID:%v 房间等级不匹配] ", self.G.Name, self.ID, person.UserID)
		return
	}

	//场景信息
	StateInfo := &protoMsg.BrcowcowSceneResp{}
	StateInfo.Inning = self.InningInfo.Number
	StateInfo.AwardAreas = self.openAreas // 录单
	StateInfo.HostID = self.bankerID
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
		bets, ok1 := value.([]*protoMsg.BrcowcowBetResp)
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
	case protoMsg.GameScene_Free:
		timeInfo.TotalTime = self.Config.Brcowcow.Duration.FreeTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrcowcowStateFreeResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Start: //抢庄
		timeInfo.TotalTime = self.Config.Brcowcow.Duration.StartTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrcowcowStateStartResp{
			Times:  timeInfo,
			HostID: self.bankerID,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = self.Config.Brcowcow.Duration.BetTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrcowcowStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = self.Config.Brcowcow.Duration.OpenTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrcowcowStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = self.Config.Brcowcow.Duration.OverTime
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.BrcowcowStateOverResp{
			Times: timeInfo,
		})
	}

}

// 开始下注前定庄
func (self *BrcowcowGame) Start(args []interface{}) {
	_ = args
	log.Release("[%v:%v]游戏創建成功", self.G.Name, self.ID)
	if self.IsStart {
		self.onDecide()
		self.IsStart = false
	}
}

// 下注 直接扣除金币
func (self *BrcowcowGame) Playing(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.BrcowcowBetReq)
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
	//log.Debug("[%v:%v] 玩家:%v Playing:->%v ", self.G.Name, self.ID, sitter.UserID, m)

	//下注成功
	//下注数目累加
	msg := &protoMsg.BrcowcowBetResp{}
	msg.UserID = sitter.UserID
	msg.BetArea = m.BetArea
	msg.BetScore = m.BetScore
	sitter.Score = m.BetScore
	sitter.Total += m.BetScore
	if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
		ok = false
		areaBetInfos, ok1 := value.([]*protoMsg.BrcowcowBetResp)
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				areaBetInfos[index].BetScore = betItem.BetScore + m.BetScore
				log.Debug("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.G.Name, self.ID, sitter.UserID, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrcowcowBetResp)
		}
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	} else {
		log.Debug("[%v:%v]\t玩家:%v 下注:%v", self.G.Name, self.ID, sitter.UserID, m)
		areaBetInfos := make([]*protoMsg.BrcowcowBetResp, 0)
		areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrcowcowBetResp)
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying
	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)

}

// 结算
func (self *BrcowcowGame) Over(args []interface{}) {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	bankerAwardScore := int64(0)
	for _, sitter := range self.GetAllSitter() {
		uid := sitter.UserID
		if uid == self.bankerID {
			continue
		}

		if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
			betInfos, ok1 := value.([]*protoMsg.BrcowcowBetResp)
			if !ok1 {
				continue
			}
			//每一次的下注信息
			for _, betInfo := range betInfos {
				log.Debug("[%v:%v]玩家:%v,下注区域:%v 下注金额:%v", self.G.Name, self.ID, uid, betInfo.BetArea, betInfo.BetScore)
				//玩家奖金
				bonusScore := self.BonusArea(betInfo.BetArea, betInfo.BetScore)
				sitter.Gain += bonusScore
				bankerAwardScore -= bonusScore
				if bonusScore < 0 {
					allAreaInfo[betInfo.BetArea] -= betInfo.BetScore
				} else {
					allAreaInfo[betInfo.BetArea] += betInfo.BetScore
				}

			}
		}
	}

	//派奖
	checkout := &protoMsg.BrcowcowOverResp{}

	if self.bankerID != INVALID {
		//查看是否够赔
		if sitter, ok := self.GetSitter(self.bankerID); ok {
			if sitter.Score+bankerAwardScore < INVALID {
				bankerAwardScore = -sitter.Score
			}
			sitter.Gain = bankerAwardScore
		}
	}
	checkout.TotalSettlement = allAreaInfo[:]
	checkout.TotalSettlement = append(checkout.TotalSettlement, bankerAwardScore)

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
func (self *BrcowcowGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]
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
	case GameUpdateHostList:
		self.hostlist(args[2:])
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
func (self *BrcowcowGame) SuperControl(args []interface{}) {
	_ = args[0]
	flag := args[0].(int32)
	switch flag {
	case GameControlClear:
		self.IsClear = true
	}
}

// 发牌
func (self *BrcowcowGame) DispatchCard() {
	//6张牌
	timeoutCount := 0
Dispatch:
	//先获取 四家的牌值,然后进行比对
	cards := Shuffle(CardListData[0:])

	self.openInfo.TianCard.Cards = Deal(cards, CardAmount, IndexStart+Area_Tian)
	self.openInfo.DiCard.Cards = Deal(cards, CardAmount, IndexStart+Area_Di)
	self.openInfo.XuanCard.Cards = Deal(cards, CardAmount, IndexStart+Area_Xuan)
	self.openInfo.HuangCard.Cards = Deal(cards, CardAmount, IndexStart+Area_Huang)
	self.openInfo.BankerCard.Cards = Deal(cards, CardAmount, IndexStart+Area_Banker)

	//	log.Debug("[百人牛牛]-->开始发牌啦<---  \n牌堆中取牌:%v ", GetCardsText(cards))

	log.Debug("[%v:%v]\t各家牌值：庄家:%v \t天:%v\t玄:%v\t地:%v\t黄:%v", self.G.Name, self.ID, GetCardsText(self.openInfo.BankerCard.Cards),
		GetCardsText(self.openInfo.TianCard.Cards), GetCardsText(self.openInfo.DiCard.Cards),
		GetCardsText(self.openInfo.XuanCard.Cards), GetCardsText(self.openInfo.HuangCard.Cards))

	var banker = Pokers{}
	var tian = Pokers{}
	var xuan = Pokers{}
	var di = Pokers{}
	var huang = Pokers{}

	list := make([]int, len(cards))
	for k, v := range cards {
		list[k] = int(v)
	}
	pokerList := CreatePoker(list)

	tian.AddPokers(pokerList[0], pokerList[1], pokerList[2], pokerList[3], pokerList[4]).ArrangeByNumber()
	di.AddPokers(pokerList[5], pokerList[6], pokerList[7], pokerList[8], pokerList[9]).ArrangeByNumber()
	xuan.AddPokers(pokerList[10], pokerList[11], pokerList[12], pokerList[13], pokerList[14]).ArrangeByNumber()
	huang.AddPokers(pokerList[15], pokerList[16], pokerList[17], pokerList[18], pokerList[19]).ArrangeByNumber()
	banker.AddPokers(pokerList[20], pokerList[21], pokerList[22], pokerList[23], pokerList[24]).ArrangeByNumber()

	tianType := CalcPoker(tian)
	xuanType := CalcPoker(xuan)
	diType := CalcPoker(di)
	huangType := CalcPoker(huang)
	bankerType := CalcPoker(banker)

	if self.logic.BetWinInfoMap == nil {
		self.logic.BetWinInfoMap = &BetWinInfoMap{}
	}

	if self.logic.BetWinInfoMap.WinInfoMap == nil {
		self.logic.BetWinInfoMap.Init()
	}

	// 清空当前彩源的开奖信息
	self.roundNumber = strconv.FormatInt(self.TimeStamp, 10)
	self.logic.BetWinInfoMap.InitSourceMap(self.roundNumber)

	//各个区域的开奖结果
	self.logic.BetWinInfoMap.Set(self.roundNumber, Area_Tian, self.logic.Compare(Area_Tian, bankerType, tianType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, Area_Di, self.logic.Compare(Area_Di, bankerType, diType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, Area_Xuan, self.logic.Compare(Area_Xuan, bankerType, xuanType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, Area_Huang, self.logic.Compare(Area_Huang, bankerType, huangType))

	//开奖结果
	self.openInfo.BankerCard.CardValue = int32(bankerType.Type)
	self.openInfo.TianCard.CardValue = int32(tianType.Type)
	self.openInfo.DiCard.CardValue = int32(diType.Type)
	self.openInfo.XuanCard.CardValue = int32(xuanType.Type)
	self.openInfo.HuangCard.CardValue = int32(huangType.Type)
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

	GlobalSender.NotifyOthers(self.PlayerList, self.openInfo)
	log.Debug("[%v:%v]\t庄家牌值:Banker:%v,T:%v D:%v X:%v H:%v 中奖区域:%v ", self.G.Name, self.ID,
		self.openInfo.BankerCard.CardValue,
		self.openInfo.TianCard.CardValue,
		self.openInfo.DiCard.CardValue,
		self.openInfo.XuanCard.CardValue,
		self.openInfo.HuangCard.CardValue,
		self.openInfo.AwardArea)

}

// 开奖区域
func (self *BrcowcowGame) DeduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	self.odds = make([]float64, AREA_MAX)

	betTianInfo, _ := self.logic.BetWinInfoMap.Get(self.roundNumber, Area_Tian)
	betDiInfo, _ := self.logic.BetWinInfoMap.Get(self.roundNumber, Area_Di)
	betXuanInfo, _ := self.logic.BetWinInfoMap.Get(self.roundNumber, Area_Xuan)
	betHuangInfo, _ := self.logic.BetWinInfoMap.Get(self.roundNumber, Area_Huang)

	if betTianInfo.IsWin {
		pWinArea[Area_Tian] = Win
		self.odds[Area_Tian] = betTianInfo.WinOdds
		log.Debug("[%v:%v] 天赢:%v ->%v", self.G.Name, self.ID, betTianInfo.WinOdds, betTianInfo.LoseOdds)
	} else {
		self.odds[Area_Tian] = 0.0
	}

	if betDiInfo.IsWin {
		pWinArea[Area_Di] = Win
		self.odds[Area_Di] = betDiInfo.WinOdds
		log.Debug("[%v:%v] 地赢:%v ->%v", self.G.Name, self.ID, betDiInfo.WinOdds, betDiInfo.LoseOdds)
	} else {
		self.odds[Area_Di] = 0.0
	}

	if betXuanInfo.IsWin {
		pWinArea[Area_Xuan] = Win
		self.odds[Area_Xuan] = betXuanInfo.WinOdds
		log.Debug("[%v:%v] 玄赢:%v ->%v", self.G.Name, self.ID, betXuanInfo.WinOdds, betXuanInfo.LoseOdds)
	} else {
		self.odds[Area_Xuan] = 0.0
	}

	if betHuangInfo.IsWin {
		pWinArea[Area_Huang] = Win
		self.odds[Area_Huang] = betHuangInfo.WinOdds
		log.Debug("[%v:%v] 黄赢:%v ->%v", self.G.Name, self.ID, betHuangInfo.WinOdds, betHuangInfo.LoseOdds)
	} else {
		self.odds[Area_Huang] = 0.0
	}

	//庄家区域
	return pWinArea
}

// 区域赔额
func (self *BrcowcowGame) BonusArea(area int32, betScore int64) int64 {
	if Win == self.openInfo.AwardArea[area] {
		bonus := int64(self.odds[int(area)]*100) * betScore / 100
		return bonus
	}
	return -betScore
}

// 定庄
func (self *BrcowcowGame) onDecide() {
	self.Reset()
	// 校准库存
	if self.confInventory != self.Config.Brcowcow.Inventory {
		self.confInventory = self.Config.Brcowcow.Inventory
		//此处会重置以往的库存值
		self.inventory = self.confInventory
	}

	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(self.Config.Brcowcow.Duration.StartTime)*time.Second, self.onPlay)

	if INVALID < self.bankerID {
		self.keepTwice++
	}
	if self.Config.Brcowcow.MaxHost < self.keepTwice {
		self.keepTwice = 0
		self.bankerID = INVALID
	}

	if self.bankerID == INVALID {
		isHave := false
		size := len(self.hostList)
		for i := 0; i < size; i++ {
			for _, pid := range self.PlayerList {
				//
				if pid == self.hostList[i] {
					person := GlobalPlayerManger.Get(pid)
					if person != nil && self.Config.Brcowcow.BankerScore <= person.Gold {
						self.bankerID = self.hostList[i]
						self.bankerScore = person.Gold
						if sit, ok := self.GetSitter(self.bankerID); ok {
							sit.Score = person.Gold
						}
						isHave = true
						break
					}

				}
			}
			if isHave {
				GlobalSender.SendResultX(self.bankerID, SUCCESS, StatusText[User17])
				break
			}
		}
	} else {
		person := GlobalPlayerManger.Get(self.bankerID)
		if person != nil {
			if self.Config.Brcowcow.BankerScore <= person.Gold {
				self.bankerScore = person.Gold
				if sit, ok := self.GetSitter(self.bankerID); ok {
					sit.Score = person.Gold
				}
			} else {
				self.bankerScore = 0
				self.bankerID = INVALID
				GlobalSender.SendResultX(person.UserID, FAILED, StatusText[Game46])

			}
		}
	}
	self.hostList = DeleteValue(self.hostList, self.bankerID).([]int64)
	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brcowcow.Duration.StartTime,
			TotalTime: self.Config.Brcowcow.Duration.StartTime,
		},
		HostID: self.bankerID,
		Inning: self.InningInfo.Number,
	})

}

// [下注减法] //开始|下注|结算|空闲
func (self *BrcowcowGame) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	time.AfterFunc(time.Duration(self.Config.Brcowcow.Duration.BetTime)*time.Second, self.onOpen)

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brcowcow.Duration.BetTime,
			TotalTime: self.Config.Brcowcow.Duration.BetTime,
		},
	})
}

func (self *BrcowcowGame) onOpen() {
	self.ChangeState(protoMsg.GameScene_Opening)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brcowcow.Duration.OpenTime,
			TotalTime: self.Config.Brcowcow.Duration.OpenTime,
		},
	})

	//【发牌】
	self.DispatchCard()
	time.AfterFunc(time.Duration(self.Config.Brcowcow.Duration.OpenTime)*time.Second, self.onOver)
}

// [结算加法]
func (self *BrcowcowGame) onOver() {
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
	time.AfterFunc(time.Duration(self.Config.Brcowcow.Duration.OverTime)*time.Second, self.onDecide)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  self.Config.Brcowcow.Duration.OverTime,
			TotalTime: self.Config.Brcowcow.Duration.OverTime,
		},
	})
}

// 抢庄
func (self *BrcowcowGame) host(args []interface{}) {
	//【消息】
	a := args[0].([]interface{})
	agent := a[1].(gate.Agent)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	host := a[0].(*protoMsg.BrcowcowHostReq)
	person := userData.(*Player)
	userID := person.UserID

	size := len(self.hostList)
	if self.bankerID == userID && host.IsWant {
		if self.G.Scene == protoMsg.GameScene_Playing {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game26])
		} else {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game26])
		}

		return
	}
	if person.Gold < self.Config.Brcowcow.BankerScore {
		GlobalSender.SendResult(agent, FAILED, fmt.Sprintf("坐庄金币不得小于%v", self.Config.Brcowcow.BankerScore))
		return
	}

	for _, pid := range self.hostList {
		if pid == userID {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game25])
			return
		}
	}

	if host.IsWant { //申请(已经是庄家)
		//列表上是否已经申请了
		for _, pid := range self.hostList {
			if pid == userID {
				GlobalSender.SendResult(agent, FAILED, StatusText[Game25])
				return
			}
		}

		if BrcowcowHostMax < size {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
		person.State = protoMsg.PlayerState_PlayerCall
		if sit, ok := self.GetSitter(userID); ok {
			sit.State = person.State
		}
		self.hostList = CopyInsert(self.hostList, size, userID).([]int64)
		GlobalSender.SendResultX(userID, SUCCESS, StatusText[User18])
	} else { //取消
		self.hostList = SliceRemoveDuplicate(self.hostList).([]int64)
		self.hostList = DeleteValue(self.hostList, userID).([]int64)
		if self.bankerID == userID {
			if self.G.Scene == protoMsg.GameScene_Playing {
				GlobalSender.SendResultX(userID, SUCCESS, StatusText[Game28])
				return
			} else if self.G.Scene == protoMsg.GameScene_Start {
				GlobalSender.SendResultX(userID, SUCCESS, StatusText[Game29])
				return
			}
			self.bankerID = INVALID
			self.bankerScore = INVALID
		}
		GlobalSender.SendResultX(userID, SUCCESS, StatusText[User16])
	}

	log.Debug("[%v:%v]有人来抢庄啦:%d 列表人数%d", self.G.Name, self.ID, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 抢庄
func (self *BrcowcowGame) hostlist(args []interface{}) {
	//【消息】
	a := args[0].([]interface{})
	agent := a[1].(gate.Agent)

	person := GlobalPlayerManger.Get(self.bankerID)
	msg := &protoMsg.BrcowcowHostListResp{}
	if person != nil {
		msg.CurHost = person.ToMsg()
	} else {
		//msg.CurHost = &protoMsg.PlayerInfo{
		//    UserID: INVALID,
		//    Name: "系统",
		//}
	}
	msg.Waitlist = self.hostList
	GlobalSender.SendData(agent, msg)
}

// 超级抢庄
func (self *BrcowcowGame) superHost(args []interface{}) {
	_ = args
	////【消息】
	//_ = args[2]
	//agent := args[1].(gate.Agent)
	//host := args[2].(*protoMsg.BrcowcowSuperHostReq)
	//
	//userData := agent.UserData()
	//if nil == userData {
	//	GlobalSender.SendResult(agent, FAILED, StatusText[User02])
	//	return
	//}
	//user := userData.(*Player)
	//if user.Level < LessLevel {
	//	GlobalSender.SendResult(agent, FAILED, StatusText[User11])
	//	return
	//}
	//userID := user.UserID
	//log.Debug("[百人牛牛]有人要超级抢庄--->:%d", userID)
	//
	//if host.IsWant {
	//	if self.superHostID == 0 {
	//		self.superHostID = userID
	//		//超级抢庄放申请列表首位
	//		self.hostList = CopyInsert(self.hostList, 0, userID).([]int64)
	//	} else {
	//		GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
	//		return
	//	}
	//}
	//
	//// 通知
	//GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowSuperHostResp{
	//	UserID: self.bankerID,
	//	IsWant: host.IsWant,
	//})

}

// 定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *BrcowcowGame) permitHost() *protoMsg.BrcowcowHostResp {
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

		banker := GlobalPlayerManger.Get(self.bankerID)
		if banker == nil {
			return nil
		}
		banker.State = protoMsg.PlayerState_PlayerPlaying
		self.bankerScore = banker.Gold
		if sit, ok := self.GetSitter(self.bankerID); ok {
			sit.State = protoMsg.PlayerState_PlayerPlaying
			sit.Score = banker.Gold
		}

		log.Debug("[%v:%v]确定庄家:%d", self.G.Name, self.ID, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		//self.bankerScore = 1000000
		log.Debug("[%v:%v]系统坐庄", self.G.Name, self.ID)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &protoMsg.BrcowcowHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Debug("[%v:%v]广播上庄", self.G.Name, self.ID)
	return msg
}

// -----------------------逻辑层---------------------------
// 重新初始化[返回结果提供给外部清场用]
func (self *BrcowcowGame) Reset() int64 {
	self.openInfo.BankerCard.Cards = make([]byte, CardAmount)
	self.openInfo.TianCard.Cards = make([]byte, CardAmount)
	self.openInfo.DiCard.Cards = make([]byte, CardAmount)
	self.openInfo.XuanCard.Cards = make([]byte, CardAmount)
	self.openInfo.HuangCard.Cards = make([]byte, CardAmount)
	//玩家下注信息
	self.personBetInfo.Range(func(key, value any) bool {
		self.personBetInfo.Delete(key)
		return true
	})
	return self.Game.Reset()
}

func (self *BrcowcowGame) simulatedResult() (int64, bool) {
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
		betInfos, ok1 := value.([]*protoMsg.BrcowcowBetResp)
		if !ok1 {
			return true
		}
		if GetPlayerManger().CheckRobot(uid) {
			return true
		}
		for _, betInfo := range betInfos {
			//玩家奖金
			if Win == self.openInfo.AwardArea[betInfo.BetArea] {
				personAwardScore -= self.BonusArea(betInfo.BetArea, betInfo.BetScore) + betInfo.BetScore
			}
		}
		return true
	})
	//
	log.Debug("[百人牛牛]模拟结算...前:%v 后:%v   当前库存:%v", self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
