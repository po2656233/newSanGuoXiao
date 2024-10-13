package player

import (
	clog "github.com/po2656233/superplace/logger"
	cproto "github.com/po2656233/superplace/net/proto"
	. "superman/internal/constant"
	commMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"
	gateMsg "superman/internal/protocol/go_file/gate"
	sqlmodel "superman/internal/sql_model/minigame"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"time"
)

// 注册接口
func (p *ActorPlayer) registerLocalMsg() {
	p.Local().Register(p.BrTuitongziHost)
	p.Local().Register(p.DisbandedTable)
	p.Local().Register(p.RankingList)
	p.Local().Register(p.ChineseChessAgreeTime)
	p.Local().Register(p.ChineseChessReady)
	p.Local().Register(p.AddRecord)
	p.Local().Register(p.MahjongOutCard)
	p.Local().Register(p.BrcowcowBet)
	p.Local().Register(p.RollDice)
	p.Local().Register(p.ChineseChessMove)
	p.Local().Register(p.BrTuitongziBet)
	p.Local().Register(p.JoinAllReadyQueue)
	p.Local().Register(p.BaccaratBet)
	p.Local().Register(p.GetBackPassword)
	p.Local().Register(p.BrtoubaoHost)
	p.Local().Register(p.DecreaseGameRun)
	p.Local().Register(p.UpdateGold)
	p.Local().Register(p.TigerXdragonHost)
	p.Local().Register(p.Trustee)
	p.Local().Register(p.Challenge)
	p.Local().Register(p.TigerXdragonSuperHost)
	p.Local().Register(p.TigerXdragonBet)
	p.Local().Register(p.BrcowcowHost)
	p.Local().Register(p.ExitGame)
	p.Local().Register(p.MahjongReady)
	p.Local().Register(p.GetInningsInfo)
	p.Local().Register(p.SanguoxiaoSwap)
	p.Local().Register(p.MahjongRoll)
	p.Local().Register(p.DisbandedGame)
	p.Local().Register(p.BrcowcowHostList)
	p.Local().Register(p.BaccaratHost)
	p.Local().Register(p.BrtoubaoBet)
	p.Local().Register(p.ChangeTable)
	p.Local().Register(p.MahjongOperate)
	p.Local().Register(p.GameOver)
	p.Local().Register(p.JoinGameReadyQueue)
	p.Local().Register(p.BrTuitongziSuperHost)
	p.Local().Register(p.ZhaocaimiaoBet)
	p.Local().Register(p.EnterGame)
	p.Local().Register(p.BaccaratSuperHost)
	p.Local().Register(p.BrtoubaoSuperHost)
	p.Local().Register(p.ChineseChessSetTime)
}
func (p *ActorPlayer) getPerson(session *cproto.Session) *mgr.Player {
	p.Session = session
	uid := session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if person == nil {
		// 获取玩家信息
		resp, err := mgr.GetDBCmpt().GetUserInfo(uid)
		if err != nil {
			clog.Errorf("[ActorPlayer] getPerson GetUserInfo err:%v", err)
			return nil
		}
		person = mgr.ToPlayer(resp.Info)
		online.BindPlayer(uid, person.UserID, p.PathString())
		mgr.GetPlayerMgr().Append(person)
	}
	p.SetUserData(person)
	mgr.GetClientMgr().Append(person.UserID, p)

	return person
}

// EnterGame 进入
func (p *ActorPlayer) EnterGame(session *cproto.Session, m *gameMsg.EnterGameReq) {
	clog.Debugf("EnterGame: sid:%v uid:%v mid:%v req:%v", session.Sid, session.Uid, session.Mid, m)
	agent := p
	uid := session.Uid
	person := p.getPerson(session)
	// 玩家仍在游戏中
	if person.GameHandle != nil {
		person.GameHandle.Scene([]interface{}{person})
		return
	}

	room := mgr.GetRoomMgr().GetRoom(m.RoomID)
	if room == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room16])
		return
	}
	var tb *mgr.Table
	tb = room.GetTable(m.TableID)
	if tb == nil {
		tbInfo, err := mgr.GetDBCmpt().CheckTable(m.TableID)
		if err != nil {
			agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room16])
			return
		}
		tb, err = room.AddTable(&commMsg.TableInfo{
			Id:         tbInfo.ID,
			Name:       tbInfo.Name,
			Rid:        tbInfo.Rid,
			Gid:        tbInfo.Gid,
			Commission: tbInfo.Commission,
			MaxRound:   tbInfo.Maxround,
			Remain:     tbInfo.Remain,
			MaxSitter:  tbInfo.MaxSitter,
			PlayScore:  tbInfo.Playscore,
			OpenTime:   tbInfo.Opentime,
		}, NewGame)
		if err != nil {
			clog.Warnf("[enter] params %+v  uid:%+v AddTable err:%v", m, uid, err)
			agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room17])
			return
		}

	}
	if tb == nil {
		clog.Warnf("[enter] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[TableInfo09])
		return
	}

	person.GameHandle = tb.GameHandle
	person.Enter([]interface{}{m, p})
}

func (p *ActorPlayer) JoinGameReadyQueue(session *cproto.Session, m *gameMsg.JoinGameReadyQueueReq) {
	agent := p
	uid := session.Uid
	clog.Debugf("[JoinGameReadyQueue]params %+v  uid:%+v", m, uid)
	person := p.getPerson(session)
	if person == nil {
		clog.Errorf("[JoinGameReadyQueue] uid:%v There are no instance objects. ", uid)
		return
	}
	// 玩家仍在游戏中
	if person.GameHandle != nil {
		clog.Warnf("[join] [GameHandle] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Login09])
		person.GameHandle.Scene([]interface{}{person})
		return
	}
	// 获取房间句柄
	room := mgr.GetRoomMgr().GetRoom(m.RoomID)
	if room == nil {
		clog.Warnf("[join] [room] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Room16])
		return
	}

	// 检测游戏是否可用
	gInfo := mgr.GetGameInfoMgr().GetGame(m.GameID)
	if gInfo == nil {
		clog.Errorf("[join] [gInfo] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[TableInfo03])
		return
	}
	if gInfo.State == commMsg.GameState_InitTB {
		clog.Warnf("[join] [State] params %+v  uid:%+v InitTB", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Game16])
		return
	}
	if gInfo.State == commMsg.GameState_CloseTB {
		clog.Warnf("[join] [State] params %+v  uid:%+v CloseTB", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Game19])
		return
	}
	// 玩家加入游戏准备列表
	if ok := room.AddWait(m.GameID, person); !ok {
		clog.Warnf("[join] [GameHandle] params %+v  uid:%+v exist", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Login09])
		return
	}
	// 检测房间的游戏是否满员
	if ok := mgr.GetRoomMgr().GetRoom(m.RoomID).CheckGameFull(m.GameID); ok {
		if m.RoomID != SYSTEMID {
			// 牌桌不足
			clog.Warnf("[join] [CheckGameFull] params %+v  uid:%+v no enought", m, uid)
			agent.SendResult(FAILED, StatusText[TableInfo10])
			return
		}
		// 系统房主动添加牌桌。 如果没有加入队列,并且是系统房,则创建牌桌
		// 请求数据库创建牌桌
		id, maxSit, err := mgr.GetDBCmpt().AddTable(sqlmodel.Table{
			Gid:        m.GameID,
			Rid:        SYSTEMID,
			Name:       gInfo.Name,
			Playscore:  Unlimited,
			Commission: INVALID, //系统房没有税收
			Maxround:   Unlimited,
			Remain:     Unlimited,
			Opentime:   time.Now().Unix(),
		})
		if err != nil {
			clog.Warnf("[join] [SendDataToDB] params %+v  uid:%+v err:%v", m, uid, err)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
		// 给系统房添加刚创建的牌桌
		tbInfo := &commMsg.TableInfo{
			Id:         id,
			Gid:        m.GameID,
			Rid:        SYSTEMID,
			Name:       gInfo.Name,
			PlayScore:  Unlimited,
			Commission: INVALID, //系统房没有税收
			MaxRound:   Unlimited,
			Remain:     Unlimited,
			OpenTime:   time.Now().Unix(),
			MaxSitter:  maxSit,
		}
		t, err := room.AddTable(tbInfo, NewGame)
		if err != nil || t == nil {
			clog.Warnf("[join] [AddTable] params %+v  uid:%+v FAIL. err:%v", m, uid, err)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
	}

	// 成功添加队列
	agent.SendResult(SUCCESS, StatusText[User30])
}

// ExitGame 退出游戏
func (p *ActorPlayer) ExitGame(session *cproto.Session, m *gameMsg.ExitGameReq) {
	//agent := p
	clog.Debugf("[exit]params %+v  uid:%+v", m, session.Uid)
	mgr.GetPlayerMgr().Get(session.Uid).Exit()
}

func ready(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	uid := agent.Session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if ok := person.UpdateState(commMsg.PlayerState_PlayerReady, args); !ok {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User02])
	}
}

// 设置时长
func setTime(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	uid := agent.Session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if ok := person.UpdateState(commMsg.PlayerState_PlayerSetTime, args); !ok {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User02])
	}
}

// 玩法操作
func playing(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	person := mgr.GetPlayerMgr().Get(agent.Session.Uid)
	if person.GameHandle != nil {
		person.GameHandle.Playing(args)
	}
}

// Suggest 意见反馈
func (p *ActorPlayer) Suggest(_ *cproto.Session, _ *gateMsg.SuggestReq) {
	agent := p
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
		msg := &gateMsg.SuggestResp{
			UserID: person.UserID,
			Feedback: &commMsg.EmailInfo{
				//EmailID:    GetPChatID(),
				TimeStamp:  time.Now().Unix(),
				AcceptName: person.Name,
				Topic:      "感谢信",
				Content:    "由衷感谢您的反馈,您的反馈是我们持续的动力!",
				Sender:     "系统邮件",
			},
		}
		mgr.GetClientMgr().SendData(agent, msg)
		mgr.GetClientMgr().SendPopResult(agent, SUCCESS, StatusText[Title007], StatusText[User22])
	}
}

func (p *ActorPlayer) NotifyNotice(_ *cproto.Session, m *gateMsg.NotifyNoticeReq) {
	agent := p
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
		if m.Timeout <= INVALID {
			m.Timeout = Minute
		}
		msg := &gateMsg.NotifyNoticeResp{
			UserID: person.UserID,
			GameID: m.GameID,
			Level:  m.Level,
			TimeInfo: &commMsg.TimeInfo{
				OutTime:   INVALID,
				WaitTime:  INVALID,
				TotalTime: m.Timeout,
				TimeStamp: time.Now().Unix(),
			},
			Title:   m.Title,
			Content: m.Content,
		}

		//写到缓存
		//RedisClient.Set("mini_Notice", msg, time.Duration(m.Timeout)*time.Second)

		//if person.UserID == SYSTEMID {
		//	msg.TimeInfo.TotalTime = Date
		//}
		mgr.GetClientMgr().NotifyAll(msg)
		//GetClientManger().SendPopResult(agent, SUCCESS, StatusText[Title001], StatusText[User22])
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) BrTuitongziHost(session *cproto.Session, m *gameMsg.BrTuitongziHostReq) {
	// TODO: 实现 BrTuitongziHost 处理逻辑
}

// 解散桌位(桌位上的游戏将处于关闭状态)
func (p *ActorPlayer) DisbandedTable(session *cproto.Session, m *gameMsg.DisbandedTableReq) {
	// TODO: 实现 DisbandedTable 处理逻辑
}

// ///////////玩家行为(与游戏弱相关的行为)/////////////////////
// [排行榜]
func (p *ActorPlayer) RankingList(session *cproto.Session, m *gameMsg.RankingListReq) {
	// TODO: 实现 RankingList 处理逻辑
}

// 是否同意对方设置的时长
func (p *ActorPlayer) ChineseChessAgreeTime(session *cproto.Session, m *gameMsg.ChineseChessAgreeTimeReq) {
	// TODO: 实现 ChineseChessAgreeTime 处理逻辑
}

// ////////////////////////////////
// 准备
func (p *ActorPlayer) ChineseChessReady(session *cproto.Session, m *gameMsg.ChineseChessReadyReq) {
	// TODO: 实现 ChineseChessReady 处理逻辑
}

// /////////////////游戏记录///////////////////////////////
// 添加游戏记录
func (p *ActorPlayer) AddRecord(session *cproto.Session, m *gameMsg.AddRecordReq) {
	// TODO: 实现 AddRecord 处理逻辑
}

// 玩家打出去的牌
func (p *ActorPlayer) MahjongOutCard(session *cproto.Session, m *gameMsg.MahjongOutCardReq) {
	// TODO: 实现 MahjongOutCard 处理逻辑
}

// 下注
func (p *ActorPlayer) BrcowcowBet(session *cproto.Session, m *gameMsg.BrcowcowBetReq) {
	// TODO: 实现 BrcowcowBet 处理逻辑
}

// 扔骰子
func (p *ActorPlayer) RollDice(session *cproto.Session, m *gameMsg.RollDiceReq) {
	// TODO: 实现 RollDice 处理逻辑
}

// 移动棋子
func (p *ActorPlayer) ChineseChessMove(session *cproto.Session, m *gameMsg.ChineseChessMoveReq) {
	// TODO: 实现 ChineseChessMove 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) BrTuitongziBet(session *cproto.Session, m *gameMsg.BrTuitongziBetReq) {
	// TODO: 实现 BrTuitongziBet 处理逻辑
}

// 加入全服游戏准备队列 也是反馈 JoinGameReadyQueueResp
func (p *ActorPlayer) JoinAllReadyQueue(session *cproto.Session, m *gameMsg.JoinAllReadyQueueReq) {
	// TODO: 实现 JoinAllReadyQueue 处理逻辑
}

// 下注
func (p *ActorPlayer) BaccaratBet(session *cproto.Session, m *gameMsg.BaccaratBetReq) {
	// TODO: 实现 BaccaratBet 处理逻辑
}

// 找回游戏密码 [创建者|群主]权限
func (p *ActorPlayer) GetBackPassword(session *cproto.Session, m *gameMsg.GetBackPasswordReq) {
	// TODO: 实现 GetBackPassword 处理逻辑
}

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) BrtoubaoHost(session *cproto.Session, m *gameMsg.BrtoubaoHostReq) {
	// TODO: 实现 BrtoubaoHost 处理逻辑
}

// 减少游戏局数
func (p *ActorPlayer) DecreaseGameRun(session *cproto.Session, m *gameMsg.DecreaseGameRunReq) {
	// TODO: 实现 DecreaseGameRun 处理逻辑
}

// /////////////////金币变化///////////////////////////////
// 更新金币
func (p *ActorPlayer) UpdateGold(session *cproto.Session, m *gameMsg.UpdateGoldReq) {
	// TODO: 实现 UpdateGold 处理逻辑
}

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) TigerXdragonHost(session *cproto.Session, m *gameMsg.TigerXdragonHostReq) {
	// TODO: 实现 TigerXdragonHost 处理逻辑
}

// 托管[暂保留]
func (p *ActorPlayer) Trustee(session *cproto.Session, m *gameMsg.TrusteeReq) {
	// TODO: 实现 Trustee 处理逻辑
}

// //////////////////////////////////////////////////
// [挑战]
func (p *ActorPlayer) Challenge(session *cproto.Session, m *gameMsg.ChallengeReq) {
	// TODO: 实现 Challenge 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) TigerXdragonSuperHost(session *cproto.Session, m *gameMsg.TigerXdragonSuperHostReq) {
	// TODO: 实现 TigerXdragonSuperHost 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) TigerXdragonBet(session *cproto.Session, m *gameMsg.TigerXdragonBetReq) {
	// TODO: 实现 TigerXdragonBet 处理逻辑
}

// ----------------------------------------------------------------------------------
// 抢庄
func (p *ActorPlayer) BrcowcowHost(session *cproto.Session, m *gameMsg.BrcowcowHostReq) {
	// TODO: 实现 BrcowcowHost 处理逻辑
}

// ////////////////////////////////////////////
// 准备
func (p *ActorPlayer) MahjongReady(session *cproto.Session, m *gameMsg.MahjongReadyReq) {
	// TODO: 实现 MahjongReady 处理逻辑
}

// 获取牌局记录 注: 房主权限
func (p *ActorPlayer) GetInningsInfo(session *cproto.Session, m *gameMsg.GetInningsInfoReq) {
	// TODO: 实现 GetInningsInfo 处理逻辑
}

// 交换
func (p *ActorPlayer) SanguoxiaoSwap(session *cproto.Session, m *gameMsg.SanguoxiaoSwapReq) {
	// TODO: 实现 SanguoxiaoSwap 处理逻辑
}

// 庄家扔骰子
func (p *ActorPlayer) MahjongRoll(session *cproto.Session, m *gameMsg.MahjongRollReq) {
	// TODO: 实现 MahjongRoll 处理逻辑
}

// 解散游戏
func (p *ActorPlayer) DisbandedGame(session *cproto.Session, m *gameMsg.DisbandedGameReq) {
	// TODO: 实现 DisbandedGame 处理逻辑
}

// 待上庄列表
func (p *ActorPlayer) BrcowcowHostList(session *cproto.Session, m *gameMsg.BrcowcowHostListReq) {
	// TODO: 实现 BrcowcowHostList 处理逻辑
}

// 游戏消息
// 抢庄
func (p *ActorPlayer) BaccaratHost(session *cproto.Session, m *gameMsg.BaccaratHostReq) {
	// TODO: 实现 BaccaratHost 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) BrtoubaoBet(session *cproto.Session, m *gameMsg.BrtoubaoBetReq) {
	// TODO: 实现 BrtoubaoBet 处理逻辑
}

// 换桌
func (p *ActorPlayer) ChangeTable(session *cproto.Session, m *gameMsg.ChangeTableReq) {
	// TODO: 实现 ChangeTable 处理逻辑
}

// 操作
func (p *ActorPlayer) MahjongOperate(session *cproto.Session, m *gameMsg.MahjongOperateReq) {
	// TODO: 实现 MahjongOperate 处理逻辑
}

// 请求游戏结束（注:返回牌局记录后,游戏随即销毁）
func (p *ActorPlayer) GameOver(session *cproto.Session, m *gameMsg.GameOverReq) {
	// TODO: 实现 GameOver 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BrTuitongziSuperHost(session *cproto.Session, m *gameMsg.BrTuitongziSuperHostReq) {
	// TODO: 实现 BrTuitongziSuperHost 处理逻辑
}

// 游戏消息
// 下注  投注总额 = 投注大小 * 投注倍数 * 基础投注
func (p *ActorPlayer) ZhaocaimiaoBet(session *cproto.Session, m *gameMsg.ZhaocaimiaoBetReq) {
	// TODO: 实现 ZhaocaimiaoBet 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BaccaratSuperHost(session *cproto.Session, m *gameMsg.BaccaratSuperHostReq) {
	// TODO: 实现 BaccaratSuperHost 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BrtoubaoSuperHost(session *cproto.Session, m *gameMsg.BrtoubaoSuperHostReq) {
	// TODO: 实现 BrtoubaoSuperHost 处理逻辑
}

// 设置时长
func (p *ActorPlayer) ChineseChessSetTime(session *cproto.Session, m *gameMsg.ChineseChessSetTimeReq) {
	// TODO: 实现 ChineseChessSetTime 处理逻辑
}
