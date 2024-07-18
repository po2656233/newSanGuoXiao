package constant

import superConst "github.com/po2656233/superplace/const"

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
	MIDGate = 1
	MIDLeaf = 2
	MIDGame = 3
	MIDPing = 4
)
const (
	NodeTypeGate = "gate"
	NodeTypeGame = "game"
	NodeTypeLeaf = "leaf"
)
const (
	ActorGate = "user"
	ActorGame = "game"
)

const (
	FuncLogin = "login"
	FuncEnter = "enter"
)

const (
	DbList       = "db_id_list"
	CenterDb     = "center_db_id"
	GameDb       = "game_db_id"
	TcpAddr      = ":10011"
	SessionClose = "sessionClose"
)

// sdk平台类型
const (
	DevMode  int32 = 1 // 开发模式，注册开发帐号登陆(开发时使用)
	QuickSDK int32 = 2 // quick sdk
)

// 源码中重要标识
const (
	INVALID    = 0  //无效(切记有效初始化,不要从零开始)
	FAILED     = 1  // 失败
	Fault      = -1 // 故障
	Default    = 1  // 默认
	SUCCESS    = 0  // 成功
	IndexStart = 1  // 起始索引

	SYSTEMID = 0

	ADD = 0 // 新增
	DEL = 1 // 删除

	Lose = 0 // 输
	Win  = 1 // 赢
	Draw = 2 // 平局

	Limit     = 99 // 限制
	Unlimited = -1
)

func Join(sources ...string) string {
	var result string
	size := len(sources)
	for i, source := range sources {
		result += source
		if i != size-1 {
			result += superConst.DOT
		}
	}
	return result
}
