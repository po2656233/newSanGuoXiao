package utils

// 各式各样的 随机数值
// 生成随机字符串
import (
	"math/rand"
	"time"
)

const (
	character = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789;:'[]{}~!@#$%^&*()<>?/\\`.,|-_+="
	letter    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomStr 随机字符
func RandomStr(count int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	b := make([]byte, count)
	for i := range b {
		b[i] = character[rand.Intn(len(character))]
	}
	return string(b)
}

// RandomStrLetter 随机字母
func RandomStrLetter(count int) string {
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, count)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// RandomNum 生成固定随机 elementMax元素的最大值(注:最后一个元素不受限制,其他元素必须等于或小于elementMax) remainder余数 count元素个数
func RandomNum(elementMax, remainder, count int) ([]int64, bool) {
	var data []int64

	if count < 1 || elementMax < 1 || remainder == 0 {
		return data, false
	}

	if remainder/2 < elementMax {
		elementMax = remainder / 2
	}

	//fmt.Printf("RandomNum-> remain:%v count:%v ", remainder, count)
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count-1; i++ {
		data = append(data, int64(rand.Int()%elementMax)+1)
	}
	sum := int64(0)
	for _, v := range data {
		sum += v
	}
	data = append(data, int64(remainder)-sum%int64(remainder))
	return data, true
}
