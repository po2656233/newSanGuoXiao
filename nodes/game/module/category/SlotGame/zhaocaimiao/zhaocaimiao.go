package zhaocaimiao

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	"runtime"
	"strings"
	protoMsg "superman/internal/protocol/go_file/common"
	mgr "superman/nodes/game/manger"
	"time"
)

// 招财喵

type row int32
type col int32

// ZhaocaimiaoGame 继承于GameItem
type ZhaocaimiaoGame struct {
	*mgr.Game
	//lk sync.Mutex
}

//
//func init() {
//	game := NewGame(&Game{
//		G: &protoMsg.GameInfo{
//			Name: "招财喵",
//		},
//		ID: 5001,
//		T: &protoMsg.TableInfo{
//			HostID: 0,
//		},
//	})
//	args := make([]interface{}, 0)
//	args = append(args, &protoMsg.ZhaocaimiaoBetReq{
//		BetAmount: 10,
//		BetRate:   20,
//	})
//	agent := &ClientAgent{}
//	agent.SetUserData(&Player{
//		PlayerInfo: &protoMsg.PlayerInfo{
//			UserID: 1,
//			Money:  10000000,
//			Gold:   8000000,
//		},
//	})
//	args = append(args, agent)
//	game.Playing(args)
//}

// NewGame 创建实例
func NewGame(game *Game) *ZhaocaimiaoGame {
	p := &ZhaocaimiaoGame{
		Game: game,
	}

	p.Init()
	return p
}

// ---------------------------%v:%v----------------------------------------------//
// 初始化信息
func (self *ZhaocaimiaoGame) Init() {
	//需要每局都重置的数据
	self.IsStart = false // 是否启动
	self.IsClear = false // 不进行清场
}

// Scene 场 景
func (self *ZhaocaimiaoGame) Scene(args []interface{}) {
	_ = args[1]
	level := args[0].(int32)
	agent := args[1].(mgr.Agent)
	userData := agent.UserData()
	if userData == nil {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game02])
		log.Release("[Error][%v:%v场景] [未能查找到相关玩家] ", self.G.Name, self.ID)
		return
	}
	person := userData.(*Player)
	if level != self.G.Level {
		GlobalSender.SendResult(agent, FAILED, StatusText[Game01])
		log.Release("[Error][%v:%v场景] [玩家ID:%v 房间等级不匹配] ", self.G.Name, self.ID, person.UserID)
		return
	}

	//场景信息
	StateInfo := &protoMsg.ZhaocaimiaoSceneResp{}
	StateInfo.Inning = self.InningInfo.Number
	StateInfo.BaseBet = BaseBet
	StateInfo.MyInfo = person.ToMsg()

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

	//反馈场景信息
	GlobalSender.SendData(agent, StateInfo)

	//通知游戏状态
	timeInfo := &protoMsg.TimeInfo{}
	timeInfo.TimeStamp = self.TimeStamp
	timeInfo.OutTime = int32(time.Now().Unix() - self.TimeStamp)

	switch person.GameState {
	case int32(protoMsg.GameScene_Start): // 开始
		timeInfo.TotalTime = Date
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.ZhaocaimiaoStateStartResp{
			Times: timeInfo,
			//Inning: self.InningInfo.Number,
		})
	case int32(protoMsg.GameScene_Playing): // 下注
		timeInfo.TotalTime = Date
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		GlobalSender.SendData(agent, &protoMsg.ZhaocaimiaoStatePlayingResp{
			Times: timeInfo,
		})
	case int32(protoMsg.GameScene_Opening): // 开奖
		timeInfo.TotalTime = WAITTIME
		timeInfo.WaitTime = timeInfo.TotalTime - timeInfo.OutTime
		resp := &protoMsg.ZhaocaimiaoOpenResp{}
		_ = BytesToPB([]byte(person.BindInfo), resp)
		GlobalSender.SendData(agent, &protoMsg.ZhaocaimiaoStateOpenResp{
			Times:    timeInfo,
			OpenInfo: resp,
		})
	}
}

// Start 开 始
func (self *ZhaocaimiaoGame) Start(args []interface{}) {
	_ = args
	log.Release("[%v:%v]游戏創建成功", self.G.Name, self.ID)
	self.Reset()
	//self.onPlay()
}

// Playing 游 戏(下分|下注)
func (self *ZhaocaimiaoGame) Playing(args []interface{}) {
	_ = args[1]
	//【消息】
	m := args[0].(*protoMsg.ZhaocaimiaoBetReq)
	//【传输对象】
	agent := args[1].(mgr.Agent)

	userData := agent.UserData()
	if nil == userData {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[User02], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[User02])
		return
	}

	person := userData.(*Player)
	if m.BetAmount <= 0 || m.BetRate <= 0 {
		log.Release("[%v:%v][%v] Playing:->uid:%v %v ", self.G.Name, self.ID, StatusText[Game07], person.UserID, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game07])
		return
	}
	if person.GameState == int32(protoMsg.GameScene_Opening) {
		log.Release("[%v:%v][%v] Playing:->%v ", self.G.Name, self.ID, StatusText[Game04], m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game04])
		return
	}
	// 下注
	msg := &protoMsg.ZhaocaimiaoBetResp{}
	msg.BetScore = m.BetAmount*int64(m.BetRate) + BaseBet
	if person.Gold <= msg.BetScore {
		log.Release("[%v:%v][%v] Playing:->uid:%v %v ", self.G.Name, self.ID, StatusText[Game05], person.UserID, m)
		GlobalSender.SendResult(agent, FAILED, StatusText[Game05])
		return
	}
	person.Gold -= msg.BetScore
	GlobalSender.SendData(agent, msg)
	log.Release("[%v:%v] [playing] -> uid:%v Bet(amount:%v rate:%v) BetScore:%v ", self.G.Name, self.ID, person.UserID,
		m.BetAmount, int64(m.BetRate), msg.BetScore)

	// 开奖
	self.OnOpen(person)

}

// Over 结 算
func (self *ZhaocaimiaoGame) Over(args []interface{}) {
	_ = args
}

// UpdateInfo 更新信息
func (self *ZhaocaimiaoGame) UpdateInfo(args []interface{}) {
	_ = args[1]

	flag := args[0].(int32)
	userID := args[1].(int64)
	log.Debug("[%v:%v]\t更新信息:%v->*** %v\n", self.G.Name, self.ID, userID, RecvMessage[int(flag)])
	switch flag {
	case GameUpdateOut: //玩家离开
		//self.UpdatePlayerList(DEL, userID, true)
	case GameUpdateEnter: //更新玩家列表
		//self.UpdatePlayerList(ADD, userID, true)
	case GameUpdateHost: //更新玩家抢庄信息
	case GameUpdateSuperHost: //更新玩家超级抢庄信息
	case GameUpdateOffline: //更新玩家超级抢庄信息
	case GameUpdateReconnect: //更新玩家超级抢庄信息
	case GameUpdateDisbanded:
		self.IsClear = true
	case GameUpdateGold:
		_ = args[2]
		GlobalSender.NotifyOthers(self.PlayerList, args[2].(*protoMsg.NotifyBalanceChange))
	}
}

// SuperControl 超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)
func (self *ZhaocaimiaoGame) SuperControl(args []interface{}) {
	_ = args[0]
	flag := args[0].(int32)
	switch flag {
	case GameControlClear:
		self.IsClear = true
	}
}

// OnOpen 开奖
func (self *ZhaocaimiaoGame) OnOpen(person *Player) {
	round := int32(0)
	fid := int64(0)
	person.GameState = int32(protoMsg.GameScene_Opening)
	openAward := &protoMsg.ZhaocaimiaoOpenResp{
		FreeRound:    0,
		TotalScore:   0,
		AllFreeRound: 0,
		TimeStamp:    time.Now().Unix(),
		Inning:       "",
		Awards:       make([]*protoMsg.ZhaocaimiaoOpenInfo, 0),
	}

	// 查找上一次的游戏记录 是否仍属于免费摇奖
	last, err1 := GlobalSqlHandle.CheckGameRecordLast(self.ID, person.UserID)
	if err1 != nil {
		log.Error("[%v:%v]\t玩家ID:%v 获取上一次游戏记录失败:%v !!", self.G.Name, self.ID, person.UserID, err1)
		log.Release("[%v:%v][%v] Opening err:%v", self.G.Name, self.ID, StatusText[Mysql16], err1)
		GlobalSender.SendResultX(person.UserID, FAILED, StatusText[Mysql16])
		person.GameState = int32(protoMsg.GameScene_Start)
		return
	}

	if 0 < last.RemainFree {
		// 父级ID
		fid = last.FatherID
		if fid == 0 {
			fid = last.ID
		}
		// 免费次数-1
		openAward.AllFreeRound = last.RemainFree - 1
		openAward.Inning = last.Inning
		round = last.Rounds
	} else {
		// 如果上把没有免费次数,则表示本轮非免费局
		openAward.Inning = self.GetInnings().Number
	}

	// 1.创建首个结果线
	lines := CreateLines(ColCount, RowCount)
	//self.linesLog(false, "原生线(按列排序)\t", lines)

	//变框
	rate := ChangeRate
	cLines := changeKuang(rate, lines)
	// 中奖 且金银框|百搭测试
	//cLines = map[col][]row{
	//	0: {0, 9, 13, 8, 4, 4},
	//	1: {8, 12, 6, 7, 13, 1},
	//	2: {5, 10, 278, -1, -1, 13},
	//	3: {6, 1, 6, 269, -1, -1},
	//	4: {3, 4, 8, -1, 11, 1},
	//	5: {0, 6, 13, 6, 12, 1},
	//}
	//// 夺宝测试
	//cLines = map[col][]row{
	//	0: {0, 11, 12, 8, 1, 8},
	//	1: {3, 13, 12, 5, 82, -1},
	//	2: {3, 12, 13, 6, 11, 4},
	//	3: {1, 7, 275, -1, -1, -1},
	//	4: {5, 2, 10, 9, 8, 10},
	//	5: {0, 2, 13, 12, 1, 9},
	//}
	self.linesLog(false, "起始线(按列排序)\t", lines)

	// 检测中奖 直到没有中

	for i := 0; i < MaxLoop; i++ {
		rLines := sortRows(cLines)
		openInfo := &protoMsg.ZhaocaimiaoOpenInfo{
			IsFreeOpen: false,
			AwardScore: 0,
			AwardRate:  0,
			Open:       make([]*protoMsg.ZhaocaimiaoCell, 0),
			Cells:      make([]*protoMsg.ZhaocaimiaoEraseCell, 0),
		}
		self.linesLog(true, fmt.Sprintf("第%v轮 栅格\t\t", i+1), rLines)

		// 获取开奖信息
		openInfo.Open = self.getOpenCells(rLines)
		openAward.Awards = append(openAward.Awards, openInfo)

		// 检测中奖
		awards := checkAward(cLines)
		if 0 == len(awards) { // 没有中奖 则退出
			break
		}

		// 消除元素
		strHints := ""
		if openInfo.Cells, strHints = self.getEraseCells(awards); strHints != "" {
			log.Release("[%v:%v] 第%v轮 玩家:%v 中奖 %v", self.G.Name, self.ID, i+1, person.UserID, strHints)
		}

		if i+1 == MaxLoop {
			rate = 100
		}
		newLines, freeCount, err := eraseRule(rate, cLines, awards)
		if err != nil {
			log.Error("[%v:%v] 消除失败 err:%v", self.G.Name, self.ID, err)
			break
		}

		if 0 < freeCount { // 存在免费次数,则退出
			openAward.FreeRound = int32(freeCount)
			log.Release("[%v:%v] 新增免费局数 :%v ", self.G.Name, self.ID, freeCount)
			break
		}

		// 更新下一轮的主线
		cLines = newLines
	}
	openDB, _ := PBToBytes(openAward)
	log.Release("[%v:%v][open] 包体大小:%v -> %v ", self.G.Name, self.ID, len(openDB), openAward)
	round++
	openAward.AllFreeRound += openAward.FreeRound
	GlobalSender.SendTo(person.UserID, openAward)
	person.BindInfo = string(openDB)
	// 结算并记录开奖信息

	// 反馈开奖记录
	self.TimeStamp = time.Now().Unix()
	GlobalSender.SendTo(person.UserID, &protoMsg.ZhaocaimiaoStateOpenResp{
		Times: &protoMsg.TimeInfo{
			TimeStamp: self.TimeStamp,
			OutTime:   0,
			WaitTime:  WAITTIME,
			TotalTime: WAITTIME,
		},
		OpenInfo: openAward,
	})

	go self.OnOver_G(round, fid, person, openAward)
}

// OnOver_G 结束并结算
func (self *ZhaocaimiaoGame) OnOver_G(round int32, fatherId int64, person *Player, openResp *protoMsg.ZhaocaimiaoOpenResp) {
	defer func() {
		<-time.After(WAITTIME * time.Second)
		if person != nil {
			person.GameState = int32(protoMsg.GameScene_Start)
			self.TimeStamp = time.Now().Unix()
			GlobalSender.SendTo(person.UserID, &protoMsg.ZhaocaimiaoStatePlayingResp{
				Times: &protoMsg.TimeInfo{
					TimeStamp: self.TimeStamp,
					OutTime:   0,
					WaitTime:  WAITTIME,
					TotalTime: WAITTIME,
				},
			})
		}
	}()

	// 记录游戏结果
	strResult := ""
	for _, award := range openResp.Awards {
		for _, cell := range award.Open {
			strResult += fmt.Sprintf("%v,%v/%v|", cell.Row, cell.Col, cell.Element)
		}
		//strResult += "\n"
	}
	strResult = strings.Trim(strResult, "|")

	uid := person.UserID
	if _, err := GlobalSqlHandle.AddGameRecord(model.GameRecord{
		Gid:        self.ID, //游戏ID
		UID:        uid,
		FatherID:   fatherId,
		RemainFree: openResp.AllFreeRound,
		Inning:     openResp.Inning,
		Rounds:     round,
		IsFree:     round != 1,
		Result:     strResult,
	}); err != nil {
		log.Error("[%v:%v]\t 游戏记录失败 玩家:%v 牌局号:%v", self.G.Name, self.ID, uid, openResp.Inning)
		return
	}

	if money, fact, ok := GlobalSqlHandle.DeductMoney(CalculateInfo{
		UserID:   uid,
		ByUID:    self.T.HostID,
		Gid:      self.ID,
		HostID:   self.T.HostID,
		PreMoney: person.Gold,
		Payment:  -openResp.TotalScore,
		Code:     CodeSettle,
		TypeID:   self.G.Type,
		Kid:      self.G.KindID,
		Level:    self.G.Level,
		Order:    self.InningInfo.Number,
		Remark:   self.G.Name,
	}); ok {
		person.Gold = money
		changeGold := &protoMsg.NotifyBalanceChange{
			UserID:    uid,
			Code:      CodeSettle,
			Coin:      money,
			AlterCoin: fact,
			Reason:    fmt.Sprintf("结算 牌局:%v", self.InningInfo.Number),
		}
		//通知金币变化
		GetClientManger().SendTo(person.UserID, changeGold)
		log.Debug("[%v:%v]\t结算成功:%v 玩家ID:%v 当前金币:%v!!", self.G.Name, self.ID, fact, uid, money)
	} else {
		GetClientManger().SendResultX(uid, FAILED, StatusText[Mysql11])
		log.Error("[%v:%v]\t结算失败:%v 玩家ID:%v 当前金币:%v!! %v", self.G.Name, self.ID, fact, uid, money, StatusText[Mysql11])
	}
	// 退出协程 该函数可以确保defer被执行
	runtime.Goexit()
}

// ///////////////////////////////////////////////////////////////////
// getOpenCells
func (self *ZhaocaimiaoGame) getOpenCells(lines map[col][]row) []*protoMsg.ZhaocaimiaoCell {
	openCells := make([]*protoMsg.ZhaocaimiaoCell, 0)
	for ri := 0; ri < RowCount; ri++ {
		for rj, e := range lines[col(ri)] {
			// 占位符不记录
			if 0 < e {
				openCells = append(openCells, &protoMsg.ZhaocaimiaoCell{
					Row:     int32(ri),
					Col:     int32(rj),
					Element: int32(e),
				})
			}
		}
	}
	return openCells
}

func (self *ZhaocaimiaoGame) getEraseCells(awards []AwardInfo) ([]*protoMsg.ZhaocaimiaoEraseCell, string) {
	eraseCells := make([]*protoMsg.ZhaocaimiaoEraseCell, 0)
	strShow := ""
	for index, award := range awards {
		eraseCell := &protoMsg.ZhaocaimiaoEraseCell{
			Index: int32(award.Index),
			Sames: make([]*protoMsg.ZhaocaimiaoCell, 0),
		}
		for _, po := range award.Pos {
			eraseCell.Sames = append(eraseCell.Sames, &protoMsg.ZhaocaimiaoCell{
				Row:     int32(po.Rx),
				Col:     int32(po.Cy),
				Element: award.Element,
			})
		}
		eraseCells = append(eraseCells, eraseCell)
		strShow += fmt.Sprintf("\t第一列的索引:%v 位置:%v ", index, award)
	}
	return eraseCells, strShow
}

func (self *ZhaocaimiaoGame) linesLog(showOriginal bool, hints string, lines map[col][]row) {
	strShow := hints
	for i := 0; i < ColCount; i++ {
		strShow += "-->"
		for _, c := range lines[col(i)] {
			element := int32(c)
			if showOriginal {
				element = originalElement(element)
			}
			strShow += fmt.Sprintf("%v\t", element)
		}
	}
	log.Release("[%v:%v] %v ", self.G.Name, self.ID, strShow)
}
