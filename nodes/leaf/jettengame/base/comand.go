package base

/*游戏中的重要标识*/

// 游戏kindID

const (
	Zhaocaimiao = 5001 //
	Mahjong     = 3002
	Sanguoxiao  = 1010
)

// 源码中重要标识
const (
	INVALID    = 0  //无效(切记有效初始化,不要从零开始)
	FAILED     = 1  // 失败
	Fault      = -1 // 故障
	Default    = 1  // 默认
	SUCCESS    = 0  // 成功
	IndexStart = 1  // 起始索引

	SuperRoom   = 1002 // 超级房卡
	GeneralRoom = 1001 // 普通房卡

	ADD = 0 // 新增
	DEL = 1 // 删除

	Lose = 0 // 输
	Win  = 1 // 赢
	Draw = 2 // 平局
)

// 限制标识 | 范围标识
const (
	MaxLoadNum   = 20       // 最大录单记录
	MaxInning    = 20       // 牌局数目
	MaxPage      = 20       // 最大分页
	MaxWait      = 20       // 最长等待时长
	MaxRounds    = 10000    // 最大牌局数
	MaxSystemNum = 99999999 // 系统最大使用数值

	SYSTEMID = 0  // 系统标识
	NameLen  = 6  // 名字长度
	Limit    = 99 // 限制

	ShowTime  = 3  // 展示效果时长
	ShowChair = 20 // 最多展示座位

	TenTwice = 10 // 十倍
	TwoTwice = 2  // 两倍
)

// 房间等级
const (
	RoomGeneral = iota //普通
	RoomMiddle         //中级
	RoomHigh           //高级
)

// 最大椅子数
const (
	MaxChairMore = iota
	OneChair
	TwoChair
	ThreeChair
	FourChair
	FiveChair
	SixChair
	SevenChair
	EightChair
	NineChair
	TenChair
	HundredChair  = 100
	ThousandChair = 1000
)

// money 0:结算 1:充值 2:平台扣除 3:平台奖励 4:冻结 5:退税 6:提取 7:购买房卡 8:消耗房卡 9:置换房卡
//
//	金币结算, 则前8位 置1。 即 Code |= 1 << 8
const (
	CodeSettle = iota + 1
	CodeRecharge
	CodeDeduct
	CodeAward
	CodeFreeze
	CodeRefund
	CodeExtract
	CodeBuyRoomCard
	CodePayRoomCard
	CodeBarterCard
)

// 游戏状态
const (
	GameStateInit = iota + 1
	GameStateOpen
	GameStateMaintain
	GameStateClose
)

// 座位状态
const (
	SeatNoCode    = iota
	SeatSettled   //已结算
	SeatCodeLook  //看牌
	SeatHaveCode  //有操作
	SeatTrustee   //托管
	SeatChi       //吃
	SeatPong      //碰
	SeatGang      //杠
	SeatTing      //听
	SeatHu        //胡
	SeatSelfDrawn //自摸
)

// 默认的个人信息
const (
	LessLevel = 1
	LessVIP   = 0
	LessMoney = 10000000 //默认金钱
)

/*----------------------------------------------------

----------------------------------------------------*/

// 时间
const (
	Date   = 24 * 60 * 60
	Hour   = 60 * 60
	Minute = 60
)

// ByteSize 流量
type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// GameControlClear 游戏控制
const (
	GameControlClear = iota //游戏清场
)

// MainGameUpdate 玩家行为
const (
	GameUpdateEnter     = iota //玩家入场
	GameUpdateOut              //玩家出场
	GameUpdateOffline          //玩家离线
	GameUpdateReconnect        //玩家重连
	GameUpdateReady            //玩家准备
	GameUpdateHost             //玩家抢庄
	GameUpdateSuperHost        //玩家超级抢庄
	GameUpdateGold             //玩家金币变化
	GameUpdateCall             //玩家叫分
	GameUpdateTrustee          //玩家托管
	GameUpdateSiting           //座位变化
	GameUpdateFollow           //跟注
	GameUpdateRaise            //加注
	GameUpdateLook             //看牌
	GameUpdateCompare          //比牌
	GameUpdateHostList         //抢庄列表
	GameUpdateRollDice         //掷骰子
	GameUpdateDisbanded        //解散
)

var OnlineCount int32 = 0 // 在线人数统计
var RecvMessage = map[int]string{
	GameUpdateEnter:     "玩家入场",
	GameUpdateHost:      "玩家抢庄",
	GameUpdateSuperHost: "玩家超级抢庄",
	GameUpdateOut:       "玩家出场",
	GameUpdateOffline:   "玩家离线",
	GameUpdateReconnect: "玩家重连",
	GameUpdateReady:     "玩家准备",
	GameUpdateGold:      "玩家金币变化",
	GameUpdateCall:      "玩家叫分",
	GameUpdateTrustee:   "玩家托管",
	GameUpdateSiting:    "座位变化",
	GameUpdateFollow:    "跟注",
	GameUpdateRaise:     "加注",
	GameUpdateLook:      "看牌",
	GameUpdateCompare:   "比牌",
	GameUpdateHostList:  "抢庄列表",
	GameUpdateRollDice:  "掷骰子",
}
