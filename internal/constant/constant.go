package constant

// sdk平台类型
const (
	DevMode  int32 = 1 // 开发模式，注册开发帐号登陆(开发时使用)
	QuickSDK int32 = 2 // quick sdk
)
const (
	MIDGate  = 1
	MIDLeaf  = 2
	MIDGame  = 3
	MIDPing  = 234
	MIDMatch = 5 //匹配服
)
const (
	DbList    = "db_id_list"
	CenterDb  = "center_db_id"
	GameDb    = "game_db_id"
	TcpAddr   = ":10011"
	KcpAddr   = ":10012"
	Empty     = ""
	MSGFile   = "config/common/message_id.json"
	RedisConf = "config/common/redis.toml"
)

// 源码中重要标识
const (
	INVALID = 0  //无效(切记有效初始化,不要从零开始)
	SUCCESS = 0  // 成功
	FINISH  = 0  // 成功
	FAILED  = 1  // 失败
	Fault   = -1 // 故障
	Default = 1  // 默认

	IndexStart = 1 // 起始索引
	SYSTEMID   = 0

	ADD = 0 // 新增
	DEL = 1 // 删除

	Lose = 0 // 输
	Win  = 1 // 赢
	Draw = 2 // 平局

	Ten       = 10
	Twenty    = 20
	FullScore = 100 // 满分
	Limit     = 99  // 限制
	Unlimited = -1  // 无限制

	MaxLoadNum = 20 //最大录单
)
const (
	NameLenMin = 3
	NameLenMax = 18
)

// money 0:结算 1:充值 2:平台扣除 3:平台奖励 4:冻结 5:退税 6:提取 7:购买房卡 8:消耗房卡 9:置换房卡
//
//	金币结算, 则前8位 置1。 即 Code |= 1 << 8
const (
	CodeSettle = iota
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

const (
	Username = "username"
	Password = "password"

	ProcessId = "pid" //包id 用于控制客户端
	ServerId  = "sid"
	RoomId    = "rid"
	TableId   = "tid"
	GameId    = "gid"
	UserId    = "uid"
	openId    = "openid" //开发者ID
)
const (
	SourcePath = ".system"
	OpsActor   = ".ops"
	DBActor    = ".db"
	AccActor   = ".account"
	GameActor  = ".game"
	MatchActor = ".match"
)

const (
	NodeTypeCenter = "center"
	NodeTypeGate   = "gate"
	NodeTypeGame   = "game"
	NodeTypeLeaf   = "leaf"
	NodeTypeMatch  = "match"
	NodeTypeChat   = "chat" // 聊天服
)
const (
	ActIdGate  = "user"
	ActIdGame  = "game"
	ActIdMatch = "match"
	ActIdChat  = "chat"
)

const (
	FuncLogin        = "login"
	FuncEnter        = "enter"
	FuncRequest      = "request"
	FuncSessionClose = "sessionClose"
)
