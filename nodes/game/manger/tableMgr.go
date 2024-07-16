package manger

import (
	"fmt"
	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
	"sort"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	"superman/nodes/leaf/jettengame/gamedata/goclib/util"
	"sync"
	"time"
)

// Package manger
/*
# -*- coding: utf-8 -*-
# @Time : 2020/4/28 9:07
# @Author : Pitter
# @File : games.go
# @Software: GoLand
*/

// TablesManger 管理房间 TablesIndex *Table
type TablesManger struct {
	sync.Map
	Tables []*Table
}

// //////////////////平台管理////////////////////////////////////
var gamesManger *TablesManger = nil
var gamesOnce sync.Once

// GetTablesManger 平台管理对象(单例模式)//manger.persons = make(map[int32]*Player)
func GetTablesManger() *TablesManger {
	gamesOnce.Do(func() {
		gamesManger = &TablesManger{
			Map: sync.Map{},
		}
	})
	return gamesManger
}

// -----------------桌子-----------------

// AddPlayer 加入玩家列表
func (self *Table) AddPlayer(userID int64) {
	self.Lock()
	defer self.Unlock()

}

// DeletePlayer 从玩家列表中踢出
func (self *Table) DeletePlayer(userID int64) {
	self.Lock()
	defer self.Unlock()

}

// AddChair 添加座位
func (self *Table) AddChair(person *Player) bool {

	return true
}

func (self *Table) NextChair(curChair *Chair) (*Chair, bool) {
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

func (self *Table) NextChairX(uid int64) (*Chair, bool) {

	return nil, false
}

func (self *Table) DeleteChair(uid int64) {
	sit, ok := self.GetSitter(uid)
	if ok && sit.Timer != nil {
		sit.Timer.Stop()
	}

}

// Loading 加载配置 [废弃:不加载公用配置,拆分至子游戏内部实现]
func (self *Table) Loading() {
	//获取文件的
	//settingFile := conf.GameYamlPath
	//fInfo, _ := os.Stat(settingFile)
	//lastFixTime := fInfo.ModTime().Unix()
	// 在文件被修改或配置内容为空时，则读取配置信息

}

// Reset 重置座位上的玩家
func (self *Table) Reset() int64 {

	self.ChangeState(protoMsg.GameScene_Free)
	self.CorrectCommission()
	self.GetInnings()
	self.Loading()
	return self.ID
}

// ChangeState 改变游戏状态
func (self *Table) ChangeState(state protoMsg.GameScene) {
	self.Lock()
	self.G.Scene = state
	self.TimeStamp = time.Now().Unix()
	self.Unlock()
	log.Debug("[%v:%v]   \t当前状态:%v 桌:%v ", self.G.Name, self.ID, protoMsg.GameScene_name[int32(self.G.Scene)], self.T.Name)
}

// CorrectCommission 修正税收
func (self *Table) CorrectCommission() {
	if self.T.Commission < 50 {
		self.T.Commission = 100 - self.T.Commission
	}
	if self.T.Commission <= 0 || self.T.Commission == 100 {
		self.T.Commission = 100
	}
}

// GetInnings 获取牌局信息
func (self *Table) GetInnings() *Inning {
	if self.InningInfo == nil {
		info := &Inning{}
		info.OpenAreas = make([]byte, 0)
		self.InningInfo = info
	}

	self.InningInfo.Number = strings.ToUpper(util.Md5Sum(time.Now().String())) //GenerateGameNum(, self.Level, int32(self.TID))
	return self.InningInfo
}

// GetPlayerListMsg 获取玩家列表
func (self *Table) GetPlayerListMsg() *protoMsg.PlayerListInfo {
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

func (self *Table) GetSitterListMsg() *protoMsg.PlayerListInfo {
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
func (self *Table) GetSitter(userID int64) (*Chair, bool) {
	sit, ok := self.Sitter.Load(userID)
	return sit.(*Chair), ok
}

// GetAllSitter 获取所有坐席
func (self *Table) GetAllSitter() []*Chair {
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
func (self *Table) GetUsefulSitter() []*Chair {
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
func (self *Table) UpdatePlayerList(code int32, userID int64, autoNotify bool) bool {
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
			GetTablesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
			person.Exit()
		}
		log.Debug("[%v:%v]无法更新玩家列表:%v 操作码:%v", self.G.Name, self.ID, userID, code)
		return false
	}

	if code == ADD { //[新增玩家]
		self.AddPlayer(userID)
		GetTablesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
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
		GetTablesManger().UpdateTableList(self.G, userID, person.PtrRoom.PageNum)
		person.Exit()

		//检测是否可以进行清场
		//if SYSTEMID != self.T.HostID && 0 == self.Sitter.Length() && IndexStart < self.RunCount {
		//	GetTablesManger().ClearTable(self.Reset)
		//}
	}
	log.Release("[%v:%v]\t 更新玩家列表 ID:%v 当前玩家列表:%v 操作码:%v", self.G.Name, self.ID, userID, self.PlayerList, code)

	return true
}

// Trustee 托管
func (self *Table) Trustee(args []interface{}, currentID int64, autoPlay func()) {
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
func (self *Table) Calculate(sql SQLCalculate, isAndRobot, isDoCommission bool) {
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
func (self *Table) CalculateOne(sql SQLCalculate, sitter *Chair, isDoCommission bool) {
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
func (self *Table) FlauntWealth() {
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
func (self *Table) GameOver() {
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
	GetTablesManger().ClearTable(self.Reset)
	self.Sitter.Range(func(key, value interface{}) bool {
		uid, _ := key.(int64)
		self.Sitter.Delete(uid)
		return true
	})
	self = nil
}

//-----------------游戏管理-----------------

// CreateGame  创建游戏(不含实例)
func (self *TablesManger) CreateGame(gid int64, gInfo *protoMsg.GameInfo, tInfo *protoMsg.TableInfo) (*Table, bool) { //添加游戏
	if "" == tInfo.Name {
		tInfo.Name = gInfo.Name
	}
	table := &Table{
		G:          gInfo,
		T:          tInfo,
		ID:         gid,
		PlayerList: make([]int64, 0),
		TimeStamp:  time.Now().Unix(),
		Sitter:     syncmap.Map{},
		InningInfo: nil,
	}

	if table.ID == INVALID {
		log.Debug("不能存放 无效的 游戏ID")
		return nil, false
	}
	log.Release("游戏ID:%v 房主ID[%v] 创建游戏[%v]成功(未启用)", gid, tInfo.HostID, gInfo.Name)
	return self.AddGame(table)
}

// AddGame 添加游戏
func (self *TablesManger) AddGame(table *Table) (*Table, bool) { //添加游戏
	if g, ok := self.GetGame(table.ID); ok {
		return g, false
	}

	if table.ID == INVALID {
		log.Fatal("不能存放 无效的 游戏ID. 信息:%v", table)
		return nil, false
	}
	self.Store(table.ID, table)
	return table, true
}

// SetGame 修改游戏信息
func (self *TablesManger) SetGame(table *Table) (*Table, bool) { //添加游戏
	if table.ID == INVALID {
		log.Fatal("不能存放 无效的 游戏ID. 信息:%v", table)
		return nil, false
	}
	self.Store(table.ID, table)
	return table, true
}

// GetGame 获取游戏
func (self *TablesManger) GetGame(gid int64) (*Table, bool) {
	if table, ok := self.Load(gid); ok {
		return table.(*Table), ok
	}
	return nil, false
}

// CheckHostID 检测房主权限
func (self *TablesManger) CheckHostID(gid, uid int64) bool {
	if table, ok := self.Load(gid); ok {
		if g, ok := table.(*Table); ok {
			return g.T.HostID == uid
		}
	}
	return false
}
func (self *TablesManger) GetGameBase(gid int64) (kindID int32, level int32) {
	if table, ok := self.Load(gid); ok && table != nil {
		return table.(*Table).G.KindID, table.(*Table).G.Level
	}
	return INVALID, INVALID
}

// GetTables 获取牌桌列表
func (self *TablesManger) GetTables(kindID int32, level int32) (tables *protoMsg.TableList) {
	tables = &protoMsg.TableList{}
	tables.Items = make([]*protoMsg.TableItem, 0)
	self.Map.Range(func(index, value interface{}) bool {
		table := value.(*Table)
		if table.G.KindID == kindID && table.G.Level == level {
			//info := *(table.T)
			//info.Password = ""
			//info.Commission = 100 - info.Commission
			item := &protoMsg.TableItem{
				GameID: table.ID,
				//Info:&info,
				Info: &protoMsg.TableInfo{
					HostID:      table.T.HostID,
					Name:        table.T.Name,
					Password:    "",
					State:       table.T.State,
					EnterScore:  table.T.EnterScore,
					LessScore:   table.T.LessScore,
					PlayScore:   table.T.PlayScore,
					Commission:  100 - table.T.Commission,
					MaxChair:    table.G.MaxCount,
					Amount:      table.T.Amount,
					SitterCount: table.T.SitterCount,
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
func (self *TablesManger) GetTableNum(gid int64, kindID, level int32) int32 {
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
func (self *TablesManger) GetTableSize(kindID int32, level int32) int {
	num := INVALID
	self.Range(func(index, value interface{}) bool {
		table := value.(*Table)
		if table.G.KindID == kindID && table.G.Level == level {
			num++
		}
		return true
	})
	return num
}

// GetTable 从已有牌桌中获取桌子
func (self *TablesManger) GetTable(gid int64) *Table {
	for _, table := range self.Tables {
		if table.GameID == gid {
			return table
		}
	}
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// SystemTables  系统创建桌子
func (self *TablesManger) SystemTables(product ProductCallback) {
	//系统房的ID固定
	self.Range(func(key, value interface{}) bool {
		table := value.(*Table)
		instance := product(table)
		if nil == instance {
			return true
		}

		table.T.Name = table.G.Name
		table := &Table{}
		table.GameID = table.ID
		table.Num = int32(self.GetTableSize(table.G.KindID, table.G.Level))
		table.Info = table.T
		table.Instance = instance
		table.Instance.Start(nil)

		self.Tables = CopyInsert(self.Tables, len(self.Tables), table).([]*Table)
		return true
	})

}

// Match 配桌
func (self *TablesManger) Match(table *Table, person *Player, product ProductCallback) (*Table, int) {
	if table == nil || person == nil {
		return nil, Game38
	}

	if nil != person.PtrTable && INVALID != person.PtrTable.GameID {
		//返回原桌
		log.Debug("[配桌]\t GID%v UID:%v 原桌返回", person.PtrTable.GameID, person.UserID)
		return person.PtrTable, SUCCESS
	}

	//无法新增座位
	//table, _ = self.AddGame(table)
	// 非电玩城类的,则为其添加座位
	if table.G.Type != protoMsg.GameType_TablesCity && !table.AddChair(person) {
		log.Debug("[配桌]\t GID%v UID:%v 无法新增座位(已满员)", table.ID, person.UserID)
		// 新增游戏 游戏ID为0,并安排新座位

		return nil, Game39
	}

	//进入场景
	enterScene := func(table *Table) (*Table, int) {
		log.Debug("[%v:%v]\t [准入条件]:(底分:%v 入场积分:%v) 状态:%v--->user:%v money:%v", table.G.Name, table.ID, table.T.LessScore, table.T.EnterScore, table.T.State, person.UserID, person.Gold)
		//检测是否满足准入条件
		isCarRoom := table.G.Type == protoMsg.GameType_TableCard
		if isCarRoom {
			//房卡入场分数不受限
			table.T.EnterScore = INVALID
		} else {
			//非房卡,则使用钱
			person.Gold = person.Money
		}
		// 是否是房主
		//isHost := table.T.HostID == person.UserID
		if table.T.State == protoMsg.TableState_InitTB || table.T.State == protoMsg.TableState_OpenTB { //[1
			if table.T.EnterScore <= int32(person.Gold) && table.T.LessScore <= int32(person.Gold) { //[2
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
		if tableItem.GameID == table.ID {
			//进入场景
			table = tableItem
			break
		}
	}
	if table == nil {
		table.InningInfo = &Inning{}
		table = &Table{}
		table.GameID = table.ID
		table.Num = int32(self.GetTableSize(table.G.KindID, table.G.Level))
		table.Info = table.T
		table.Instance = product(table)
		table.Instance.Start(nil)
		self.Tables = CopyInsert(self.Tables, len(self.Tables), table).([]*Table)
	}

	return enterScene(table)
}

// ChangeTable 换桌
func (self *TablesManger) ChangeTable(gid int64, person *Player, product ProductCallback) (*Table, int) { //换桌
	if person.State == protoMsg.PlayerState_PlayerStandUp || person.State == protoMsg.PlayerState_PlayerLookOn {
		if temGame, ok1 := self.GetGame(gid); ok1 {
			kid := temGame.G.KindID
			level := temGame.G.Level

			gList := self.GetTables(kid, level)
			for _, item := range gList.Items {
				if item.GameID == person.GameID {
					continue
				}
				if table, ok := self.GetGame(item.GameID); ok && "" == table.T.Password && (table.G.Scene == protoMsg.GameScene_Over || table.G.Scene == protoMsg.GameScene_Free) {
					return self.Match(table, person, product)
				}
			}
		}

	}
	return nil, Room09
}

func (self *TablesManger) control(gid int64, state protoMsg.TableState) (*Table, bool) { //维护
	if v, ok := self.Load(gid); ok {
		item := v.(*Table)
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
func (self *TablesManger) Start(gid int64) (*Table, bool) {
	return self.control(gid, protoMsg.TableState_OpenTB)
}

// Maintain 维护
func (self *TablesManger) Maintain(gid int64) (*Table, bool) {
	return self.control(gid, protoMsg.TableState_RepairTB)
}

// Clear 清场
func (self *TablesManger) Clear(gid int64) (*Table, bool) {
	return self.control(gid, protoMsg.TableState_ClearTB)
}

// Close 关闭
func (self *TablesManger) Close(gid int64) (*Table, bool) {
	//需要将游戏删除
	return self.control(gid, protoMsg.TableState_CloseTB)
}

// ClearTable ClearCallback
func (self *TablesManger) ClearTable(clear ClearCallback) bool { //清场
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
func (self *TablesManger) UpdateTableList(gInfo *protoMsg.GameInfo, butID int64, pageNum int32) int32 {
	//广播给当前游戏种类内的玩家
	maxPage := int32(IndexStart)
	num := pageNum
	if tables := GetTablesManger().GetTables(gInfo.KindID, gInfo.Level); tables != nil {
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
func (self *TablesManger) DownCaches() {
	//GetTablesManger()
	////
	//f, err := os.OpenFile("./log/table.dat", os.O_CREATE|os.O_WRONLY, 0600)
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
func (self *TablesManger) UpCaches() {
	//GetTablesManger()
	//data, err := ioutil.ReadFile("./log/table.dat")
	//if err != nil {
	//	log.Fatal("%v", err)
	//}
	//self.Length, err = strconv.ParseUint(string(data), 10, 64)
	//if err != nil {
	//	log.Fatal("读取缓存文件:table.dat  失败:%v", err)
	//}
	//
	//log.Release("起步时 最大GID:%v", self.Length)
}
