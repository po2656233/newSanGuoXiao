package constant

import "fmt"

const (
	KeyToken    = "token_"
	KeyMsgProto = "msg_proto"
)

func GetTokenKey(account string) string {
	return fmt.Sprintf("%s%s", KeyToken, account)
}
