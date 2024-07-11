// Package manger
/*
# -*- coding: utf-8 -*-
# @Time : 2020/4/28 9:07
# @Author : Pitter
# @File : games.go
# @Software: GoLand
*/
package manger

import (
	"fmt"
	"github.com/po2656233/goleaf/gate"
	"os"
	"sort"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/base/syncmap"
	"superman/nodes/leaf/jettengame/conf"
	"superman/nodes/leaf/jettengame/sql/redis"
	"sync"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/po2656233/goleaf/log"
	"superman/nodes/leaf/jettengame/gamedata/goclib/util"
)

// ----------游戏----------------------

// 配置文件最后一次修改的时间戳
var settingFixTime int64 = 0

// ProductCallback 实例回调(根据桌子号创建游戏)
type ProductCallback func(game *Game) IGameOperate

type SQLCalculate func(info CalculateInfo) (nowMoney, factDeduct int64, isOK bool)

// ClearCallback 清场回调
type ClearCallback func() (gameID int64)

// Chair 椅子
type Chair struct {
	ChairID int                  //椅子号(仅表示序号) 0:空
	UserID  int64                //玩家
	State   protoMsg.PlayerState //玩家状态
	Code    byte                 //牌值状态码(扎金花用作是否看牌|听牌|碰牌|杠牌 )

	HandCard []byte //手牌
	Multiple int    //倍数|番数

	Score int64 //当前注额(斗地主叫分阶段为叫分分值)
	Total int64 //总共下注
	Gain  int64 //获利
	Gold  int64 //金币

	Challenge []int64     //挑战过的玩家
	Timer     *time.Timer //座位上的定时器
}

// Table 桌子
type Table struct {
	protoMsg.TableInfo
	Instance IGameOperate //游戏实例
}

// Room 房间
type Room struct {
	//ID      int64
	Num     int64
	Kind    int32 //玩家获取的当前游戏类型 [主要用于记录,并反馈最新列表信息]
	Level   int32 //玩家获取的当前游戏级别 [主要用于记录,并反馈最新列表信息]
	PageNum int32
}

// Game 游戏
type Game struct {
	G          *protoMsg.GameInfo
	T          *protoMsg.TableInfo
	ID         int64
	Config     *conf.Yaml
	IsStart    bool        // 第一次启动
	IsClear    bool        // 是否清场
	ReadyCount int32       // 已准备人数
	RunCount   int32       // 运行次数
	TimeStamp  int64       // 当前时间戳
	PlayerList []int64     // 玩家列表
	Sitter     syncmap.Map //map[int64]*Chair // 座位信息
	InningInfo *Inning     // 牌局信息
	//HaveRobot  bool              // 是否已经有机器人
	lock sync.RWMutex
}

// Inning 每局信息
type Inning struct {
	Number    string //牌局号(游戏名+桌号+时间戳)
	OpenAreas []byte //开奖纪录(路单)
}

// IGameOperate 子游戏接口
type IGameOperate interface {
	Scene(args []interface{})        //场 景
	Start(args []interface{})        //开 始
	Playing(args []interface{})      //游 戏(下分|下注)
	Over(args []interface{})         //结 算
	UpdateInfo(args []interface{})   //更新信息
	SuperControl(args []interface{}) //超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)
}

// IMultiPlayer 百人类
type IMultiPlayer interface {
	DispatchCard()                              //发牌
	DeduceWin() []byte                          //开奖区域
	BonusArea(area int32, betScore int64) int64 //区域赔额
}

// IAgainst 对战类
type IAgainst interface {
	Ready(arg []interface{})      //准备
	DispatchCard()                //发牌
	CallScore(args []interface{}) //叫分
	OutCard(args []interface{})   //出牌
	Discard(args []interface{})   //操作
}

//----------游戏管理----------------------
////游戏索引
//type GameIndex struct {
//    KindID int32
//    Level  int32
//    HostID int64 //房主ID  0:系统
//}

// GamesManger 管理房间 GamesIndex *Game
type GamesManger struct {
	sync.Map
	Tables []*Table
}

// IGameManagement 子游戏接口
type IGameManagement interface {
	CreateGame(gid int64, gInfo *protoMsg.GameInfo, tInfo *protoMsg.TableInfo) (*Game, bool) //创建游戏
	AddGame(game *Game) (*Game, bool)                                                        //添加游戏

	Match(game *Game, person *Player, product ProductCallback) (*Table, int)      //配桌
	ChangeTable(gid int64, person *Player, product ProductCallback) (*Table, int) //换桌

	Start(gid int64) (*Game, bool)    //开启(游戏)
	Maintain(gid int64) (*Game, bool) //维护
	Clear(gid int64) (*Game, bool)    //清场
	Close(gid int64) (*Game, bool)    //关闭

}

// //////////////////平台管理////////////////////////////////////
var gamesManger *GamesManger = nil
var gamesOnce sync.Once

// GetGamesManger 平台管理对象(单例模式)//manger.persons = make(map[int32]*Player)
func GetGamesManger() *GamesManger {
	gamesOnce.Do(func() {
		gamesManger = &GamesManger{
			Map: sync.Map{},
		}
	})
	return gamesManger
}

// ToMsg --------------椅子----------------------
func (self *Chair) ToMsg() *protoMsg.PlayerInfo {
	player := GetPlayerManger().Get(self.UserID)
	person := *player
	person.ChairID = int32(self.ChairID)
	person.UserID = self.UserID
	person.Gold = self.Gold
	person.State = self.State
	if person.Sex == 0x0F && person.State == protoMsg.PlayerState_PlayerPickUp {
		person.State = protoMsg.PlayerState_PlayerPlaying
	}
	ptrPerson := &person
	return ptrPerson.ToMsg()
}

// -----------------桌子-----------------

// AddPlayer 加入玩家列表
func (self *Game) AddPlayer(userID int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.PlayerList = CopyInsert(self.PlayerList, len(self.PlayerList), userID).([]int64)
	self.PlayerList = SliceRemoveDuplicate(self.PlayerList).([]int64)
	self.T.SitterCount = int32(len(self.PlayerList))
	redis.RedisHandle().Set(GetGameKey(self.ID), self.T.SitterCount, 0)
	log.Release("[%v:%v]\t[新增]玩家ID:%v 当前玩家数:%v ", self.G.Name, self.ID, userID, self.T.SitterCount)
}

// DeletePlayer 从玩家列表中踢出
func (self *Game) DeletePlayer(userID int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.PlayerList = SliceRemoveDuplicate(self.PlayerList).([]int64)
	self.PlayerList = DeleteValue(self.PlayerList, userID).([]int64)
	self.T.SitterCount = int32(len(self.PlayerList))
	//if self.T.SitterCount < INVALID {
	//	self.T.SitterCount = INVALID
	//}
	redis.RedisHandle().Set(GetGameKey(self.ID), self.T.SitterCount, 0)
	log.Release("[%v:%v]\t[剔除]玩家ID:%v 当前玩家数:%v ", self.G.Name, self.ID, userID, self.T.SitterCount)

}

// AddChair 添加座位
func (self *Game) AddChair(person *Player) bool {
	//是否之前已经占座了
	have := false
	isOk := false

	// 读取座位信息
	self.Sitter.Range(func(kUid, vChairObj interface{}) bool {
		sit, ok := vChairObj.(*Chair)
		if !ok {
			return true
		}
		// 已经占座
		if person.ChairID != INVALID && sit.ChairID == int(person.ChairID) {
			have = true
		}
		// 同步用户id上的座位信息
		if sit.UserID == person.UserID {
			person.ChairID = int32(sit.ChairID)
			person.Gold = sit.Gold
			isOk = true
			return false
		}
		return true
	})

	if isOk {
		return true
	}
	size := self.Sitter.Length() + 1
	if self.G.MaxCount != INVALID && self.G.MaxCount < int32(size) {
		return false
	}
	//添加桌子号
	if !have {
		// 坐在
		for i := IndexStart; i <= int(size); i++ {
			have = false
			self.Sitter.Range(func(kUid, vChairObj interface{}) bool {
				sit, ok := vChairObj.(*Chair)
				if !ok {
					return true
				}
				if sit.ChairID == i {
					have = true
					return false
				}
				return true
			})
			if !have {
				person.ChairID = int32(i)
				break
			}
		}

	}

	chair := &Chair{
		ChairID:   int(person.ChairID),
		UserID:    person.UserID,
		HandCard:  make([]byte, 0),
		Challenge: make([]int64, 0),
		Score:     INVALID,
		State:     protoMsg.PlayerState_PlayerLookOn,
		Multiple:  INVALID,
		Code:      SeatNoCode,
		Gain:      INVALID,
		Gold:      self.T.PlayScore,
		Timer:     nil,
	}

	if self.T.HostID == SYSTEMID {
		chair.Gold = person.Money
		person.Gold = person.Money
	} else if chair.Gold == INVALID {
		chair.Gold = person.Gold
	}
	// 不是房卡类型\电玩城 且 在空闲场景下
	if self.G.Scene == protoMsg.GameScene_Free &&
		self.G.Type != protoMsg.GameType_TableCard &&
		self.G.Type != protoMsg.GameType_GamesCity {
		chair.Timer = time.AfterFunc(time.Duration(MaxWait)*time.Second, func() {
			if chair.State == protoMsg.PlayerState_PlayerLookOn && self.G.Scene == protoMsg.GameScene_Free {
				self.UpdatePlayerList(DEL, chair.UserID, true)
				GetClientManger().SendTo(chair.UserID, &protoMsg.ResultPopResp{
					Flag:  SUCCESS,
					Title: StatusText[Title001],
					Hints: StatusText[Game09],
				})
			}
		})
	}
	person.Gold = chair.Gold
	// 添加座位信息
	self.Sitter.Store(person.UserID, chair)
	return true
}

func (self *Game) NextChair(curChair *Chair) (*Chair, bool) {
	if curChair == nil {
		return nil, false
	}
	// 获取当前玩家的坐席
	chair, ok := self.GetSitter(curChair.UserID)
	if !ok {
		return nil, false
	}
	chairID := chair.ChairID
	keys := self.GetAllSitter()
	size := len(keys)

	// 与当前玩家座次比较，找出下一位玩家
	sort.Slice(keys, func(i, j int) bool {
		//绝对值
		iChair := keys[i].ChairID - chairID
		if iChair < 0 {
			iChair = keys[i].ChairID + size
		}
		jChair := keys[j].ChairID - chairID
		if jChair < 0 {
			jChair = keys[j].ChairID + size
		}
		return iChair < jChair
	})

	for _, sitter := range keys {
		if sitter.State < protoMsg.PlayerState_PlayerAgree || sitter.ChairID == chairID {
			continue
		}
		return sitter, true
	}
	return nil, false
}

func (self *Game) NextChairX(uid int64) (*Chair, bool) {
	chair, ok := self.Sitter.Load(uid)
	return chair.(*Chair), ok
}

func (self *Game) DeleteChair(uid int64) {
	sit, ok := self.GetSitter(uid)
	if ok && sit.Timer != nil {
		sit.Timer.Stop()
	}
	self.Sitter.Delete(uid)
}

// Loading 加载配置 [废弃:不加载公用配置,拆分至子游戏内部实现]
func (self *Game) Loading() {

	//获取文件的
	settingFile := conf.GameYamlPath
	fInfo, _ := os.Stat(settingFile)
	lastFixTime := fInfo.ModTime().Unix()
	// 在文件被修改或配置内容为空时，则读取配置信息
	if lastFixTime != settingFixTime || self.Config == nil {
		//读取配置信息
		yamlFile, err := os.ReadFile(settingFile)
		//log.Debug("yamlFile:%v", yamlFile)
		if err != nil {
			log.Fatal("yamlFile.Get err #%v ", err)
		}

		err = yaml.Unmarshal(yamlFile, self.Config)
		// err = yaml.Unmarshal(yamlFile, &resultMap)
		if err != nil {
			log.Error("Unmarshal: %v", err)
		}
	}
	settingFixTime = lastFixTime

	//log.Debug("conf:%v", self.Config)
}

// Reset 重置座位上的玩家
func (self *Game) Reset() int64 {

	self.ReadyCount = INVALID
	self.RunCount++
	log.Debug("[%v:%v]   \t扫地僧出来干活了...List %v", self.G.Name, self.ID, self.PlayerList)
	// 遍历过程中存在slice的数据删除
	self.lock.RLocker()
	persons := self.PlayerList
	self.lock.RUnlock()
	for _, uid := range persons {
		v, ok := self.GetSitter(uid)
		if ok {
			v.State = protoMsg.PlayerState_PlayerSitDown
			v.Code = SeatNoCode
			v.Gain = INVALID
			v.Score = INVALID
			v.Total = INVALID
			v.Multiple = INVALID
			v.HandCard = make([]byte, 0)
			v.Challenge = make([]int64, 0)
			if v.Timer != nil {
				v.Timer.Stop()
				v.Timer = nil
			}
		}

		if _, ok = GetClientManger().Get(uid); !ok {
			log.Release("[%v:%v]   \t扫地僧出来干活了...%v", self.G.Name, self.ID, uid)
			self.UpdatePlayerList(DEL, uid, true)
		}
	}
	self.Sitter.Range(func(key, value interface{}) bool {
		uid, ok := key.(int64)
		sit, ok1 := value.(*Chair)
		if ok || !ok1 {
			return true
		}
		if sit.Gold < int64(self.T.LessScore) || sit.State != protoMsg.PlayerState_PlayerSitDown && sit.State != protoMsg.PlayerState_PlayerLookOn {
			if sit.Timer != nil {
				sit.Timer.Stop()
				sit.Timer = nil
			}

			GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.ExitGameResp{
				UserID: uid,
				GameID: self.ID,
			})
			if sit.Gold < int64(self.T.LessScore) {
				GetClientManger().SendPopResultX(sit.UserID, FAILED, StatusText[Title001], StatusText[User12])
			}
			self.Sitter.Delete(uid)
			sit = nil
		}
		return true
	})

	self.ChangeState(protoMsg.GameScene_Free)
	self.CorrectCommission()
	self.GetInnings()
	self.Loading()
	return self.ID
}

// ChangeState 改变游戏状态
func (self *Game) ChangeState(state protoMsg.GameScene) {
	self.lock.Lock()
	self.G.Scene = state
	self.TimeStamp = time.Now().Unix()
	self.lock.Unlock()
	log.Debug("[%v:%v]   \t当前状态:%v 桌:%v ", self.G.Name, self.ID, protoMsg.GameScene_name[int32(self.G.Scene)], self.T.Name)
}

// CorrectCommission 修正税收
func (self *Game) CorrectCommission() {
	if self.T.Commission < 50 {
		self.T.Commission = 100 - self.T.Commission
	}
	if self.T.Commission <= 0 || self.T.Commission == 100 {
		self.T.Commission = 100
	}
}

// GetInnings 获取牌局信息
func (self *Game) GetInnings() *Inning {
	if self.InningInfo == nil {
		info := &Inning{}
		info.OpenAreas = make([]byte, 0)
		self.InningInfo = info
	}

	self.InningInfo.Number = strings.ToUpper(util.Md5Sum(time.Now().String())) //GenerateGameNum(, self.Level, int32(self.TID))
	return self.InningInfo
}

// GetPlayerListMsg 获取玩家列表
func (self *Game) GetPlayerListMsg() *protoMsg.PlayerListInfo {
	msg := &protoMsg.PlayerListInfo{}
	msg.AllInfos = make([]*protoMsg.PlayerInfo, 0)
	for _, pid := range self.PlayerList {
		person := GetPlayerManger().Get(pid)
		if person != nil {
			msg.AllInfos = CopyInsert(msg.AllInfos, len(msg.AllInfos), person.ToMsg()).([]*protoMsg.PlayerInfo)
		}

	}

	return msg
}

func (self *Game) GetSitterListMsg() *protoMsg.PlayerListInfo {
	msg := &protoMsg.PlayerListInfo{}
	msg.AllInfos = make([]*protoMsg.PlayerInfo, 0)
	showCount := IndexStart

	self.Sitter.Range(func(key, value interface{}) bool {
		sit, ok := value.(*Chair)
		if !ok {
			return true
		}
		if showCount < ShowChair {
			showCount++
			msg.AllInfos = CopyInsert(msg.AllInfos, len(msg.AllInfos), sit.ToMsg()).([]*protoMsg.PlayerInfo)
		}
		return true
	})
	return msg
}

// GetSitter 获取座位
func (self *Game) GetSitter(userID int64) (*Chair, bool) {
	sit, ok := self.Sitter.Load(userID)
	return sit.(*Chair), ok
}

// GetAllSitter 获取所有坐席
func (self *Game) GetAllSitter() []*Chair {
	list := make([]*Chair, 0)
	self.Sitter.Range(func(key, value interface{}) bool {
		sit, ok := value.(*Chair)
		if !ok {
			return true
		}
		list = CopyInsert(list, len(list), sit).([]*Chair)
		return true
	})
	return list
}

// GetUsefulSitter 获取有用的坐席
func (self *Game) GetUsefulSitter() []*Chair {
	list := make([]*Chair, 0)
	self.Sitter.Range(func(key, value interface{}) bool {
		sit, ok := value.(*Chair)
		if !ok || sit.State < protoMsg.PlayerState_PlayerAgree {
			return true
		}
		list = CopyInsert(list, len(list), sit).([]*Chair)
		return true
	})
	return list
}

// UpdatePlayerList 更新玩家列表
func (self *Game) UpdatePlayerList(code int32, userID int64, autoNotify bool) bool {
	//获取椅子号
	person := GetPlayerManger().Get(userID)
	chair, ok := self.GetSitter(userID)
	if !ok || person == nil {
		//反馈出场场消息
		GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.ExitGameResp{
			UserID: userID,
			GameID: self.ID,
		})

		//删除玩家
		self.DeleteChair(userID)
		self.DeletePlayer(userID)
		if person != nil {
			GetGamesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
			person.Exit()
		}
		log.Debug("[%v:%v]无法更新玩家列表:%v 操作码:%v", self.G.Name, self.ID, userID, code)
		return false
	}

	if code == ADD { //[新增玩家]
		self.AddPlayer(userID)
		GetGamesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
		if autoNotify {
			//反馈入场消息
			GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.EnterGameResp{
				ChairNum: int32(chair.ChairID),
				UserInfo: chair.ToMsg(),
				GameID:   self.ID,
			})
		}

	} else if code == DEL { //[玩家退出]

		if chair.State == protoMsg.PlayerState_PlayerPickUp &&
			self.G.Scene == protoMsg.GameScene_Free && person.Sex == 0x0F {
			self.ReadyCount--
		} else if INVALID != chair.Score || (chair.State == protoMsg.PlayerState_PlayerAgree &&
			self.G.Scene != protoMsg.GameScene_Free && self.G.Scene != protoMsg.GameScene_Over) ||
			(protoMsg.PlayerState_PlayerAgree < chair.State &&
				chair.State != protoMsg.PlayerState_PlayerGiveUp &&
				chair.State != protoMsg.PlayerState_PlayerStandUp &&
				chair.State != protoMsg.PlayerState_PlayerCompareLose) {
			GetClientManger().SendPopResultX(userID, FAILED, StatusText[Title002], StatusText[Game23])
			log.Release("[%v:%v]无法踢除玩家:%v 状态:%v 操作码:%v", self.G.Name, self.ID, userID, chair.State, code)
			return false
		}
		if chair.State == protoMsg.PlayerState_PlayerAgree && 0 < self.ReadyCount &&
			(self.G.Scene == protoMsg.GameScene_Free || self.G.Scene == protoMsg.GameScene_Over) {
			self.ReadyCount--
		}

		//反馈出场场消息
		if autoNotify {
			GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.ExitGameResp{
				UserID: userID,
				GameID: self.ID,
			})
		}

		//删除玩家
		self.DeleteChair(userID)
		self.DeletePlayer(userID)
		GetGamesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
		person.Exit()

		//检测是否可以进行清场
		//if SYSTEMID != self.T.HostID && 0 == self.Sitter.Length() && IndexStart < self.RunCount {
		//	GetGamesManger().ClearTable(self.Reset)
		//}
	}
	log.Release("[%v:%v]\t 更新玩家列表 ID:%v 当前玩家列表:%v 操作码:%v", self.G.Name, self.ID, userID, self.PlayerList, code)

	return true
}

// Trustee 托管
func (self *Game) Trustee(args []interface{}, currentID int64, autoPlay func()) {
	_ = args[0]
	a := args[0].([]interface{})
	agent := a[1].(gate.Agent)
	person := agent.UserData().(*Player)
	m := a[0].(*protoMsg.TrusteeReq)

	if sitter, ok := self.GetSitter(person.UserID); ok {
		// 机器人 则不发起通知
		//if m.IsTrustee != (sitter.State == protoMsg.PlayerState_PlayerPickUp) && person.Sex != 0x0F {
		//	GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.LandLordsTrusteeResp{
		//		UserID:    person.UserID,
		//		IsTrustee: m.IsTrustee,
		//	})
		//}
		if m.IsTrustee {
			if person.UserID == currentID && self.G.Scene == protoMsg.GameScene_Playing && sitter.State != protoMsg.PlayerState_PlayerPickUp {
				sitter.State = protoMsg.PlayerState_PlayerPickUp
				autoPlay()
			} else {
				sitter.State = protoMsg.PlayerState_PlayerPickUp
			}

			log.Release("[%v:%v]\t玩家:%v 托管了", self.G.Name, self.ID, person.UserID)
		} else {
			sitter.State = protoMsg.PlayerState_PlayerPlaying
			log.Release("[%v:%v]\t玩家:%v 取消托管了", self.G.Name, self.ID, person.UserID)
		}
	}

}

// Calculate 玩家结算
func (self *Game) Calculate(sql SQLCalculate, isAndRobot, isDoCommission bool) {
	for _, sitter := range self.GetAllSitter() {
		uid := sitter.UserID
		sitter.Score = INVALID //清空下注积分
		if !isAndRobot && 0x0F == GetPlayerManger().Get(uid).Sex {
			continue
		}
		//(已结算的 不再结算)
		if sitter.Gain != INVALID && sitter.Code != SeatSettled {
			if INVALID < sitter.Gain && !isDoCommission && 0 < self.T.Commission { //赔率包含税收 对未扣除税收的进行扣税
				sitter.Gain = sitter.Gain * int64(self.T.Commission) * 100 / 10000
			}
			//写入数据库
			if money, fact, ok := sql(CalculateInfo{
				UserID:   uid,
				ByUID:    self.T.HostID,
				Gid:      self.ID,
				HostID:   self.T.HostID,
				PreMoney: sitter.Gold,
				Payment:  -sitter.Gain,
				Code:     CodeSettle,
				TypeID:   self.G.Type,
				Kid:      self.G.KindID,
				Level:    self.G.Level,
				Order:    self.InningInfo.Number,
				Remark:   self.G.Name,
			}); ok {
				sitter.Gold = money
				sitter.Gain = -fact
				sitter.Code = CodeSettle
				changeGold := &protoMsg.NotifyBalanceChange{
					UserID:    uid,
					Code:      CodeSettle,
					Coin:      money,
					AlterCoin: sitter.Gain,
					Reason:    self.InningInfo.Number,
				}
				//通知金币变化
				GetClientManger().SendTo(uid, changeGold)
				log.Debug("[%v:%v]\t结算成功:%v 玩家ID:%v 当前金币:%v!!", self.G.Name, self.ID, sitter.Gain, uid, money)
			} else {
				GetClientManger().SendResultX(uid, FAILED, StatusText[Mysql11])
				log.Debug("[%v:%v]\t结算失败:%v 玩家ID:%v 当前金币:%v!! %v", self.G.Name, self.ID, sitter.Gain, uid, money, StatusText[Mysql11])
			}
		}

	}
}

// CalculateOne 仅对某一位进行结算
func (self *Game) CalculateOne(sql SQLCalculate, sitter *Chair, isDoCommission bool) {
	uid := sitter.UserID
	if sitter.Gain != INVALID && sitter.Code != SeatSettled {
		if INVALID < sitter.Gain && !isDoCommission && 0 < self.T.Commission { //赔率包含税收 对未扣除税收的进行扣税
			sitter.Gain = sitter.Gain * int64(self.T.Commission) * 100 / 10000
		}
		//写入数据库
		if money, fact, ok := sql(CalculateInfo{
			UserID:   uid,
			ByUID:    self.T.HostID,
			Gid:      self.ID,
			HostID:   self.T.HostID,
			PreMoney: sitter.Gold,
			Payment:  -sitter.Gain,
			Code:     CodeSettle,
			TypeID:   self.G.Type,
			Kid:      self.G.KindID,
			Level:    self.G.Level,
			Order:    self.InningInfo.Number,
			Remark:   self.G.Name,
		}); ok {
			sitter.Gold = money
			sitter.Gain = -fact
			sitter.Code = CodeSettle
			changeGold := &protoMsg.NotifyBalanceChange{
				UserID:    uid,
				Code:      CodeSettle,
				Coin:      money,
				AlterCoin: sitter.Gain,
				Reason:    self.InningInfo.Number,
			}
			//通知金币变化
			GetClientManger().SendTo(uid, changeGold)
			log.Debug("[%v:%v]\t结算成功:%v 玩家ID:%v 当前金币:%v!!", self.G.Name, self.ID, sitter.Gain, uid, money)
		} else {
			GetClientManger().SendResultX(uid, FAILED, StatusText[Mysql11])
			log.Debug("[%v:%v]\t结算失败:%v 玩家ID:%v 当前金币:%v!! %v", self.G.Name, self.ID, sitter.Gain, uid, money, StatusText[Mysql11])
		}
	}
}

// FlauntWealth 盈利者炫耀
func (self *Game) FlauntWealth() {
	hints := ""
	for _, sitter := range self.GetAllSitter() {
		if sitter != nil && 0 < sitter.Gain {
			player := GetPlayerManger().Get(sitter.UserID)

			hints = fmt.Sprintf("玩家%v在[%v]赢得了%0.2f金币", player.Name, self.G.Name, Unwrap(sitter.Gain, 2))
			for _, pid := range self.PlayerList {
				GetClientManger().SendTo(pid, &protoMsg.NotifyNoticeResp{
					UserID: pid,
					GameID: self.ID,
					Level:  protoMsg.NTFLevel_PraiseNTF,
					TimeInfo: &protoMsg.TimeInfo{
						OutTime:   INVALID,
						WaitTime:  INVALID,
						TotalTime: ShowTime,
						TimeStamp: time.Now().Unix(),
					},
					Title:   StatusText[Title006],
					Content: hints,
				})
			}
		}
	}
}

// GameOver 游戏解散
func (self *Game) GameOver() {
	hints := fmt.Sprintf("非常感谢您的参与!!!\n\n已经达到最大局数:%v局,所有玩家将自动离桌!", self.RunCount)
	GetClientManger().NotifyOthers(self.PlayerList, &protoMsg.ResultPopResp{
		Flag:  SUCCESS,
		Title: StatusText[Title001],
		Hints: hints,
	})
	for _, uid := range self.PlayerList {
		GetClientManger().SendTo(uid, &protoMsg.ExitGameResp{
			UserID: uid,
			GameID: self.ID,
		})
		person := GetPlayerManger().Get(uid)
		person.Exit()
	}
	GetGamesManger().ClearTable(self.Reset)
	self.Sitter.Range(func(key, value interface{}) bool {
		uid, _ := key.(int64)
		self.Sitter.Delete(uid)
		return true
	})
	self = nil
}

//-----------------游戏管理-----------------

// CreateGame  创建游戏(不含实例)
func (self *GamesManger) CreateGame(gid int64, gInfo *protoMsg.GameInfo, tInfo *protoMsg.TableInfo) (*Game, bool) { //添加游戏
	if "" == tInfo.Name {
		tInfo.Name = gInfo.Name
	}
	game := &Game{
		G:          gInfo,
		T:          tInfo,
		ID:         gid,
		PlayerList: make([]int64, 0),
		TimeStamp:  time.Now().Unix(),
		Sitter:     syncmap.Map{},
		InningInfo: nil,
	}

	if game.ID == INVALID {
		log.Debug("不能存放 无效的 游戏ID")
		return nil, false
	}
	log.Release("游戏ID:%v 房主ID[%v] 创建游戏[%v]成功(未启用)", gid, tInfo.HostID, gInfo.Name)
	return self.AddGame(game)
}

// AddGame 添加游戏
func (self *GamesManger) AddGame(game *Game) (*Game, bool) { //添加游戏
	if g, ok := self.GetGame(game.ID); ok {
		return g, false
	}

	if game.ID == INVALID {
		log.Fatal("不能存放 无效的 游戏ID. 信息:%v", game)
		return nil, false
	}
	self.Store(game.ID, game)
	return game, true
}

// SetGame 修改游戏信息
func (self *GamesManger) SetGame(game *Game) (*Game, bool) { //添加游戏
	if game.ID == INVALID {
		log.Fatal("不能存放 无效的 游戏ID. 信息:%v", game)
		return nil, false
	}
	self.Store(game.ID, game)
	return game, true
}

// GetGame 获取游戏
func (self *GamesManger) GetGame(gid int64) (*Game, bool) {
	if game, ok := self.Load(gid); ok {
		return game.(*Game), ok
	}
	return nil, false
}

// CheckHostID 检测房主权限
func (self *GamesManger) CheckHostID(gid, uid int64) bool {
	if game, ok := self.Load(gid); ok {
		if g, ok := game.(*Game); ok {
			return g.T.HostID == uid
		}
	}
	return false
}
func (self *GamesManger) GetGameBase(gid int64) (kindID int32, level int32) {
	if game, ok := self.Load(gid); ok && game != nil {
		return game.(*Game).G.KindID, game.(*Game).G.Level
	}
	return INVALID, INVALID
}

// GetTables 获取牌桌列表
func (self *GamesManger) GetTables(kindID int32, level int32) (tables *protoMsg.TableList) {
	tables = &protoMsg.TableList{}
	tables.Items = make([]*protoMsg.TableItem, 0)
	self.Map.Range(func(index, value interface{}) bool {
		game := value.(*Game)
		if game.G.KindID == kindID && game.G.Level == level {
			//info := *(game.T)
			//info.Password = ""
			//info.Commission = 100 - info.Commission
			item := &protoMsg.TableItem{
				GameID: game.ID,
				//Info:&info,
				Info: &protoMsg.TableInfo{
					HostID:      game.T.HostID,
					Name:        game.T.Name,
					Password:    "",
					State:       game.T.State,
					EnterScore:  game.T.EnterScore,
					LessScore:   game.T.LessScore,
					PlayScore:   game.T.PlayScore,
					Commission:  100 - game.T.Commission,
					MaxChair:    game.G.MaxCount,
					Amount:      game.T.Amount,
					SitterCount: game.T.SitterCount,
				},
			}
			tables.Items = CopyInsert(tables.Items, len(tables.Items), item).([]*protoMsg.TableItem)
		}
		return true
	})
	if 0 == len(tables.Items) {
		return nil
	}

	//排序一下
	sort.Slice(tables.Items, func(i, j int) bool { return tables.Items[i].GameID < tables.Items[j].GameID })
	for k, item := range tables.Items {
		item.Num = int32(k + 1)
	}

	return tables
}

// GetTableNum 获取牌桌编号
func (self *GamesManger) GetTableNum(gid int64, kindID, level int32) int32 {
	if tables := self.GetTables(kindID, level); tables != nil {
		for _, v := range tables.Items {
			if v.GameID == gid {
				return v.Num
			}
		}
	}
	return 0
}

// GetTableSize 获取牌桌编号
func (self *GamesManger) GetTableSize(kindID int32, level int32) int {
	num := INVALID
	self.Range(func(index, value interface{}) bool {
		game := value.(*Game)
		if game.G.KindID == kindID && game.G.Level == level {
			num++
		}
		return true
	})
	return num
}

// GetTable 从已有牌桌中获取桌子
func (self *GamesManger) GetTable(gid int64) *Table {
	for _, table := range self.Tables {
		if table.GameID == gid {
			return table
		}
	}
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// SystemTables  系统创建桌子
func (self *GamesManger) SystemTables(product ProductCallback) {
	//系统房的ID固定
	self.Range(func(key, value interface{}) bool {
		game := value.(*Game)
		instance := product(game)
		if nil == instance {
			return true
		}

		game.T.Name = game.G.Name
		table := &Table{}
		table.GameID = game.ID
		table.Num = int32(self.GetTableSize(game.G.KindID, game.G.Level))
		table.Info = game.T
		table.Instance = instance
		table.Instance.Start(nil)

		self.Tables = CopyInsert(self.Tables, len(self.Tables), table).([]*Table)
		return true
	})

}

// Match 配桌
func (self *GamesManger) Match(game *Game, person *Player, product ProductCallback) (*Table, int) {
	if game == nil || person == nil {
		return nil, Game38
	}

	if nil != person.PtrTable && INVALID != person.PtrTable.GameID {
		//返回原桌
		log.Debug("[配桌]\t GID%v UID:%v 原桌返回", person.PtrTable.GameID, person.UserID)
		return person.PtrTable, SUCCESS
	}

	//无法新增座位
	//game, _ = self.AddGame(game)
	// 非电玩城类的,则为其添加座位
	if game.G.Type != protoMsg.GameType_GamesCity && !game.AddChair(person) {
		log.Debug("[配桌]\t GID%v UID:%v 无法新增座位(已满员)", game.ID, person.UserID)
		// 新增游戏 游戏ID为0,并安排新座位

		return nil, Game39
	}

	//进入场景
	enterScene := func(table *Table) (*Table, int) {
		log.Debug("[%v:%v]\t [准入条件]:(底分:%v 入场积分:%v) 状态:%v--->user:%v money:%v", game.G.Name, game.ID, game.T.LessScore, game.T.EnterScore, game.T.State, person.UserID, person.Gold)
		//检测是否满足准入条件
		isCarRoom := game.G.Type == protoMsg.GameType_TableCard
		if isCarRoom {
			//房卡入场分数不受限
			game.T.EnterScore = INVALID
		} else {
			//非房卡,则使用钱
			person.Gold = person.Money
		}
		// 是否是房主
		//isHost := game.T.HostID == person.UserID
		if game.T.State == protoMsg.TableState_InitTB || game.T.State == protoMsg.TableState_OpenTB { //[1
			if game.T.EnterScore <= int32(person.Gold) && game.T.LessScore <= int32(person.Gold) { //[2
				person.GameID = table.GameID
				person.PtrTable = table

				var args []interface{}
				args = append(args, int32(GameUpdateEnter), person.UserID)
				table.Instance.UpdateInfo(args)
				return table, SUCCESS
			}
			return nil, Game15
		}
		return nil, Game16
	}

	//查找现有桌子
	var table *Table = nil
	for _, tableItem := range self.Tables {
		if tableItem.GameID == game.ID {
			//进入场景
			table = tableItem
			break
		}
	}
	if table == nil {
		game.InningInfo = &Inning{}
		table = &Table{}
		table.GameID = game.ID
		table.Num = int32(self.GetTableSize(game.G.KindID, game.G.Level))
		table.Info = game.T
		table.Instance = product(game)
		table.Instance.Start(nil)
		self.Tables = CopyInsert(self.Tables, len(self.Tables), table).([]*Table)
	}

	return enterScene(table)
}

// ChangeTable 换桌
func (self *GamesManger) ChangeTable(gid int64, person *Player, product ProductCallback) (*Table, int) { //换桌
	if person.State == protoMsg.PlayerState_PlayerStandUp || person.State == protoMsg.PlayerState_PlayerLookOn {
		if temGame, ok1 := self.GetGame(gid); ok1 {
			kid := temGame.G.KindID
			level := temGame.G.Level

			gList := self.GetTables(kid, level)
			for _, item := range gList.Items {
				if item.GameID == person.GameID {
					continue
				}
				if game, ok := self.GetGame(item.GameID); ok && "" == game.T.Password && (game.G.Scene == protoMsg.GameScene_Over || game.G.Scene == protoMsg.GameScene_Free) {
					return self.Match(game, person, product)
				}
			}
		}

	}
	return nil, Room09
}

func (self *GamesManger) control(gid int64, state protoMsg.TableState) (*Game, bool) { //维护
	if v, ok := self.Load(gid); ok {
		item := v.(*Game)
		item.T.State = state
		item.G.Scene = protoMsg.GameScene_Closing
		if table := self.GetTable(item.ID); table != nil && table.Instance != nil {
			var args []interface{}
			args = append(args, state)
			table.Instance.SuperControl(args)
		}
		return item, true
	}
	return nil, false
}

// Start 开启
func (self *GamesManger) Start(gid int64) (*Game, bool) {
	return self.control(gid, protoMsg.TableState_OpenTB)
}

// Maintain 维护
func (self *GamesManger) Maintain(gid int64) (*Game, bool) {
	return self.control(gid, protoMsg.TableState_RepairTB)
}

// Clear 清场
func (self *GamesManger) Clear(gid int64) (*Game, bool) {
	return self.control(gid, protoMsg.TableState_ClearTB)
}

// Close 关闭
func (self *GamesManger) Close(gid int64) (*Game, bool) {
	//需要将游戏删除
	return self.control(gid, protoMsg.TableState_CloseTB)
}

// ClearTable ClearCallback
func (self *GamesManger) ClearTable(clear ClearCallback) bool { //清场
	gid := clear()
	bRet := false
	self.Range(func(key, value interface{}) bool {
		if key.(int64) == gid {
			bRet = true
			self.Delete(key)
			for _, v := range self.Tables {
				if v.GameID == gid {
					v.Instance = nil
					*v = Table{}
					self.Tables = DeleteValue(self.Tables, v).([]*Table)
					return false
				}
			}
		}
		return true
	})
	log.Debug("[清场]\tGameID:%v 清场中(%v)......................   ", gid, bRet)
	return bRet
}

/////////////////////////////////////////////////////////

// UpdateTableList 通知更新牌桌列表 不通知某位玩家
func (self *GamesManger) UpdateTableList(gInfo *protoMsg.GameInfo, butID int64, pageNum int32) int32 {
	//广播给当前游戏种类内的玩家
	maxPage := int32(IndexStart)
	num := pageNum
	if tables := GetGamesManger().GetTables(gInfo.KindID, gInfo.Level); tables != nil {
		msgTbs := &protoMsg.ChooseGameResp{
			Info: gInfo,
			Tables: &protoMsg.TableList{
				Items: make([]*protoMsg.TableItem, 0),
			},
			PageNum: INVALID,
		}
		//仅记录最后一页数据
		if MaxPage < len(tables.Items) {
			items := make([]*protoMsg.TableItem, 0)
			maxPage = IndexStart
			for i := 0; i < len(tables.Items); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						if num == maxPage {
							break
						}
						maxPage++
					}
					items = make([]*protoMsg.TableItem, 0)
				}
				items = CopyInsert(items, len(items), tables.Items[i]).([]*protoMsg.TableItem)

			}
			if 0 < len(items) {
				msgTbs.Tables.Items = items
			}
		} else {
			msgTbs.Tables.Items = tables.Items
		}
		if num == INVALID {
			num = maxPage
		}
		//通知所在得当前玩家更新桌位列表
		allPlayer := GetPlayerManger()
		allPlayer.Range(func(key, value interface{}) bool {
			uid, ok := key.(int64)
			if !ok {
				return true
			}
			man, ok1 := value.(*Player)
			if !ok1 || man.PtrRoom == nil {
				return true
			}
			if INVALID == man.GameID && butID != uid && man.PtrRoom.Kind == gInfo.KindID &&
				man.PtrRoom.Level == gInfo.Level && man.PtrRoom.PageNum == num {
				msgTbs.PageNum = num
				GetClientManger().SendTo(uid, msgTbs)
			}
			return true
		})
	}
	return maxPage
}

// DownCaches 记录游戏当前最大ID
func (self *GamesManger) DownCaches() {
	//GetGamesManger()
	////
	//f, err := os.OpenFile("./log/game.dat", os.O_CREATE|os.O_WRONLY, 0600)
	//defer f.Close()
	//if err != nil {
	//	log.Debug("创建文件出错:%v", err)
	//	return
	//}
	//
	//if self.Length != INVALID {
	//	maxGid := strconv.FormatUint(self.Length, 10)
	//	if maxGid != "" {
	//		f.Write([]byte(maxGid))
	//		f.Sync()
	//		log.Release("落地时 最大GID:%s", maxGid)
	//	}
	//}
}

// 缓存落地
func (self *GamesManger) UpCaches() {
	//GetGamesManger()
	//data, err := ioutil.ReadFile("./log/game.dat")
	//if err != nil {
	//	log.Fatal("%v", err)
	//}
	//self.Length, err = strconv.ParseUint(string(data), 10, 64)
	//if err != nil {
	//	log.Fatal("读取缓存文件:game.dat  失败:%v", err)
	//}
	//
	//log.Release("起步时 最大GID:%v", self.Length)
}
