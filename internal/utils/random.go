package utils

// 各式各样的 随机数值
// 生成随机字符串
import (
	"bytes"
	rand2 "crypto/rand"
	"math/big"
	"math/rand"
	"strconv"
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

// GetRandString 生成n位随机数字字符串
func GetRandString(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(RandIntn(10)))
	}
	return buffer.String()
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

// RandIntn 获取一个 0 ~ n 之间的随机值
func RandIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func GenRandNum(min, max int) int {
	mus := max - min
	if mus < 0 {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(mus+1)
}
func GenRandNum32(min, max int32) int32 {
	mus := max - min
	if mus < 0 {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int31n(mus+1)
}
func GenRandNum64(min, max int64) int64 {
	mus := max - min
	if mus < 0 {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(mus+1)
}

// GetRandList64 从指定的列表中获取随机元素
func GetRandList64(src []int64, count int, canRepeat bool) []int64 {
	size := len(src)
	if size < 2 {
		return src
	}
	dest := make([]int64, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 可重复
	if canRepeat {
		for i := 0; i < count; i++ {
			dest = append(dest, src[r.Int()%size])
		}
		return dest
	}

	// 不重复
	if size <= count {
		return src
	}
	copy(dest, src)
	r.Shuffle(count, func(i, j int) {
		dest[i], dest[j] = dest[j], dest[i]
	})
	return dest[:count]
}

// IsSatisfy 是否命中 百分比
func IsSatisfy(rate int) bool {
	r2, _ := rand2.Int(rand2.Reader, big.NewInt(100))
	return 0 == r2.Int64()%int64(rate)
}

// GetOneDay 获取当天0点和24点时间戳
// beginTimeNum  0点
// endTimeNum  24点
func GetOneDay() (beginTimeNum, endTimeNum int64) {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	beginTimeNum = t.Unix()
	endTimeNum = beginTimeNum + 86400
	return beginTimeNum, endTimeNum
}
