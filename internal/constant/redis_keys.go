package constant

import "fmt"

const (
	KeyLock      = "redis_lock"
	KeyToken     = "token_"
	KeyUser      = "user_"
	KeyMsgProto  = "msg_proto"
	KeyMsgTables = "tables_"
)

func GetTokenKey(account string) string {
	return fmt.Sprintf("%s%s", KeyToken, account)
}

// GetMatchKey
// value的结构 tableId,maxSit,nowSitCount,nodeServerId
// score即时maxSit-nowSitCount 剩余的是结构
func GetMatchKey(gameId, roomId int64) string {
	return fmt.Sprintf("%sg%d_r%d", KeyMsgTables, gameId, roomId)
}

func GetUserKey(uid int64) string {
	return fmt.Sprintf("%s%d", KeyUser, uid)
}
