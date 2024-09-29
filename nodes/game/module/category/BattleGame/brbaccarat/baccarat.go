package brbaccarat

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/go_file/common"
	//gateMsg "superman/internal/protocol/go_file/gate"
	gameMsg "superman/internal/protocol/go_file/game"
	"superman/internal/utils"
	. "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
	"sync"
	"time"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// BaccaratGame 继承于GameItem
type BaccaratGame struct {
	*Game
	T *Table

	cbPlayerCards *protoMsg.CardInfo //闲家扑克
	cbBankerCards *protoMsg.CardInfo //庄家扑克
	cbCardCount   []byte             //首次拿牌数

	personBetInfo *sync.Map                 // map[int64][]*protoMsg.BaccaratBetResp //下注信息
	openAreas     [][]byte                  //开奖纪录
	openInfo      *gameMsg.BaccaratOpenResp //开奖结果

	hostList    []int64 //申请列表(默认11个)
	bankerID    int64   //庄家ID
	superHostID int64   //超级抢庄ID
	bankerScore int64   //庄家积分
	keepTwice   uint8   //连续坐庄次数

	inventory     int64 //库存
	confInventory int64 // 记录配置上的库存,用以判断库存信息是否被更新
}

// New 创建百家乐实例
func New(game *Game, tb *Table) *BaccaratGame {
	p := &BaccaratGame{
		Game: game,
		T:    tb,
	}
	p.bankerScore = 0
	p.Init()
	return p
}

// ---------------------------百家乐----------------------------------------------//

// Init 初始化信息
func (self *BaccaratGame) Init() {
	self.bankerID = 0         // 庄家ID
	self.superHostID = 0      // 超级抢庄ID
	self.keepTwice = 0        // 连续抢庄次数
	self.hostList = []int64{} // 申请列表(默认11个)

	self.cbCardCount = []byte{2, 2}             // 首次拿牌张数
	self.openInfo = &gameMsg.BaccaratOpenResp{} // 结算结果

	self.personBetInfo = new(sync.Map)      // make(map[int64][]*protoMsg.BaccaratBetResp) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 20) // 开奖纪录
	// 闲家手牌
	self.cbPlayerCards = &protoMsg.CardInfo{
		Cards: make([]byte, CardCountPlayer),
	}
	// 庄家手牌
	self.cbBankerCards = &protoMsg.CardInfo{
		Cards: make([]byte, CardCountBanker),
	}
}

//获取玩家列表
//func (self *BaccaratGame) getPlayersName() []string {
//	var name []string
//	for _, pid := range self.PlayerList {
//		log.Debugf("----玩家：%v", pid)
//		name = append(name, sqlhandle.CheckName(pid))
//	}
//	return name
//}

// //////////////////////////////////////////////////////////////

// Scene 场景信息
func (self *BaccaratGame) Scene(args []interface{}) bool {
	if !self.Game.Scene(args) {
		return false
	}
	person := args[0].(*Player)

	//场景信息
	StateInfo := &gameMsg.BaccaratSceneResp{
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
		bets, ok1 := value.([]*gameMsg.BaccaratBetResp)
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
	GlobalSender.SendTo(person.UserID, StateInfo)

	//通知游戏状态
	timeInfo := &protoMsg.TimeInfo{
		TimeStamp: self.TimeStamp,
		OutTime:   int32(time.Now().Unix() - self.TimeStamp),
	}
	switch self.GameInfo.Scene {
	case protoMsg.GameScene_Start: //准备
		timeInfo.TotalTime = YamlObj.Baccarat.Duration.Start
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(person.UserID, &gameMsg.BaccaratStateStartResp{
			Times:  timeInfo,
			Inning: self.Inning,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = YamlObj.Baccarat.Duration.Play
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(person.UserID, &gameMsg.BaccaratStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = YamlObj.Baccarat.Duration.Open
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(person.UserID, &gameMsg.BaccaratStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = YamlObj.Baccarat.Duration.Over
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(person.UserID, &gameMsg.BaccaratStateOverResp{
			Times: timeInfo,
		})
	}
	return true
}

// Start 开始下注前定庄
func (self *BaccaratGame) Start(args []interface{}) bool {
	_ = args

	if !self.IsStart {
		log.Infof("[%v] [%v:%v]  游戏成功启动!!! ", self.T.Name, self.GameInfo.Name, self.T.Id)
		self.IsStart = true
		// 游戏开始事件
		// 开始状态
		startResp := &gameMsg.BaccaratStateStartResp{
			Times: &protoMsg.TimeInfo{
				OutTime:   0,
				WaitTime:  YamlObj.Baccarat.Duration.Start,
				TotalTime: YamlObj.Baccarat.Duration.Start,
			},
		}
		self.Game.RegisterEvent(protoMsg.GameScene_Start, startResp, startResp.Times.TotalTime, func() bool {
			// 校准库存
			if self.confInventory != YamlObj.Baccarat.Inventory {
				self.confInventory = YamlObj.Baccarat.Inventory
				//此处会重置以往的库存值
				self.inventory = self.confInventory
			}

			return true
		}, func() {
			startResp.Inning = self.Inning
			startResp.Times.TimeStamp = self.TimeStamp
			self.reset()
		}, nil)

		// 游戏中
		playResp := &gameMsg.BaccaratStatePlayingResp{
			Times: &protoMsg.TimeInfo{
				OutTime:   0,
				WaitTime:  YamlObj.Baccarat.Duration.Play,
				TotalTime: YamlObj.Baccarat.Duration.Play,
			},
		}
		self.Game.RegisterEvent(protoMsg.GameScene_Playing, playResp, playResp.Times.TotalTime, func() bool {
			return true
		}, func() {
			playResp.Times.TimeStamp = self.TimeStamp
		}, nil)

		// 开牌
		openResp := &gameMsg.BaccaratStateOpenResp{
			Times: &protoMsg.TimeInfo{
				OutTime:   0,
				WaitTime:  YamlObj.Baccarat.Duration.Open,
				TotalTime: YamlObj.Baccarat.Duration.Open,
			},
		}
		self.Game.RegisterEvent(protoMsg.GameScene_Opening, openResp, openResp.Times.TotalTime, func() bool {
			//【发牌】
			self.DispatchCard()
			return true
		}, func() {
			openResp.Times.TimeStamp = self.TimeStamp
		}, nil)

		// 结算
		overResp := &gameMsg.BaccaratStateOpenResp{
			Times: &protoMsg.TimeInfo{
				OutTime:   0,
				WaitTime:  YamlObj.Baccarat.Duration.Over,
				TotalTime: YamlObj.Baccarat.Duration.Over,
			},
		}
		self.Game.RegisterEvent(protoMsg.GameScene_Over, overResp, openResp.Times.TotalTime, func() bool {
			self.Over(nil)
			if self.IsClear || self.Close(func() *Table {
				return self.T
			}) { // 剩余次数为零,释放资源
				return false //表示不执行after
			}
			return true
		}, func() {
			// 校准剩余次数 有人玩,才减去
			self.personBetInfo.Range(func(key, value any) bool {
				self.T.CalibratingRemain(Default)
				return false
			})
			openResp.Times.TimeStamp = self.TimeStamp
		}, nil)
	}
	self.ChangeStateAndWork(protoMsg.GameScene_Start)
	return true
}

// Playing 下注 直接扣除金币
func (self *BaccaratGame) Playing(args []interface{}) bool {
	if !self.Game.Playing(args) {
		log.Debugf("[%v:%v][%v] Playing data fail! ", self.GameInfo.Name, self.T.Id, StatusText[User02])
		return false
	}
	_ = args[1]
	//【消息】
	m := args[0].(*gameMsg.BaccaratBetReq)
	//【传输对象】
	agent := args[1].(Agent)

	userData := agent.UserData()
	person := userData.(*Player)
	uid := person.UserID
	if m.BetScore <= 0 {
		log.Debugf("[%v:%v][%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, StatusText[Game07], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return false
	}

	sitter := self.T.GetChair(uid)
	if sitter == nil {
		log.Debugf("[%v:%v][%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, StatusText[Game33], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game33])
		return false
	}
	if self.bankerID == uid {
		log.Debugf("[%v:%v][%v:%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, uid, StatusText[Game27], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game27])
		return false
	}
	gold := sitter.PlayerInfo.Gold
	if gold < m.BetScore+sitter.Total {
		log.Debugf("[%v:%v][%v:%v] %v Playing:->%v ", self.GameInfo.Name, self.T.Id, uid, StatusText[Game05], gold, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
		return false
	}
	if AREA_MAX <= m.BetArea {
		log.Debugf("[%v:%v][%v:%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, uid, StatusText[Game06], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game06])
		return false
	}
	// log.Debugf("[%v:%v] 玩家:%v Playing:->%v ", self.GameInfo.Name, self.T.Id, sitter.UserID, m)

	//下注成功
	msg := &gameMsg.BaccaratBetResp{}
	msg.UserID = sitter.UserID
	msg.BetArea = m.BetArea
	msg.BetScore = m.BetScore
	sitter.Score = m.BetScore
	sitter.Total += m.BetScore
	if value, ok := self.personBetInfo.Load(msg.UserID); ok {
		ok = false
		areaBetInfos, ok1 := value.([]*gameMsg.BaccaratBetResp)
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				areaBetInfos[index].BetScore = betItem.BetScore + m.BetScore
				log.Debugf("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.GameInfo.Name, self.T.Id, uid, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = utils.CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*gameMsg.BaccaratBetResp)
		}
		self.personBetInfo.Store(uid, areaBetInfos)
	} else {
		log.Debugf("[%v:%v]\t玩家:%v 下注:%+v", self.GameInfo.Name, self.T.Id, uid, m)
		areaBetInfos := make([]*gameMsg.BaccaratBetResp, 0)
		areaBetInfos = utils.CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*gameMsg.BaccaratBetResp)
		self.personBetInfo.Store(uid, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying
	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	return true
}

// Over 结算
func (self *BaccaratGame) Over(args []interface{}) bool {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		betInfos, ok1 := value.([]*gameMsg.BaccaratBetResp)
		if !ok || !ok1 {
			return true
		}
		sitter := self.T.GetChair(uid)
		//每一次的下注信息
		for _, betInfo := range betInfos {
			log.Infof("[%v:%v]玩家:%v,下注区域:%v 下注金额:%v", self.GameInfo.Name, self.T.Id, uid, betInfo.BetArea, betInfo.BetScore)
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

	// 对平局的先结算
	result := fmt.Sprintf("闲:[%v] 庄:[%v] 中奖区域:%v", GetCardsText(self.cbPlayerCards.Cards), GetCardsText(self.cbBankerCards.Cards), self.openInfo.AwardArea)
	// 结算
	self.T.ChairSettle(self.Name, self.Inning, result)
	// 移除已离线玩家
	noLiveIds := self.T.CheckChairsNoLive()
	for _, uid := range noLiveIds {
		self.RemovePlayer(uid)
	}
	// 发送游戏结算数据
	//派奖
	checkout := &gameMsg.BaccaratCheckoutResp{}
	checkout.Acquires = allAreaInfo
	self.T.ChairWork(func(chair *Chair) {
		if chair.Total != INVALID {
			checkout.MyAcquire = chair.Gain
		}
		GetClientMgr().SendTo(chair.UserID, checkout)
	})

	log.Infof("[%v:%v]   \t结算注单... 各区域情况:%v", self.GameInfo.Name, self.T.Id, allAreaInfo)
	return true
}

// UpdateInfo 检测是否可以站起
func (self *BaccaratGame) UpdateInfo(args []interface{}) bool {
	flag, ok := args[0].(protoMsg.PlayerState)
	if !ok {
		return false
	}
	uid, ok1 := args[1].(int64)
	if !ok1 {
		return false
	}
	bRet := false
	switch flag {
	case protoMsg.PlayerState_PlayerStandUp, protoMsg.PlayerState_PlayerGiveUp:
		self.personBetInfo.Range(func(key, value any) bool {
			if uid1, ok2 := key.(int64); ok2 && uid == uid1 {
				bRet = true
				return false
			}
			return true
		})
		if bRet {
			return false
		}
		self.RemovePlayer(uid)
		return true
	}

	return self.Game.UpdateInfo(args)
}

// 发牌
func (self *BaccaratGame) DispatchCard() {
	//6张牌
	timeoutCount := 0
Dispatch:
	self.cbCardCount = []byte{2, 2}
	tableCards := RandCardList(CardCountPlayer + CardCountBanker)

	//首次发牌
	log.Debugf("[%v:%v]\t-->开始发牌啦<---  牌堆中取牌:%v ", self.GameInfo.Name, self.T.Id, GetCardsText(tableCards))
	copy(self.cbPlayerCards.Cards[:CardCountPlayer], tableCards[0:self.cbCardCount[INDEX_PLAYER]])
	copy(self.cbBankerCards.Cards[:CardCountBanker], tableCards[CardCountBanker:CardCountBanker+self.cbCardCount[INDEX_BANKER]])
	log.Debugf("[%v:%v]\t第一次发牌 \t闲家:%v \t庄家:%v", self.GameInfo.Name, self.T.Id, GetCardsText(self.cbPlayerCards.Cards), GetCardsText(self.cbBankerCards.Cards))
	//计算点数
	self.cbBankerCards.CardValue = int32(GetCardListPip(self.cbBankerCards.Cards))
	self.cbPlayerCards.CardValue = int32(GetCardListPip(self.cbPlayerCards.Cards))

	//闲家补牌
	var cbPlayerThirdCardValue byte = 0 //第三张牌点数
	if self.cbPlayerCards.CardValue <= 5 && self.cbBankerCards.CardValue < 8 {
		//计算点数
		self.cbCardCount[INDEX_PLAYER]++
		cbPlayerThirdCardValue = GetCardPip(tableCards[CardCountPlayer-1])
		self.cbPlayerCards.Cards[CardCountPlayer-1] = tableCards[CardCountPlayer-1]

		//补牌后,再计算一次点数
		self.cbBankerCards.CardValue = int32(GetCardListPip(self.cbBankerCards.Cards))
		self.cbPlayerCards.CardValue = int32(GetCardListPip(self.cbPlayerCards.Cards))
	}

	//庄家补牌
	if self.cbPlayerCards.CardValue < 8 && self.cbBankerCards.CardValue < 8 {
		switch self.cbBankerCards.CardValue {
		case 0:
			self.cbCardCount[INDEX_BANKER]++
		case 1:
			self.cbCardCount[INDEX_BANKER]++
		case 2:
			self.cbCardCount[INDEX_BANKER]++
		case 3:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 8) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 4:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 1 && cbPlayerThirdCardValue != 8 && cbPlayerThirdCardValue != 9 && cbPlayerThirdCardValue != 0) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 5:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 1 && cbPlayerThirdCardValue != 2 && cbPlayerThirdCardValue != 3 && cbPlayerThirdCardValue != 8 && cbPlayerThirdCardValue != 9 && cbPlayerThirdCardValue != 0) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 6:
			if self.cbCardCount[INDEX_PLAYER] == 3 && (cbPlayerThirdCardValue == 6 || cbPlayerThirdCardValue == 7) {
				self.cbCardCount[INDEX_BANKER]++
			}
		}
		self.cbBankerCards.CardValue = int32(GetCardListPip(self.cbBankerCards.Cards))
	}
	if self.cbCardCount[INDEX_BANKER] == 3 {
		self.cbBankerCards.Cards[CardCountBanker-1] = tableCards[CardCountBanker+CardCountPlayer-1]
	}

	// 开奖状态
	self.openInfo.PlayerCard = self.cbPlayerCards
	self.openInfo.BankerCard = self.cbBankerCards
	self.openInfo.AwardArea = self.DeduceWin()
	strOpen := ""
	for i, b := range self.openInfo.AwardArea {
		if b == 1 {
			strOpen += GetAreaName(i)
		}
	}

	//是否满足库存
	personAwardScore, ok := self.simulatedResult()
	if !ok && timeoutCount < 10000 {
		timeoutCount++
		//log.Errorf("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v 尝试次数:%v[库存警告]", self.GameInfo.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore, timeoutCount)
		goto Dispatch
	}
	log.Infof("[%v:%v]\t当前开奖区域:%v 闲家:[%v]  庄家:[%v]", self.GameInfo.Name, self.T.Id, strOpen, self.openInfo.PlayerCard.Cards, self.openInfo.BankerCard.Cards)

	log.Debugf("[%v:%v]\t当前库存:%v 结算后:%v 扣算前:%v", self.GameInfo.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore)
	self.inventory = personAwardScore
	self.bankerScore = personAwardScore - self.inventory
	if self.bankerScore < INVALID {
		self.bankerScore = INVALID
	}
	if self.inventory < INVALID {
		log.Errorf("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v [库存警告]", self.GameInfo.Name, self.T.Id, self.inventory, self.bankerScore)
	}
	//// ===>>>
	//self.bankerScore = INVALID
	//self.inventory = INVALID
	log.Debugf("[%v:%v]\t 库存信息 当前:%v 配置:%v 盈利:[%v]", self.GameInfo.Name, self.T.Id, self.inventory, self.confInventory, self.inventory-self.confInventory)

	//录单记录
	self.openAreas = append(self.openAreas, self.openInfo.AwardArea)
	lenOpen := len(self.openAreas)
	if MaxLoadNum < lenOpen {
		self.openAreas = self.openAreas[lenOpen-MaxLoadNum : lenOpen]
	}

	GlobalSender.NotifyOthers(self.PlayerList, self.openInfo)
	// 再次计算点数
	log.Debugf("[%v:%v]\t补牌结果\t闲家:%v  庄家:%v 开奖记录:%+v", self.GameInfo.Name, self.T.Id, GetCardsText(self.cbPlayerCards.Cards), GetCardsText(self.cbBankerCards.Cards), len(self.openAreas))
}

// 开奖区域
func (self *BaccaratGame) DeduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	//胜利区域--------------------------
	//平
	if self.cbPlayerCards.CardValue == self.cbBankerCards.CardValue {

		pWinArea[AREA_PING] = Win
		log.Debugf("[%v:%v]\t平", self.GameInfo.Name, self.T.Id)
		// 同平点
		if self.cbCardCount[INDEX_PLAYER] == self.cbCardCount[INDEX_BANKER] {
			var wCardIndex byte = 0
			for wCardIndex = 0; wCardIndex < self.cbCardCount[INDEX_PLAYER]; wCardIndex++ {
				cbBankerValue := GetCardValue(self.cbBankerCards.Cards[wCardIndex])
				cbPlayerValue := GetCardValue(self.cbPlayerCards.Cards[wCardIndex])
				if cbBankerValue != cbPlayerValue {
					break
				}
			}

			if wCardIndex == self.cbCardCount[INDEX_PLAYER] {
				pWinArea[AREA_TONG_DUI] = Win
				log.Debugf("[%v:%v]\t同点平", self.GameInfo.Name, self.T.Id)
			}
		}
	} else if self.cbPlayerCards.CardValue < self.cbBankerCards.CardValue { //庄
		pWinArea[AREA_ZHUANG] = Win
		log.Debugf("[%v:%v]\t庄", self.GameInfo.Name, self.T.Id)
		//天王判断
		if self.cbBankerCards.CardValue == 8 || self.cbBankerCards.CardValue == 9 {
			pWinArea[AREA_ZHUANG_TIAN] = Win
			log.Debugf("[%v:%v]\t庄天王", self.GameInfo.Name, self.T.Id)
		}
	} else { //闲
		pWinArea[AREA_XIAN] = Win
		log.Debugf("[%v:%v]\t闲", self.GameInfo.Name, self.T.Id)
		//天王判断
		if self.cbPlayerCards.CardValue == 8 || self.cbPlayerCards.CardValue == 9 {
			pWinArea[AREA_XIAN_TIAN] = Win
			log.Debugf("[%v:%v]\t闲天王", self.GameInfo.Name, self.T.Id)
		}
	}

	//对子判断(前两张牌比较)
	if GetCardValue(self.cbPlayerCards.Cards[0]) == GetCardValue(self.cbPlayerCards.Cards[1]) ||
		(0 != self.cbPlayerCards.Cards[2] && (GetCardValue(self.cbPlayerCards.Cards[0]) == GetCardValue(self.cbPlayerCards.Cards[2]) ||
			GetCardValue(self.cbPlayerCards.Cards[1]) == GetCardValue(self.cbPlayerCards.Cards[2]))) {
		pWinArea[AREA_XIAN_DUI] = Win
		log.Debugf("[%v:%v]\t闲对子", self.GameInfo.Name, self.T.Id)
	}

	if GetCardValue(self.cbBankerCards.Cards[0]) == GetCardValue(self.cbBankerCards.Cards[1]) ||
		(0 != self.cbBankerCards.Cards[2] && (GetCardValue(self.cbBankerCards.Cards[0]) == GetCardValue(self.cbBankerCards.Cards[2]) ||
			GetCardValue(self.cbBankerCards.Cards[1]) == GetCardValue(self.cbBankerCards.Cards[2]))) {
		pWinArea[AREA_ZHUANG_DUI] = Win
		log.Debugf("[%v:%v]\t庄对子", self.GameInfo.Name, self.T.Id)
	}
	return pWinArea
}

// BonusArea 区域赔额
func (self *BaccaratGame) BonusArea(area int32, betScore int64) int64 {
	if self.openInfo.AwardArea[AREA_PING] == Win &&
		(area == AREA_XIAN || area == AREA_ZHUANG) {
		return 0 //返还本金
	}
	if Win == self.openInfo.AwardArea[area] {
		bonus := int64(Odds[int(area)]*100) * betScore / 100
		return bonus
	}
	return -betScore
}

// 抢庄
func (self *BaccaratGame) host(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	host := args[2].(*gameMsg.BaccaratHostReq)
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

		if BaccaratHostMax < size {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
		self.hostList = append(self.hostList, userID)
		self.hostList = utils.CopyInsert(self.hostList, size, userID).([]int64)

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

	log.Debugf("[%v:%v]\t有人来抢庄啦:%d 列表人数%d", self.GameInfo.Name, self.T.Id, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.BaccaratHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 超级抢庄
func (self *BaccaratGame) superHost(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)
	host := args[2].(*gameMsg.BaccaratSuperHostReq)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}
	user := userData.(*Player)
	userID := user.UserID
	log.Debugf("[%v:%v]\t有人要超级抢庄--->:%d", self.GameInfo.Name, self.T.Id, userID)

	if host.IsWant {
		if self.superHostID == 0 {
			self.superHostID = userID
			//超级抢庄放申请列表首位
			self.hostList = utils.CopyInsert(self.hostList, 0, userID).([]int64)
		} else {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
	}

	// 通知
	GlobalSender.NotifyOthers(self.PlayerList, &gameMsg.BaccaratSuperHostResp{
		UserID: self.bankerID,
		IsWant: host.IsWant,
	})

}

// 定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *BaccaratGame) permitHost() *gameMsg.BaccaratHostResp {
	//校验是否满足庄家条件 [5000 < 金额] 不可连续坐庄15次
	tempList := self.hostList
	log.Debugf("[%v:%v]\t定庄.... 列表数据:%v", self.GameInfo.Name, self.T.Id, self.hostList)

	//befBankerID := self.bankerID 避免重复
	for index, pid := range tempList {
		person := GlobalPlayerMgr.Get(pid)
		if person == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == person.UserID {
			log.Debugf("[%v:%v]\t不再连续坐庄：%v", self.GameInfo.Name, self.T.Id, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if person.Gold < 5000 {
			log.Debugf("[%v:%v]\t玩家%d 金币%lf 少于5000不能申请坐庄", self.GameInfo.Name, self.T.Id, person.UserID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Debugf("[%v:%v]\t连续坐庄次数:%d", self.GameInfo.Name, self.T.Id, self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]
		if banker := GlobalPlayerMgr.Get(self.bankerID); banker != nil {
			self.bankerScore = banker.Gold
		}

		log.Debugf("[%v:%v]\t确定庄家:%d", self.GameInfo.Name, self.T.Id, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		self.bankerScore = self.inventory
		log.Debugf("[%v:%v]\t系统坐庄", self.GameInfo.Name, self.T.Id)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &gameMsg.BaccaratHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Debugf("[%v:%v]\t广播上庄", self.GameInfo.Name, self.T.Id)
	return msg
}

// -----------------------逻辑层---------------------------

// reset 重新初始化[返回结果提供给外部清场用]
func (self *BaccaratGame) reset() {
	self.cbCardCount = []byte{2, 2}             // 首次拿牌张数
	self.openInfo = &gameMsg.BaccaratOpenResp{} // 结算结果

	// 闲家手牌
	self.cbPlayerCards = &protoMsg.CardInfo{
		Cards: make([]byte, CardCountPlayer),
	}
	// 庄家手牌
	self.cbBankerCards = &protoMsg.CardInfo{
		Cards: make([]byte, CardCountBanker),
	}
	//玩家下注信息
	self.personBetInfo.Range(func(key, value any) bool {
		self.personBetInfo.Delete(key)
		return true
	})
	self.Game.Init()
	return
}

func (self *BaccaratGame) simulatedResult() (int64, bool) {
	//var personAwardScore int64 = self.bankerScore + self.inventory
	if self.T.Rid != SYSTEMID {
		return self.bankerScore, true
	}
	var personAwardScore int64 = self.inventory
	self.personBetInfo.Range(func(key, value any) bool {
		_, ok := key.(int64)
		if !ok {
			return true
		}
		betInfos, ok1 := value.([]*gameMsg.BaccaratBetResp)
		if !ok1 {
			return true
		}
		//if GetPlayerMgr().CheckRobot(uid) {
		//	return true
		//}
		for _, betInfo := range betInfos {
			if self.openInfo.AwardArea[AREA_PING] == Win &&
				(betInfo.BetArea == AREA_XIAN || betInfo.BetArea == AREA_ZHUANG) {
				continue
			}
			if Win == self.openInfo.AwardArea[betInfo.BetArea] { //玩家奖金
				// 需扣除本金
				personAwardScore -= self.BonusArea(betInfo.BetArea, betInfo.BetScore) + betInfo.BetScore
			}
		}
		return true
	})

	log.Debugf("[%v:%v]\t模拟结算...前:%v 后:%v   当前库存:%v", self.GameInfo.Name, self.T.Id, self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
