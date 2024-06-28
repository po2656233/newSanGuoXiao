package base

import "fmt"

const (
	KeyNormalNotify = "notice_normal"
	KeyOnlineCount  = "online_count"
	KeyMsgProto     = "msg_proto_list"
)

func GetToken(uid int64) string {
	return fmt.Sprintf("TOKEN_%v", uid)
}

func GetGameKey(gid int64) string {
	return fmt.Sprintf("Game_%v", gid)
}
func GetGameKindKey(gid int64) string {
	return fmt.Sprintf("GameKind_%v", gid)
}
func GetAccountKey(acc string) string {
	return fmt.Sprintf("acc_%v", acc)
}

// GetAddressKey 对应的value是平台ID+账号
func GetAddressKey(addr string) string {
	return fmt.Sprintf("addr_%v", addr)
}

//func GetMaxMsgIdKey() string {
//	return "max_proto_"
//}
