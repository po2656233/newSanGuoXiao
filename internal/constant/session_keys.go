package constant

import "github.com/po2656233/superplace/const"

const (
	ServerID = "server_id" // int32 游戏服务器ID
	OpenID   = "open_id"   // string 第三方登陆sdk的用户唯一标识
	PID      = "pid"       // int32 sdk包id
	PlayerID = "player_id" // int64 玩家id
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
