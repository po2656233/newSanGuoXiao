package player

import (
	clog "github.com/po2656233/superplace/logger"
	cproto "github.com/po2656233/superplace/net/proto"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"time"
)

// p.Call(session.AgentPath, "Response", &protoMsg.Response{})

// 注册接口
func (p *ActorPlayer) registerLocalMsg() {

	p.Local().Register(p.CheckKnapsack)
	p.Local().Register(p.ZhaocaimiaoBet)
	p.Local().Register(p.HowPlay)
	p.Local().Register(p.BaccaratHost)
	p.Local().Register(p.EmailDel)
	p.Local().Register(p.BrtoubaoHost)
	p.Local().Register(p.TigerXdragonHost)
	p.Local().Register(p.JoinGameReadyQueue)
	p.Local().Register(p.ChooseHero)
	p.Local().Register(p.Ping)
	p.Local().Register(p.UpdateGold)
	p.Local().Register(p.MahjongOutCard)
	p.Local().Register(p.FixNickName)
	p.Local().Register(p.BarterYuanBao)
	p.Local().Register(p.GetClassList)
	p.Local().Register(p.ChineseChessSetTime)
	p.Local().Register(p.GetUserID)
	p.Local().Register(p.Trustee)
	p.Local().Register(p.GetBackPassword)
	p.Local().Register(p.TigerXdragonBet)
	p.Local().Register(p.GetGameList)
	p.Local().Register(p.MahjongRoll)
	p.Local().Register(p.EmailRead)
	p.Local().Register(p.ChineseChessAgreeTime)
	p.Local().Register(p.ChineseChessReady)
	p.Local().Register(p.GetAllHero)
	p.Local().Register(p.BarterCoin)
	p.Local().Register(p.NotifyNotice)
	p.Local().Register(p.DisbandedTable)
	p.Local().Register(p.Suggest)
	p.Local().Register(p.MahjongReady)
	p.Local().Register(p.BrcowcowHost)
	p.Local().Register(p.BrTuitongziBet)
	p.Local().Register(p.DeleteTable)
	p.Local().Register(p.Claim)
	p.Local().Register(p.MahjongOperate)
	p.Local().Register(p.SanguoxiaoSwap)
	p.Local().Register(p.UpdateMoney)
	p.Local().Register(p.BaccaratSuperHost)
	p.Local().Register(p.BrTuitongziSuperHost)
	p.Local().Register(p.AddRecord)
	p.Local().Register(p.ChineseChessMove)
	p.Local().Register(p.GetRecharges)
	p.Local().Register(p.Barter)
	p.Local().Register(p.BrcowcowHostList)
	p.Local().Register(p.Recharge)
	p.Local().Register(p.JoinAllReadyQueue)
	p.Local().Register(p.ChooseRoom)
	p.Local().Register(p.Challenge)
	p.Local().Register(p.BrTuitongziHost)
	p.Local().Register(p.BaccaratBet)
	p.Local().Register(p.GetRoomList)
	p.Local().Register(p.GetInningsInfo)
	p.Local().Register(p.DisbandedGame)
	p.Local().Register(p.CheckIn)
	p.Local().Register(p.GetTable)
	p.Local().Register(p.ChangeTable)
	p.Local().Register(p.DrawHero)
	p.Local().Register(p.GetAllGoods)
	p.Local().Register(p.BarterMoney)
	p.Local().Register(p.CreateTable)
	p.Local().Register(p.TigerXdragonSuperHost)
	p.Local().Register(p.GetMyHero)
	p.Local().Register(p.DecreaseGameRun)
	p.Local().Register(p.GameOver)
	p.Local().Register(p.BuyGoods)
	p.Local().Register(p.BrtoubaoSuperHost)
	p.Local().Register(p.CreateRoom)
	p.Local().Register(p.EnterGame)
	p.Local().Register(p.GetGoods)
	p.Local().Register(p.GetTaskList)
	p.Local().Register(p.CheckHero)
	p.Local().Register(p.GetCheckIn)
	p.Local().Register(p.DownHero)
	p.Local().Register(p.BrcowcowBet)
	p.Local().Register(p.ChooseClass)
	p.Local().Register(p.RollDice)
	p.Local().Register(p.GetTableList)
	p.Local().Register(p.BrtoubaoBet)
	p.Local().Register(p.ChooseTable)
	p.Local().Register(p.RankingList)
	p.Local().Register(p.ExitGame)
	p.Local().Register(p.Email)
	p.Local().Register(p.GetUserInfo)
}
func (p *ActorPlayer) getPerson(session *cproto.Session) *mgr.Player {
	p.Session = session
	uid := session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if person == nil {
		// 获取玩家信息
		data, errCode := rpc.SendDataToDB(p.App(), &protoMsg.GetUserInfoReq{Uid: uid})
		if errCode == 0 && data != nil {
			resp, ok := data.(*protoMsg.GetUserInfoResp)
			if ok && resp.Info != nil {
				person = mgr.ToPlayer(resp.Info)
				online.BindPlayer(uid, person.UserID, p.PathString())
				mgr.GetPlayerMgr().Append(person)
			}
		}
	}
	if person == nil {
		userData := p.UserData()
		if userData != nil {
			person = userData.(*mgr.Player)
			mgr.GetPlayerMgr().Append(person)
		} else {
			clog.Errorf("uid:%v There are no instance objects. ", uid)
		}
		return person
	}

	p.SetUserData(person)
	mgr.GetClientMgr().Append(person.UserID, p)

	return person
}

// EnterGame 进入
func (p *ActorPlayer) EnterGame(session *cproto.Session, m *protoMsg.EnterGameReq) {
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
		data, errCode := rpc.SendDataToDB(agent.App(), &protoMsg.GetTableReq{Tid: m.TableID})
		if errCode == 0 {
			if tbResp, ok := data.(*protoMsg.GetTableResp); ok {
				var err error
				tb, err = room.AddTable(tbResp.Info, NewGame)
				if err != nil {
					clog.Warnf("[enter] params %+v  uid:%+v AddTable err:%v", m, uid, err)
					agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room17])
					return
				}
			}
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

func (p *ActorPlayer) JoinGameReadyQueue(session *cproto.Session, m *protoMsg.JoinGameReadyQueueReq) {
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
	if gInfo.State == protoMsg.GameState_InitTB {
		clog.Warnf("[join] [State] params %+v  uid:%+v InitTB", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Game16])
		return
	}
	if gInfo.State == protoMsg.GameState_CloseTB {
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
		data, errCode := rpc.SendDataToDB(agent.App(), &protoMsg.CreateTableReq{
			Rid:        SYSTEMID,
			Gid:        m.GameID,
			PlayScore:  Unlimited,
			Name:       gInfo.Name,
			Opentime:   time.Now().Unix(),
			Commission: INVALID, //系统房没有税收
			MaxRound:   Unlimited,
		})
		if errCode != SUCCESS {
			clog.Warnf("[join] [SendDataToDB] params %+v  uid:%+v FAIL", m, uid)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
		resp, ok1 := data.(*protoMsg.CreateTableResp)
		if !ok1 || resp == nil || resp.Table == nil {
			clog.Warnf("[join]  params %+v  uid:%+v FAIL", m, uid)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
		// 给系统房添加刚创建的牌桌
		t, err := room.AddTable(resp.Table, NewGame)
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
func (p *ActorPlayer) ExitGame(session *cproto.Session, m *protoMsg.ExitGameReq) {
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
	if ok := person.UpdateState(protoMsg.PlayerState_PlayerReady, args); !ok {
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
	if ok := person.UpdateState(protoMsg.PlayerState_PlayerSetTime, args); !ok {
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
func (p *ActorPlayer) Suggest(_ *cproto.Session, _ *protoMsg.SuggestReq) {
	agent := p
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
		msg := &protoMsg.SuggestResp{
			UserID: person.UserID,
			Feedback: &protoMsg.EmailInfo{
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

func (p *ActorPlayer) NotifyNotice(_ *cproto.Session, m *protoMsg.NotifyNoticeReq) {
	agent := p
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
		if m.Timeout <= INVALID {
			m.Timeout = Minute
		}
		msg := &protoMsg.NotifyNoticeResp{
			UserID: person.UserID,
			GameID: m.GameID,
			Level:  m.Level,
			TimeInfo: &protoMsg.TimeInfo{
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

// 获取用户简要信息
func (p *ActorPlayer) GetUserInfo(session *cproto.Session, m *protoMsg.GetUserInfoReq) {
	// TODO: 实现 GetUserInfo 处理逻辑
}

// 游戏消息
// 抢庄
func (p *ActorPlayer) BaccaratHost(session *cproto.Session, m *protoMsg.BaccaratHostReq) {
	// TODO: 实现 BaccaratHost 处理逻辑
}

// 操作
func (p *ActorPlayer) MahjongOperate(session *cproto.Session, m *protoMsg.MahjongOperateReq) {
	// TODO: 实现 MahjongOperate 处理逻辑
}

// ////////////////////////////////////////////
// 准备
func (p *ActorPlayer) MahjongReady(session *cproto.Session, m *protoMsg.MahjongReadyReq) {
	// TODO: 实现 MahjongReady 处理逻辑
}

// 选择桌台
func (p *ActorPlayer) ChooseTable(session *cproto.Session, m *protoMsg.ChooseTableReq) {
	// TODO: 实现 ChooseTable 处理逻辑
}

// 庄家扔骰子
func (p *ActorPlayer) MahjongRoll(session *cproto.Session, m *protoMsg.MahjongRollReq) {
	// TODO: 实现 MahjongRoll 处理逻辑
}

// 游戏消息
// 下注  投注总额 = 投注大小 * 投注倍数 * 基础投注
func (p *ActorPlayer) ZhaocaimiaoBet(session *cproto.Session, m *protoMsg.ZhaocaimiaoBetReq) {
	// TODO: 实现 ZhaocaimiaoBet 处理逻辑
}

// 解散桌位(桌位上的游戏将处于关闭状态)
func (p *ActorPlayer) DisbandedTable(session *cproto.Session, m *protoMsg.DisbandedTableReq) {
	// TODO: 实现 DisbandedTable 处理逻辑
}

// 加入全服游戏准备队列 也是反馈 JoinGameReadyQueueResp
func (p *ActorPlayer) JoinAllReadyQueue(session *cproto.Session, m *protoMsg.JoinAllReadyQueueReq) {
	// TODO: 实现 JoinAllReadyQueue 处理逻辑
}

// 换成元宝
func (p *ActorPlayer) BarterYuanBao(session *cproto.Session, m *protoMsg.BarterYuanBaoReq) {
	// TODO: 实现 BarterYuanBao 处理逻辑
}

// 获取房间列表
func (p *ActorPlayer) GetRoomList(session *cproto.Session, m *protoMsg.GetRoomListReq) {
	// TODO: 实现 GetRoomList 处理逻辑
}

// ///////////[优秀如你]-->Req:请求 Resp:反馈<--[交互专用]///////////////////////////////////
// ///////////[优秀如你]-->Req:请求 Resp:反馈<--[交互专用]///////////////////////////////////
// ///////////////选择操作///////////////////////////////////
// 获取分类列表
func (p *ActorPlayer) GetClassList(session *cproto.Session, m *protoMsg.GetClassListReq) {
	// TODO: 实现 GetClassList 处理逻辑
}

// ////////////////////武将/////////////////////////////////////////
// [招募]抽取武将
func (p *ActorPlayer) DrawHero(session *cproto.Session, m *protoMsg.DrawHeroReq) {
	// TODO: 实现 DrawHero 处理逻辑
}

// 选择房间
func (p *ActorPlayer) ChooseRoom(session *cproto.Session, m *protoMsg.ChooseRoomReq) {
	// TODO: 实现 ChooseRoom 处理逻辑
}

// [布阵]获取我的英雄(武将)
func (p *ActorPlayer) GetMyHero(session *cproto.Session, m *protoMsg.GetMyHeroReq) {
	// TODO: 实现 GetMyHero 处理逻辑
}

// 购买商品
func (p *ActorPlayer) BuyGoods(session *cproto.Session, m *protoMsg.BuyGoodsReq) {
	// TODO: 实现 BuyGoods 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) BrTuitongziBet(session *cproto.Session, m *protoMsg.BrTuitongziBetReq) {
	// TODO: 实现 BrTuitongziBet 处理逻辑
}

// 减少游戏局数
func (p *ActorPlayer) DecreaseGameRun(session *cproto.Session, m *protoMsg.DecreaseGameRunReq) {
	// TODO: 实现 DecreaseGameRun 处理逻辑
}

// /////////////////金币变化///////////////////////////////
// 更新金币
func (p *ActorPlayer) UpdateGold(session *cproto.Session, m *protoMsg.UpdateGoldReq) {
	// TODO: 实现 UpdateGold 处理逻辑
}

// 换成金币
func (p *ActorPlayer) BarterCoin(session *cproto.Session, m *protoMsg.BarterCoinReq) {
	// TODO: 实现 BarterCoin 处理逻辑
}

// 获取单张桌信息
func (p *ActorPlayer) GetTable(session *cproto.Session, m *protoMsg.GetTableReq) {
	// TODO: 实现 GetTable 处理逻辑
}

// /////////////////游戏记录///////////////////////////////
// 添加游戏记录
func (p *ActorPlayer) AddRecord(session *cproto.Session, m *protoMsg.AddRecordReq) {
	// TODO: 实现 AddRecord 处理逻辑
}

// 获取游戏列表
func (p *ActorPlayer) GetGameList(session *cproto.Session, m *protoMsg.GetGameListReq) {
	// TODO: 实现 GetGameList 处理逻辑
}

// 移动棋子
func (p *ActorPlayer) ChineseChessMove(session *cproto.Session, m *protoMsg.ChineseChessMoveReq) {
	// TODO: 实现 ChineseChessMove 处理逻辑
}

// 设置时长
func (p *ActorPlayer) ChineseChessSetTime(session *cproto.Session, m *protoMsg.ChineseChessSetTimeReq) {
	// TODO: 实现 ChineseChessSetTime 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) TigerXdragonSuperHost(session *cproto.Session, m *protoMsg.TigerXdragonSuperHostReq) {
	// TODO: 实现 TigerXdragonSuperHost 处理逻辑
}

// ////////////////////充值///////////////////////////////////////////
// 充值
func (p *ActorPlayer) Recharge(session *cproto.Session, m *protoMsg.RechargeReq) {
	// TODO: 实现 Recharge 处理逻辑
}

// ----------------------------------------------------------------------------------
// 抢庄
func (p *ActorPlayer) BrcowcowHost(session *cproto.Session, m *protoMsg.BrcowcowHostReq) {
	// TODO: 实现 BrcowcowHost 处理逻辑
}

// 获取任务列表
func (p *ActorPlayer) GetTaskList(session *cproto.Session, m *protoMsg.GetTaskListReq) {
	// TODO: 实现 GetTaskList 处理逻辑
}

// [选将]或[更换]选择携带的武将
func (p *ActorPlayer) ChooseHero(session *cproto.Session, m *protoMsg.ChooseHeroReq) {
	// TODO: 实现 ChooseHero 处理逻辑
}

// 请求游戏结束（注:返回牌局记录后,游戏随即销毁）
func (p *ActorPlayer) GameOver(session *cproto.Session, m *protoMsg.GameOverReq) {
	// TODO: 实现 GameOver 处理逻辑
}

// 查看背包
func (p *ActorPlayer) CheckKnapsack(session *cproto.Session, m *protoMsg.CheckKnapsackReq) {
	// TODO: 实现 CheckKnapsack 处理逻辑
}

// ////////////////邮箱建议////////////////////////////////////
// 邮箱信息
func (p *ActorPlayer) Email(session *cproto.Session, m *protoMsg.EmailReq) {
	// TODO: 实现 Email 处理逻辑
}

// // 注册开发帐号
// message RegisterDevReq {
// string accountName = 1; // 帐号名
// string password = 2;    // 密码
// string ip = 3;          // ip地址
// }
// 用户信息
func (p *ActorPlayer) GetUserID(session *cproto.Session, m *protoMsg.GetUserIDReq) {
	// TODO: 实现 GetUserID 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BrtoubaoSuperHost(session *cproto.Session, m *protoMsg.BrtoubaoSuperHostReq) {
	// TODO: 实现 BrtoubaoSuperHost 处理逻辑
}

// 配置游戏(房卡模式)
func (p *ActorPlayer) CreateTable(session *cproto.Session, m *protoMsg.CreateTableReq) {
	// TODO: 实现 CreateTable 处理逻辑
}

// 查找英雄
func (p *ActorPlayer) CheckHero(session *cproto.Session, m *protoMsg.CheckHeroReq) {
	// TODO: 实现 CheckHero 处理逻辑
}

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) TigerXdragonHost(session *cproto.Session, m *protoMsg.TigerXdragonHostReq) {
	// TODO: 实现 TigerXdragonHost 处理逻辑
}

// 下注
func (p *ActorPlayer) BaccaratBet(session *cproto.Session, m *protoMsg.BaccaratBetReq) {
	// TODO: 实现 BaccaratBet 处理逻辑
}

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) BrTuitongziHost(session *cproto.Session, m *protoMsg.BrTuitongziHostReq) {
	// TODO: 实现 BrTuitongziHost 处理逻辑
}

// ///////////玩家行为(与游戏弱相关的行为)/////////////////////
// [排行榜]
func (p *ActorPlayer) RankingList(session *cproto.Session, m *protoMsg.RankingListReq) {
	// TODO: 实现 RankingList 处理逻辑
}

// 修改用户昵称
func (p *ActorPlayer) FixNickName(session *cproto.Session, m *protoMsg.FixNickNameReq) {
	// TODO: 实现 FixNickName 处理逻辑
}

// 登录
func (p *ActorPlayer) Login(session *cproto.Session, m *protoMsg.LoginReq) {
	// TODO: 实现 Login 处理逻辑
}

// 获取牌局记录 注: 房主权限
func (p *ActorPlayer) GetInningsInfo(session *cproto.Session, m *protoMsg.GetInningsInfoReq) {
	// TODO: 实现 GetInningsInfo 处理逻辑
}

// //////////////选择分类//////////////////////////////////////
// 选择分类
func (p *ActorPlayer) ChooseClass(session *cproto.Session, m *protoMsg.ChooseClassReq) {
	// TODO: 实现 ChooseClass 处理逻辑
}

// 换桌
func (p *ActorPlayer) ChangeTable(session *cproto.Session, m *protoMsg.ChangeTableReq) {
	// TODO: 实现 ChangeTable 处理逻辑
}

// 获取全部英雄(武将)
func (p *ActorPlayer) GetAllHero(session *cproto.Session, m *protoMsg.GetAllHeroReq) {
	// TODO: 实现 GetAllHero 处理逻辑
}

// 交换
func (p *ActorPlayer) SanguoxiaoSwap(session *cproto.Session, m *protoMsg.SanguoxiaoSwapReq) {
	// TODO: 实现 SanguoxiaoSwap 处理逻辑
}

// ////////////////////////////////
// 准备
func (p *ActorPlayer) ChineseChessReady(session *cproto.Session, m *protoMsg.ChineseChessReadyReq) {
	// TODO: 实现 ChineseChessReady 处理逻辑
}

// ////////////////////签到///////////////////////////////////////////
// 签到
func (p *ActorPlayer) CheckIn(session *cproto.Session, m *protoMsg.CheckInReq) {
	// TODO: 实现 CheckIn 处理逻辑
}

// 物品转换 仅支持通用房卡和超级房卡的转换,游戏房卡之间不能置换
func (p *ActorPlayer) Barter(session *cproto.Session, m *protoMsg.BarterReq) {
	// TODO: 实现 Barter 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BrTuitongziSuperHost(session *cproto.Session, m *protoMsg.BrTuitongziSuperHostReq) {
	// TODO: 实现 BrTuitongziSuperHost 处理逻辑
}

// 删除桌牌
func (p *ActorPlayer) DeleteTable(session *cproto.Session, m *protoMsg.DeleteTableReq) {
	// TODO: 实现 DeleteTable 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) BrtoubaoBet(session *cproto.Session, m *protoMsg.BrtoubaoBetReq) {
	// TODO: 实现 BrtoubaoBet 处理逻辑
}

// 获取充值纪录
func (p *ActorPlayer) GetRecharges(session *cproto.Session, m *protoMsg.GetRechargesReq) {
	// TODO: 实现 GetRecharges 处理逻辑
}

// 扔骰子
func (p *ActorPlayer) RollDice(session *cproto.Session, m *protoMsg.RollDiceReq) {
	// TODO: 实现 RollDice 处理逻辑
}

// 查看游戏玩法
func (p *ActorPlayer) HowPlay(session *cproto.Session, m *protoMsg.HowPlayReq) {
	// TODO: 实现 HowPlay 处理逻辑
}

// 超级抢庄
func (p *ActorPlayer) BaccaratSuperHost(session *cproto.Session, m *protoMsg.BaccaratSuperHostReq) {
	// TODO: 实现 BaccaratSuperHost 处理逻辑
}

// [下阵]选择携带的武将
func (p *ActorPlayer) DownHero(session *cproto.Session, m *protoMsg.DownHeroReq) {
	// TODO: 实现 DownHero 处理逻辑
}

// 待上庄列表
func (p *ActorPlayer) BrcowcowHostList(session *cproto.Session, m *protoMsg.BrcowcowHostListReq) {
	// TODO: 实现 BrcowcowHostList 处理逻辑
}

// ///////////预留协议///////////////////////////
// 抢庄
func (p *ActorPlayer) BrtoubaoHost(session *cproto.Session, m *protoMsg.BrtoubaoHostReq) {
	// TODO: 实现 BrtoubaoHost 处理逻辑
}

// 换成money
func (p *ActorPlayer) BarterMoney(session *cproto.Session, m *protoMsg.BarterMoneyReq) {
	// TODO: 实现 BarterMoney 处理逻辑
}

// 读取邮件
func (p *ActorPlayer) EmailRead(session *cproto.Session, m *protoMsg.EmailReadReq) {
	// TODO: 实现 EmailRead 处理逻辑
}

// ///////////[优秀如你]-->Req:请求 Resp:反馈<--[交互专用]///////////////////////////////////
// 走web服时,该项忽略
// 注册
func (p *ActorPlayer) Register(session *cproto.Session, m *protoMsg.RegisterReq) {
	// TODO: 实现 Register 处理逻辑
}

// 托管[暂保留]
func (p *ActorPlayer) Trustee(session *cproto.Session, m *protoMsg.TrusteeReq) {
	// TODO: 实现 Trustee 处理逻辑
}

// 获取所有商品信息
func (p *ActorPlayer) GetAllGoods(session *cproto.Session, m *protoMsg.GetAllGoodsReq) {
	// TODO: 实现 GetAllGoods 处理逻辑
}

// 重连
func (p *ActorPlayer) Reconnect(session *cproto.Session, m *protoMsg.ReconnectReq) {
	// TODO: 实现 Reconnect 处理逻辑
}

// 更新余额
func (p *ActorPlayer) UpdateMoney(session *cproto.Session, m *protoMsg.UpdateMoneyReq) {
	// TODO: 实现 UpdateMoney 处理逻辑
}

// 获取牌桌列表(同获取游戏列表)
func (p *ActorPlayer) GetTableList(session *cproto.Session, m *protoMsg.GetTableListReq) {
	// TODO: 实现 GetTableList 处理逻辑
}

// ////////////heart//////////////////////////////////////////
// 心跳包 默认20秒 网关等待读取数据时长为35秒
func (p *ActorPlayer) Ping(session *cproto.Session, m *protoMsg.PingReq) {
	// TODO: 实现 Ping 处理逻辑
}

// 领取奖励
func (p *ActorPlayer) Claim(session *cproto.Session, m *protoMsg.ClaimReq) {
	// TODO: 实现 Claim 处理逻辑
}

// 获取签到
func (p *ActorPlayer) GetCheckIn(session *cproto.Session, m *protoMsg.GetCheckInReq) {
	// TODO: 实现 GetCheckIn 处理逻辑
}

// 玩家打出去的牌
func (p *ActorPlayer) MahjongOutCard(session *cproto.Session, m *protoMsg.MahjongOutCardReq) {
	// TODO: 实现 MahjongOutCard 处理逻辑
}

// ////////////////////////////////////////////////
// 创建房间---------------
func (p *ActorPlayer) CreateRoom(session *cproto.Session, m *protoMsg.CreateRoomReq) {
	// TODO: 实现 CreateRoom 处理逻辑
}

// //////////////////////////////////////////////////
// [挑战]
func (p *ActorPlayer) Challenge(session *cproto.Session, m *protoMsg.ChallengeReq) {
	// TODO: 实现 Challenge 处理逻辑
}

// 游戏消息
// 下注
func (p *ActorPlayer) TigerXdragonBet(session *cproto.Session, m *protoMsg.TigerXdragonBetReq) {
	// TODO: 实现 TigerXdragonBet 处理逻辑
}

// 解散游戏
func (p *ActorPlayer) DisbandedGame(session *cproto.Session, m *protoMsg.DisbandedGameReq) {
	// TODO: 实现 DisbandedGame 处理逻辑
}

// 找回游戏密码 [创建者|群主]权限
func (p *ActorPlayer) GetBackPassword(session *cproto.Session, m *protoMsg.GetBackPasswordReq) {
	// TODO: 实现 GetBackPassword 处理逻辑
}

// 是否同意对方设置的时长
func (p *ActorPlayer) ChineseChessAgreeTime(session *cproto.Session, m *protoMsg.ChineseChessAgreeTimeReq) {
	// TODO: 实现 ChineseChessAgreeTime 处理逻辑
}

// 删除邮件
func (p *ActorPlayer) EmailDel(session *cproto.Session, m *protoMsg.EmailDelReq) {
	// TODO: 实现 EmailDel 处理逻辑
}

// 下注
func (p *ActorPlayer) BrcowcowBet(session *cproto.Session, m *protoMsg.BrcowcowBetReq) {
	// TODO: 实现 BrcowcowBet 处理逻辑
}

// ///////////////物品/////////////////////////
// 获取商品信息
func (p *ActorPlayer) GetGoods(session *cproto.Session, m *protoMsg.GetGoodsReq) {
	// TODO: 实现 GetGoods 处理逻辑
}
