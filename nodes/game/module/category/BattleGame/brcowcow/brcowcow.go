package brcowcow

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	"strconv"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	. "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
	"sync"
	"time"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// BrcowcowGame 继承于GameItem
type BrcowcowGame struct {
	*Game
	T *Table

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

// New 创建百人牛牛实例
func New(game *Game, tb *Table) *BrcowcowGame {
	p := &BrcowcowGame{
		Game: game,
		T:    tb,
	}

	p.bankerScore = 0
	p.logic = &StuCowCow{}
	p.openInfo = &protoMsg.BrcowcowOpenResp{} // 结算结果
	p.Init()
	return p
}

// ---------------------------百人牛牛----------------------------------------------//

// Init 初始化信息
func (self *BrcowcowGame) Init() {
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

// //////////////////////////////////////////////////////////////
// 场景信息
func (self *BrcowcowGame) Scene(args []interface{}) bool {
	if !self.Game.Scene(args) {
		return false
	}
	person := args[0].(*Player)
	//场景信息
	StateInfo := &protoMsg.BrcowcowSceneResp{
		AllPlayers: &protoMsg.PlayerList{
			Items: make([]*protoMsg.PlayerInfo, 0),
		},
		Inning:     self.Inning,
		TimeStamp:  self.TimeStamp,
		AwardAreas: self.openAreas,
		AreaBets:   make([]int64, AREA_MAX),
		MyBets:     make([]int64, AREA_MAX),
	}

	log.Infof("[%v:%v] [Scene:%v] [当前人数:%v] ", self.GameInfo.Name, self.T.Id, self.GameInfo.Scene, self.T.SitCount())
	self.T.ChairWork(func(chair *Chair) {
		StateInfo.AllPlayers.Items = append(StateInfo.AllPlayers.Items, chair.PlayerInfo)
	})
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
	uid := person.UserID
	GlobalSender.SendTo(uid, StateInfo)

	//通知游戏状态
	timeInfo := &protoMsg.TimeInfo{
		TimeStamp: self.TimeStamp,
		OutTime:   int32(time.Now().Unix() - self.TimeStamp),
	}
	switch self.GameInfo.Scene {
	case protoMsg.GameScene_Free:
		timeInfo.TotalTime = YamlObj.BrCowcow.Duration.Free
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrcowcowStateFreeResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Start: //抢庄
		timeInfo.TotalTime = YamlObj.BrCowcow.Duration.Start
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrcowcowStateStartResp{
			Times:  timeInfo,
			HostID: self.bankerID,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = YamlObj.BrCowcow.Duration.Play
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrcowcowStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = YamlObj.BrCowcow.Duration.Open
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrcowcowStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = YamlObj.BrCowcow.Duration.Over
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrcowcowStateOverResp{
			Times: timeInfo,
		})
	}
	return true
}

// 开始下注前定庄
func (self *BrcowcowGame) Start(args []interface{}) bool {
	_ = args
	if !self.IsStart {
		log.Infof("[%v] [%v:%v]  游戏成功启动!!! ", self.T.Name, self.GameInfo.Name, self.T.Id)
		self.onDecide()
		self.IsStart = true
	}
	return true
}

// 下注 直接扣除金币
func (self *BrcowcowGame) Playing(args []interface{}) bool {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.BrcowcowBetReq)
	//【传输对象】
	agent := args[1].(Agent)

	userData := agent.UserData()
	if m.BetScore <= 0 {
		log.Infof("[%v:%v][%v] Playing:->%v ", self.Name, self.T.Id, StatusText[Game07], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return false
	}

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
	//log.Infof("[%v:%v] 玩家:%v Playing:->%v ", self.Name, self.T.Id, sitter.UserID, m)

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
				log.Infof("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.Name, self.T.Id, sitter.UserID, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = utils.CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrcowcowBetResp)
		}
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	} else {
		log.Infof("[%v:%v]\t玩家:%v 下注:%+v", self.Name, self.T.Id, sitter.UserID, m)
		areaBetInfos := make([]*protoMsg.BrcowcowBetResp, 0)
		areaBetInfos = utils.CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrcowcowBetResp)
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying
	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)

	return true
}

// 结算
func (self *BrcowcowGame) Over(args []interface{}) {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	bankerAwardScore := int64(0)
	self.personBetInfo.Range(func(key, value any) bool {
		uid := key.(int64)
		if uid == self.bankerID {
			return true
		}

		betInfos, ok1 := value.([]*protoMsg.BrcowcowBetResp)
		if !ok1 {
			return true
		}
		sitter := self.T.GetChair(uid)
		//每一次的下注信息
		for _, betInfo := range betInfos {
			log.Infof("[%v:%v]玩家:%v,下注区域:%v 下注金额:%v", self.Name, self.T.Id, uid, betInfo.BetArea, betInfo.BetScore)
			//玩家奖金
			bonusScore := self.BonusArea(betInfo.BetArea, betInfo.BetScore)
			if sitter != nil {
				sitter.Gain += bonusScore
			}
			bankerAwardScore -= bonusScore
			if bonusScore < 0 {
				allAreaInfo[betInfo.BetArea] -= betInfo.BetScore
			} else {
				allAreaInfo[betInfo.BetArea] += betInfo.BetScore
			}

		}
		return true
	})

	//派奖
	checkout := &protoMsg.BrcowcowOverResp{}

	if self.bankerID != INVALID {
		//查看是否够赔
		if sitter := self.T.GetChair(self.bankerID); sitter != nil {
			if sitter.Score+bankerAwardScore < INVALID {
				bankerAwardScore = -sitter.Score
			}
			sitter.Gain = bankerAwardScore
		}
	}
	// 结算
	result := fmt.Sprintf("Banker[%v] T[%v] D[%v] X[%v] H[%v] 中奖区域:%v",
		self.openInfo.BankerCard.CardValue,
		self.openInfo.TianCard.CardValue,
		self.openInfo.DiCard.CardValue,
		self.openInfo.XuanCard.CardValue,
		self.openInfo.HuangCard.CardValue,
		self.openInfo.AwardArea)
	self.T.ChairSettle(self.Name, self.Inning, result)

	// 派奖信息
	checkout.TotalSettlement = allAreaInfo[:]
	checkout.TotalSettlement = append(checkout.TotalSettlement, bankerAwardScore)
	self.T.ChairWork(func(chair *Chair) {
		if chair.Gain != INVALID {
			checkout.MyAcquire = chair.Gain
			GlobalSender.SendTo(chair.UserID, checkout)
		}
	})

	log.Infof("[%v:%v]   \t结算注单... 各区域情况:%v", self.Name, self.T.Id, allAreaInfo)
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

	//	log.Infof("[百人牛牛]-->开始发牌啦<---  \n牌堆中取牌:%v ", GetCardsText(cards))

	log.Infof("[%v:%v]\t各家牌值：庄家:%v \t天:%v\t玄:%v\t地:%v\t黄:%v", self.Name, self.T.Id, GetCardsText(self.openInfo.BankerCard.Cards),
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

	GlobalSender.NotifyOthers(self.PlayerList, self.openInfo)
	log.Infof("[%v:%v]\t庄家牌值:Banker:%v,T:%v D:%v X:%v H:%v 中奖区域:%v ", self.Name, self.T.Id,
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
		log.Infof("[%v:%v] 天赢:%v ->%v", self.Name, self.T.Id, betTianInfo.WinOdds, betTianInfo.LoseOdds)
	} else {
		self.odds[Area_Tian] = 0.0
	}

	if betDiInfo.IsWin {
		pWinArea[Area_Di] = Win
		self.odds[Area_Di] = betDiInfo.WinOdds
		log.Infof("[%v:%v] 地赢:%v ->%v", self.Name, self.T.Id, betDiInfo.WinOdds, betDiInfo.LoseOdds)
	} else {
		self.odds[Area_Di] = 0.0
	}

	if betXuanInfo.IsWin {
		pWinArea[Area_Xuan] = Win
		self.odds[Area_Xuan] = betXuanInfo.WinOdds
		log.Infof("[%v:%v] 玄赢:%v ->%v", self.Name, self.T.Id, betXuanInfo.WinOdds, betXuanInfo.LoseOdds)
	} else {
		self.odds[Area_Xuan] = 0.0
	}

	if betHuangInfo.IsWin {
		pWinArea[Area_Huang] = Win
		self.odds[Area_Huang] = betHuangInfo.WinOdds
		log.Infof("[%v:%v] 黄赢:%v ->%v", self.Name, self.T.Id, betHuangInfo.WinOdds, betHuangInfo.LoseOdds)
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
	self.reset()
	// 校准库存
	if self.confInventory != YamlObj.BrCowcow.Inventory {
		self.confInventory = YamlObj.BrCowcow.Inventory
		//此处会重置以往的库存值
		self.inventory = self.confInventory
	}

	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(YamlObj.BrCowcow.Duration.Start)*time.Second, self.onPlay)

	if INVALID < self.bankerID {
		self.keepTwice++
	}
	if YamlObj.BrCowcow.MaxHost < self.keepTwice {
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
					person := GlobalPlayerMgr.Get(pid)
					if person != nil && YamlObj.BrCowcow.BankerScore <= person.Gold {
						self.bankerID = self.hostList[i]
						self.bankerScore = person.Gold
						if sit := self.T.GetChair(self.bankerID); sit != nil {
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
		person := GlobalPlayerMgr.Get(self.bankerID)
		if person != nil {
			if YamlObj.BrCowcow.BankerScore <= person.Gold {
				self.bankerScore = person.Gold
				if sit := self.T.GetChair(self.bankerID); sit != nil {
					sit.Score = person.Gold
				}
			} else {
				self.bankerScore = 0
				self.bankerID = INVALID
				GlobalSender.SendResultX(person.UserID, FAILED, StatusText[Game46])

			}
		}
	}
	self.hostList = utils.RemoveValue(self.hostList, self.bankerID)
	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.BrCowcow.Duration.Start,
			TotalTime: YamlObj.BrCowcow.Duration.Start,
		},
		HostID: self.bankerID,
		Inning: self.Inning,
	})

}

// [下注减法] //开始|下注|结算|空闲
func (self *BrcowcowGame) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	time.AfterFunc(time.Duration(YamlObj.BrCowcow.Duration.Play)*time.Second, self.onOpen)

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.BrCowcow.Duration.Play,
			TotalTime: YamlObj.BrCowcow.Duration.Play,
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
			WaitTime:  YamlObj.BrCowcow.Duration.Open,
			TotalTime: YamlObj.BrCowcow.Duration.Open,
		},
	})

	//【发牌】
	self.DispatchCard()
	time.AfterFunc(time.Duration(YamlObj.BrCowcow.Duration.Open)*time.Second, self.onOver)
}

// [结算加法]
func (self *BrcowcowGame) onOver() {
	self.ChangeState(protoMsg.GameScene_Over)

	// 玩家结算(框架消息)
	self.Over(nil)

	self.T.CalibratingRemain(Default)

	//自动清场
	if self.IsClear || (SYSTEMID != self.T.Rid && self.T.Remain <= INVALID) {
		self.Close(func() *Table {
			return self.T
		})
		return
	}

	// 开奖状态
	time.AfterFunc(time.Duration(YamlObj.BrCowcow.Duration.Over)*time.Second, self.onDecide)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  YamlObj.BrCowcow.Duration.Over,
			TotalTime: YamlObj.BrCowcow.Duration.Over,
		},
	})
}

// 抢庄
func (self *BrcowcowGame) host(args []interface{}) {
	//【消息】
	a := args[0].([]interface{})
	agent := a[1].(Agent)

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
		if self.GameInfo.Scene == protoMsg.GameScene_Playing {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game26])
		} else {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game26])
		}

		return
	}
	if person.Gold < YamlObj.BrCowcow.BankerScore {
		GlobalSender.SendResult(agent, FAILED, fmt.Sprintf("坐庄金币不得小于%v", YamlObj.BrCowcow.BankerScore))
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
		if sit := self.T.GetChair(userID); sit != nil {
			sit.State = person.State
		}
		self.hostList = utils.CopyInsert(self.hostList, size, userID).([]int64)
		GlobalSender.SendResultX(userID, SUCCESS, StatusText[User18])
	} else { //取消
		self.hostList = utils.Unique(self.hostList)
		self.hostList = utils.RemoveValue(self.hostList, userID)
		if self.bankerID == userID {
			if self.GameInfo.Scene == protoMsg.GameScene_Playing {
				GlobalSender.SendResultX(userID, SUCCESS, StatusText[Game28])
				return
			} else if self.GameInfo.Scene == protoMsg.GameScene_Start {
				GlobalSender.SendResultX(userID, SUCCESS, StatusText[Game29])
				return
			}
			self.bankerID = INVALID
			self.bankerScore = INVALID
		}
		GlobalSender.SendResultX(userID, SUCCESS, StatusText[User16])
	}

	log.Infof("[%v:%v]有人来抢庄啦:%d 列表人数%d", self.Name, self.T.Id, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrcowcowHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 抢庄
func (self *BrcowcowGame) hostlist(args []interface{}) {
	//【消息】
	_ = args
	//a := args[0].([]interface{})
	person := GlobalPlayerMgr.Get(self.bankerID)
	msg := &protoMsg.BrcowcowHostListResp{}
	if person != nil {
		msg.CurHost = person.PlayerInfo
		msg.Waitlist = self.hostList
		GlobalSender.SendTo(person.UserID, msg)
	}

}

// 超级抢庄
func (self *BrcowcowGame) superHost(args []interface{}) {
	_ = args
	////【消息】
	//_ = args[2]
	//agent := args[1].(Agent)
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
	//log.Infof("[百人牛牛]有人要超级抢庄--->:%d", userID)
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
	log.Infof("[%v:%v]定庄.... 列表数据:%v", self.Name, self.T.Id, self.hostList)

	//befBankerID := self.bankerID 避免重复
	for index, pid := range tempList {
		person := GlobalPlayerMgr.Get(pid)
		if person == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == person.UserID {
			log.Infof("[%v:%v]不再连续坐庄：%v", self.Name, self.T.Id, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if person.Gold < 5000 {
			log.Infof("[%v:%v]玩家%d 金币%lf 少于5000不能申请坐庄", self.Name, self.T.Id, person.UserID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Infof("[%v:%v]连续坐庄次数:%d", self.Name, self.T.Id, self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]

		banker := GlobalPlayerMgr.Get(self.bankerID)
		if banker == nil {
			return nil
		}
		banker.State = protoMsg.PlayerState_PlayerPlaying
		self.bankerScore = banker.Gold
		if sit := self.T.GetChair(self.bankerID); sit != nil {
			sit.State = protoMsg.PlayerState_PlayerPlaying
			sit.Score = banker.Gold
		}

		log.Infof("[%v:%v]确定庄家:%d", self.Name, self.T.Id, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		//self.bankerScore = 1000000
		log.Infof("[%v:%v]系统坐庄", self.Name, self.T.Id)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &protoMsg.BrcowcowHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Infof("[%v:%v]广播上庄", self.Name, self.T.Id)
	return msg
}

// -----------------------逻辑层---------------------------
// reset 重新初始化[返回结果提供给外部清场用]
func (self *BrcowcowGame) reset() {
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
	self.Game.Init()
	return
}

func (self *BrcowcowGame) simulatedResult() (int64, bool) {
	//var personAwardScore int64 = self.bankerScore + self.inventory
	if self.T.Rid != SYSTEMID {
		return self.bankerScore, true
	}
	var personAwardScore int64 = self.inventory
	self.personBetInfo.Range(func(key, value any) bool {
		betInfos, ok := value.([]*protoMsg.BrcowcowBetResp)
		if !ok {
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
	log.Infof("[百人牛牛]模拟结算...前:%v 后:%v   当前库存:%v", self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
