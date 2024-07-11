package constant

import "fmt"

const (
	KeyToken = "token_"
)

func GetTokenKey(account string) string {
	return fmt.Sprintf("%s%s", KeyToken, account)
}
