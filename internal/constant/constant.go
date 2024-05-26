package constant

const (
	GameNodeType = "game"
)

const (
	DbList   = "db_id_list"
	CenterDb = "center_db_id"
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

	Limit = 99 // 限制
)
