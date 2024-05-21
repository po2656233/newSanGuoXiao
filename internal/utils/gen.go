package utils

import (
	"strconv"
	"strings"
	"time"
)

// 随机生成字符 或 指定字符生成
// GenerateGameNum 生成游戏牌局号
func GenerateGameNum(gameName string, level, tableID int32) string {
	return gameName + strconv.Itoa(int(level)) + strconv.Itoa(int(tableID)) + strconv.FormatInt(time.Now().Unix(), 10)
}

// GenerateRobotNum 生成机器人牌局号
func GenerateRobotNum(tableID int64) string {
	return strconv.Itoa(int(tableID)) + strconv.FormatInt(time.Now().Unix(), 10)
}
func ReplaceLast(s, old, new string) string {
	lastIndex := strings.LastIndex(s, old)
	if lastIndex == -1 {
		// 如果old不在s中，直接返回s
		return s
	}

	// 计算替换后的字符串长度
	newLen := len(new)
	oldLen := len(old)
	resultLen := len(s) - oldLen + newLen

	// 创建一个新的切片，长度为替换后的字符串长度
	result := make([]byte, resultLen)

	// 将s中的内容复制到结果切片中，直到最后一个匹配项的前面
	copy(result, s[:lastIndex])

	// 将新字符串复制到结果切片中，替换最后一个匹配项
	copy(result[lastIndex:], new)

	// 如果原字符串在最后一个匹配项之后还有内容，将其追加到结果切片中
	if oldLen < newLen {
		copy(result[lastIndex+newLen:], s[lastIndex+oldLen:])
	}

	return string(result)
}
