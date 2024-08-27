package brtoubao

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/internal/utils"
	. "superman/nodes/game/manger"
	. "superman/nodes/game/module/category"
	"sync"
	"time"
)

//var lock *sync.Mutex = &sync.Mutex{} //锁

// BrToubaoGame 百人类
type BrToubaoGame struct {
	*Game
	T             *Table
	personBetInfo *sync.Map                  //map[int64][]*protoMsg.BrtoubaoBetResp //下注信息
	openAreas     [][]byte                   //开奖纪录
	openInfo      *protoMsg.BrtoubaoOpenResp //开奖结果

	hostList    []int64 //申请列表(默认11个)
	bankerID    int64   //庄家ID
	superHostID int64   //超级抢庄ID
	bankerScore int64   //庄家积分
	keepTwice   uint8   //连续坐庄次数

	inventory     int64 //库存
	confInventory int64
}

// New 创建百人骰宝实例
func New(game *Game, table *Table) *BrToubaoGame {
	p := &BrToubaoGame{
		Game: game,
		T:    table,
	}

	p.bankerScore = 0
	p.Init()
	return p
}

// ---------------------------百人骰宝----------------------------------------------//
// 初始化信息
func (self *BrToubaoGame) Init() {
	self.bankerID = 0         // 庄家ID
	self.superHostID = 0      // 超级抢庄ID
	self.keepTwice = 0        // 连续抢庄次数
	self.hostList = []int64{} // 申请列表(默认11个)

	self.openInfo = &protoMsg.BrtoubaoOpenResp{} // 结算结果

	self.personBetInfo = new(sync.Map)      //make(map[int64][]*protoMsg.BrtoubaoBetResp) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 20) // 开奖纪录
}

//获取玩家列表
//func (self *BrToubaoGame) getPlayersName() []string {
//	var name []string
//	for _, pid := range self.PlayerList {
//		log.Infof("----玩家：%v", pid)
//		name = append(name, sqlhandle.CheckName(pid))
//	}
//	return name
//}

// //////////////////////////////////////////////////////////////
// 场景信息
func (self *BrToubaoGame) Scene(args []interface{}) bool {
	if !self.Game.Scene(args) {
		return false
	}
	person := args[0].(*Player)

	//场景信息
	StateInfo := &protoMsg.BrtoubaoSceneResp{
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
		bets, ok1 := value.([]*protoMsg.BrtoubaoBetResp)
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
	case protoMsg.GameScene_Start: //准备
		timeInfo.TotalTime = YamlObj.Brtoubao.Duration.Start
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrtoubaoStateStartResp{
			Times:  timeInfo,
			Inning: self.Inning,
		})
	case protoMsg.GameScene_Playing: //下注
		timeInfo.TotalTime = YamlObj.Brtoubao.Duration.Play
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrtoubaoStatePlayingResp{
			Times: timeInfo,
		})
	case protoMsg.GameScene_Opening: //开奖
		timeInfo.TotalTime = YamlObj.Brtoubao.Duration.Open
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrtoubaoStateOpenResp{
			Times:    timeInfo,
			OpenInfo: self.openInfo,
		})
	case protoMsg.GameScene_Over: //结算
		timeInfo.TotalTime = YamlObj.Brtoubao.Duration.Over
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendTo(uid, &protoMsg.BrtoubaoStateOverResp{
			Times: timeInfo,
		})
	}
	return true
}

// 开始下注前定庄
func (self *BrToubaoGame) Start(args []interface{}) bool {
	_ = args
	log.Infof("[%v:%v]游戏創建成功", self.GameInfo.Name, self.T.Id)
	if self.IsStart {
		self.onStart()
		self.IsStart = false
	}
	return true
}

// Playing 下注 直接扣除金币
func (self *BrToubaoGame) Playing(args []interface{}) bool {
	if !self.Game.Playing(args) {
		return false
	}
	//【传输对象】
	agent := args[1].(Agent)
	m := args[0].(*protoMsg.BrtoubaoBetReq)
	if m.BetScore <= 0 {
		log.Infof("[%v:%v][%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, StatusText[Game07], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return false
	}

	userData := agent.UserData()
	person := userData.(*Player)
	sitter := self.T.GetChair(person.UserID)
	if sitter == nil {
		log.Infof("[%v:%v][%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, StatusText[Game33], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game33])
		return false
	}
	if self.bankerID == sitter.UserID {
		log.Infof("[%v:%v][%v:%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, sitter.UserID, StatusText[Game27], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game27])
		return false
	}
	if sitter.Gold < m.BetScore+sitter.Total {
		log.Infof("[%v:%v][%v:%v] %v Playing:->%v ", self.GameInfo.Name, self.T.Id, sitter.UserID, StatusText[Game05], sitter.Gold, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
		return false
	}
	if AREA_MAX <= m.BetArea {
		log.Infof("[%v:%v][%v:%v] Playing:->%v ", self.GameInfo.Name, self.T.Id, sitter.UserID, StatusText[Game06], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game06])
		return false
	}
	//log.Infof("[%v:%v] 玩家:%v Playing:->%v ", self.GameInfo.Name, self.T.Id, sitter.UserID, m)

	//数据库中冻结玩家金币[下注成功]
	//if money, fact, ok := GlobalSqlHandle.DeductMoney(person.UserID, m.BetScore, CodeRefund, SYSTEMID, self.Inning); ok {
	//	//log.Infof("下注成功 玩家ID:%v 现有金币:%v 下注金币:%v", person.UserID, money, m.BetScore)
	//	GlobalSender.SendTo(uid, &protoMsg.GoldChangeInfo{
	//		UserID:    person.UserID,
	//		Gold:      money,
	//		AlterGold: -fact,
	//		Code:      CodeRefund,
	//		Reason:    self.Inning,
	//	})
	//} else {
	//	log.Infof("下注失败 玩家ID:%v 现有金币:%v 下注金币:%v", person.UserID, money, m.BetScore)
	//	GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
	//	return
	//}

	//下注成功
	//下注数目累加
	msg := &protoMsg.BrtoubaoBetResp{}
	msg.UserID = sitter.UserID
	msg.BetArea = m.BetArea
	msg.BetScore = m.BetScore

	sitter.Score = m.BetScore
	sitter.Total += m.BetScore
	if value, ok := self.personBetInfo.Load(sitter.UserID); ok {
		ok = false
		areaBetInfos, ok1 := value.([]*protoMsg.BrtoubaoBetResp)
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				areaBetInfos[index].BetScore = betItem.BetScore + m.BetScore
				log.Infof("[%v:%v]\t玩家:%v 区域:%v 累加:->%v", self.GameInfo.Name, self.T.Id, sitter.UserID, m.BetArea, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok || !ok1 {
			areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrtoubaoBetResp)
		}
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	} else {
		log.Infof("[%v:%v]\t玩家:%v 下注:%v", self.GameInfo.Name, self.T.Id, sitter.UserID, m)
		areaBetInfos := make([]*protoMsg.BrtoubaoBetResp, 0)
		areaBetInfos = CopyInsert(areaBetInfos, len(areaBetInfos), msg).([]*protoMsg.BrtoubaoBetResp)
		self.personBetInfo.Store(sitter.UserID, areaBetInfos)
	}
	person.State = protoMsg.PlayerState_PlayerPlaying

	self.bankerScore += msg.BetScore
	//通知其他玩家
	GlobalSender.NotifyOthers(self.PlayerList, msg)
	return true
}

// 结算
func (self *BrToubaoGame) Over(args []interface{}) bool {
	_ = args
	allAreaInfo := make([]int64, AREA_MAX)
	self.personBetInfo.Range(func(key, value any) bool {
		uid, ok := key.(int64)
		betInfos, ok1 := value.([]*protoMsg.BrtoubaoBetResp)
		if !ok || !ok1 {
			return true
		}
		sitter := self.T.GetChair(uid)
		if sitter == nil {
			return true
		}

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

	result := fmt.Sprintf("%v 中奖区域:%v", self.openInfo.Dice, self.openInfo.AwardArea)
	self.T.ChairSettle(self.Name, self.Inning, result)
	// 移除已离线玩家
	noLiveIds := self.T.CheckChairsNoLive()
	for _, uid := range noLiveIds {
		self.RemovePlayer(uid)
	}
	//派奖
	checkout := &protoMsg.BrtoubaoCheckoutResp{}
	checkout.Acquires = allAreaInfo
	// 统一结算
	self.T.ChairWork(func(chair *Chair) {
		if chair.Total != INVALID {
			checkout.MyAcquire = chair.Gain
		}
		GetClientMgr().SendTo(chair.UserID, checkout)
	})

	log.Infof("[%v:%v]   \t结算注单... 各区域情况:%v", self.GameInfo.Name, self.T.Id, allAreaInfo)
	return true
}

// 发牌
func (self *BrToubaoGame) DispatchCard() {
	//6张牌
	timeoutCount := 0
Dispatch:
	self.openInfo.Dice = RollDice(CardCount)
	// 中奖区域
	self.openInfo.AwardArea = self.DeduceWin()

	//是否满足库存
	personAwardScore, ok := self.simulatedResult()
	if !ok && timeoutCount < 1000 {
		timeoutCount++
		//log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v 尝试次数:%v[库存警告]", self.GameInfo.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore, timeoutCount)
		goto Dispatch
	}
	log.Infof("[%v:%v]\t当前库存:%v 庄家积分:%v 扣算之前:%v", self.GameInfo.Name, self.T.Id, self.inventory, personAwardScore, self.bankerScore)
	self.inventory = personAwardScore
	self.bankerScore = personAwardScore - self.inventory
	if self.bankerScore < INVALID {
		self.bankerScore = INVALID
	}
	if self.inventory < INVALID {
		log.Error("[库存警告][%v:%v]\t当前库存:%v 庄家积分:%v [库存警告]", self.GameInfo.Name, self.T.Id, self.inventory, self.bankerScore)
	}
	// ===>>>
	//self.bankerScore = INVALID
	//self.inventory = INVALID
	log.Infof("[%v:%v]\t 库存信息 当前:%v 配置:%v 盈利:[%v]", self.GameInfo.Name, self.T.Id, self.inventory, self.confInventory, self.inventory-self.confInventory)

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
func (self *BrToubaoGame) DeduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	//胜利区域--------------------------
	sum := byte(INVALID)
	for _, num := range self.openInfo.Dice {
		sum += num
	}

	if Small < sum && sum < Largest {
		//大小
		if Middle < sum {
			pWinArea[AREA_Big] = Win
		} else {
			pWinArea[AREA_Small] = Win
		}

		//单双
		if 0 == sum%2 {
			pWinArea[AREA_Double] = Win
		} else {
			pWinArea[AREA_Single] = Win
		}

		//数字区域
		pWinArea[AREA_4+sum-4] = Win
	}

	//baozi
	if self.openInfo.Dice[0] == self.openInfo.Dice[1] && self.openInfo.Dice[1] == self.openInfo.Dice[2] {
		pWinArea[AREA_Baozi] = Win
		pWinArea[AREA_Baozi+4+self.openInfo.Dice[0]] = Win
	}

	return pWinArea
}

// 区域赔额
func (self *BrToubaoGame) BonusArea(area int32, betScore int64) int64 {
	if Win == self.openInfo.AwardArea[area] {
		bonus := int64(Odds[int(area)]*100) * betScore / 100
		return bonus
	}
	return -betScore
}

// 开始|下注|结算|空闲
func (self *BrToubaoGame) onStart() {
	self.reset()
	// 校准库存
	if self.confInventory != YamlObj.Brtoubao.Inventory {
		self.confInventory = YamlObj.Brtoubao.Inventory
		//此处会重置以往的库存值
		self.inventory = self.confInventory
	}

	self.ChangeState(protoMsg.GameScene_Start)
	time.AfterFunc(time.Duration(YamlObj.Brtoubao.Duration.Start)*time.Second, self.onPlay)

	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoStateStartResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  YamlObj.Brtoubao.Duration.Start,
			TotalTime: YamlObj.Brtoubao.Duration.Start,
		},
		Inning: self.Inning,
	})
}

// [下注减法]
func (self *BrToubaoGame) onPlay() {
	self.ChangeState(protoMsg.GameScene_Playing)
	time.AfterFunc(time.Duration(YamlObj.Brtoubao.Duration.Play)*time.Second, self.onOpen)

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoStatePlayingResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  YamlObj.Brtoubao.Duration.Play,
			TotalTime: YamlObj.Brtoubao.Duration.Play,
		},
	})
}

func (self *BrToubaoGame) onOpen() {
	self.ChangeState(protoMsg.GameScene_Opening)
	// 开始状态
	//m := self.permitHost() //反馈定庄信息
	//GlobalSender.NotifyOthers(self.PlayerList, MainGameState, protoMsg.GameScene_Start, m)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  YamlObj.Brtoubao.Duration.Open,
			TotalTime: YamlObj.Brtoubao.Duration.Open,
		},
		OpenInfo: self.openInfo,
	})

	//【发牌】
	self.DispatchCard()

	time.AfterFunc(time.Duration(YamlObj.Brtoubao.Duration.Open)*time.Second, self.onOver)

}

// [结算加法]
func (self *BrToubaoGame) onOver() {
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
	time.AfterFunc(time.Duration(YamlObj.Brtoubao.Duration.Over)*time.Second, self.onStart)
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoStateOverResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   INVALID,
			WaitTime:  YamlObj.Brtoubao.Duration.Over,
			TotalTime: YamlObj.Brtoubao.Duration.Over,
		},
	})
}

// 抢庄
func (self *BrToubaoGame) host(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	host := args[2].(*protoMsg.BrtoubaoHostReq)
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

		if BrtoubaoHostMax < size {
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

	log.Infof("[%v:%v]有人来抢庄啦:%d 列表人数%d", self.GameInfo.Name, self.T.Id, userID, len(self.hostList))
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoHostResp{
		UserID: userID,
		IsWant: host.IsWant,
	})
}

// 超级抢庄
func (self *BrToubaoGame) superHost(args []interface{}) {
	//【消息】
	_ = args[2]
	agent := args[1].(Agent)
	host := args[2].(*protoMsg.BrtoubaoSuperHostReq)

	userData := agent.UserData()
	if nil == userData {
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}
	user := userData.(*Player)
	userID := user.UserID
	log.Infof("[%v:%v]有人要超级抢庄--->:%d", self.GameInfo.Name, self.T.Id, userID)

	if host.IsWant {
		if self.superHostID == INVALID {
			self.superHostID = userID
			//超级抢庄放申请列表首位
			self.hostList = CopyInsert(self.hostList, 0, userID).([]int64)
		} else {
			GlobalSender.SendResult(agent, FAILED, StatusText[Game11])
			return
		}
	}

	// 通知
	GlobalSender.NotifyOthers(self.PlayerList, &protoMsg.BrtoubaoSuperHostResp{
		UserID: self.bankerID,
		IsWant: host.IsWant,
	})

}

// 定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *BrToubaoGame) permitHost() *protoMsg.BrtoubaoHostResp {
	//校验是否满足庄家条件 [5000 < 金额] 不可连续坐庄15次
	tempList := self.hostList
	log.Infof("[%v:%v]定庄.... 列表数据:%v", self.GameInfo.Name, self.T.Id, self.hostList)

	//befBankerID := self.bankerID 避免重复
	for index, pid := range tempList {
		person := GlobalPlayerMgr.Get(pid)
		if person == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == person.UserID {
			log.Infof("[%v:%v]不再连续坐庄：%v", self.GameInfo.Name, self.T.Id, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if person.Gold < 5000 {
			log.Infof("[%v:%v]玩家%d 金币%lf 少于5000不能申请坐庄", self.GameInfo.Name, self.T.Id, person.UserID, person.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Infof("[%v:%v]连续坐庄次数:%d", self.GameInfo.Name, self.T.Id, self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]

		if banker := GlobalPlayerMgr.Get(self.bankerID); banker != nil {
			self.bankerScore = banker.Gold
		}

		log.Infof("[%v:%v]确定庄家:%d", self.GameInfo.Name, self.T.Id, self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		//self.bankerScore = 1000000
		log.Infof("[%v:%v]系统坐庄", self.GameInfo.Name, self.T.Id)
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &protoMsg.BrtoubaoHostResp{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Infof("[%v:%v]广播上庄", self.GameInfo.Name, self.T.Id)
	return msg
}

// -----------------------逻辑层---------------------------
// 重新初始化[返回结果提供给外部清场用]
func (self *BrToubaoGame) reset() {
	self.personBetInfo.Range(func(key, value any) bool {
		self.personBetInfo.Delete(key)
		return true
	})
	self.Game.Init()
	return
}

func (self *BrToubaoGame) simulatedResult() (int64, bool) {
	//var personAwardScore int64 = self.bankerScore + self.inventory
	if self.T.Rid != SYSTEMID {
		return self.bankerScore, true
	}
	var personAwardScore int64 = self.inventory
	self.personBetInfo.Range(func(key, value any) bool {
		betInfos, ok1 := value.([]*protoMsg.BrtoubaoBetResp)
		if !ok1 {
			return true
		}
		for _, betInfo := range betInfos {
			//玩家奖金
			if Win == self.openInfo.AwardArea[betInfo.BetArea] {
				personAwardScore -= int64(Odds[int(betInfo.BetArea)])*betInfo.BetScore + betInfo.BetScore
			}
		}
		return true
	})

	log.Infof("[%v:%v]模拟结算...前:%v 后:%v   当前库存:%v", self.GameInfo.Name, self.T.Id, self.bankerScore, personAwardScore, self.inventory)
	return personAwardScore, 0 <= personAwardScore
}
